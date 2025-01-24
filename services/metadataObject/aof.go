package metadataObject

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Aof struct {
	file *os.File
	mu   sync.Mutex
}

var AOF *Aof

func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: f,
	}

	go func() {
		for {
			aof.mu.Lock()

			aof.file.Sync()

			aof.mu.Unlock()

			time.Sleep(time.Second)
		}
	}()

	return aof, nil
}

func (aof *Aof) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	return aof.file.Close()
}

func (aof *Aof) Write(value FileMapType) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return fmt.Errorf("failed to write to AOF: %v", err)
	}

	return nil
}

func (aof *Aof) Read(fn func(value FileMapType)) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	aof.file.Seek(0, io.SeekStart)

	reader := NewResp(aof.file)

	for {
		value, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		fn(value)
	}

	return nil
}

func InitializeAoF() {
	var err error
	AOF, err = NewAof("metadata.aof")
	if err != nil {
		log.Fatal("Failed to create AOF file", err)
		return
	}

	AOF.Read(func(value FileMapType) {
		command := strings.ToUpper(value.Type)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(value)
	})
}
