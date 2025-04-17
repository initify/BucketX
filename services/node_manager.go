package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/lafikl/consistent"
	"github.com/gin-gonic/gin"
)

type Node struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var (
	nodesFile = "config/nodes.json"

	nodes   []Node
	hash    *consistent.Consistent
	mu      sync.RWMutex
)

// loads nodes
func Initialize() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(nodesFile)
	if err != nil {
		if os.IsNotExist(err) {
			nodes = []Node{}
			hash = consistent.New()
			return nil
		}
		return fmt.Errorf("failed to read nodes file: %w", err)
	}

	if err := json.Unmarshal(data, &nodes); err != nil {
		return fmt.Errorf("failed to parse nodes file: %w", err)
	}

	hash = consistent.New()
	for _, node := range nodes {
		hash.Add(node.URL)
	}

	return nil
}

// GetNodeForKey returns the URL of the node for the given key
func GetNodeForKey(key string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	if hash == nil {
		return "", fmt.Errorf("hash ring not initialized")
	}

	return hash.Get(key)
}

// GetAllNodes returns a copy of the node list
func GetAllNodes() []Node {
	mu.RLock()
	defer mu.RUnlock()

	return append([]Node(nil), nodes...)
}

func LoadNodes() ([]Node, error) {
	data, err := os.ReadFile(nodesFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read nodes file: %w", err)
	}

	var nodes []Node
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nodes file: %w", err)
	}

	return nodes, nil
}

func SaveNodes(nodes []Node) error {
	data, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize nodes: %w", err)
	}

	err = os.WriteFile(nodesFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write nodes file: %w", err)
	}

	return nil
}

func AddNode(newNode Node) error {
	mu.Lock()
	defer mu.Unlock()

	nodes, err := LoadNodes()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	for _, n := range nodes {
		if n.URL == newNode.URL {
			return fmt.Errorf("node already exists")
		}
	}

	nodes = append(nodes, newNode)
	hash.Add(newNode.URL)

	// Update global list and save
	SaveNodes(nodes)
	return Initialize()
}

func isLocalNode(nodeURL string) bool {
	return strings.Contains(nodeURL, "localhost") || strings.Contains(nodeURL, "127.0.0.1")
}

func ForwardUploadRequest(c *gin.Context, nodeURL string) (*http.Response, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)

	writer.WriteField("bucket_id", c.PostForm("bucket_id"))
	writer.WriteField("file_key", c.PostForm("file_key"))

	writer.Close()

	req, err := http.NewRequest("POST", nodeURL+"/api/v1/files/upload", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	return client.Do(req)
}

func ForwardFetchRequest(fileKey string, fileQuery string, nodeURL string) (string, error) {
	url := nodeURL + "/api/v1/files/" + fileKey + "?tr=" + fileQuery

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err!= nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch file from node: %s", nodeURL)
	}
	body, err := io.ReadAll(resp.Body)
	if err!= nil {
		return "", err
	}
	return string(body), nil
}
