package metadataObject

import (
	"bufio"
	"io"
	"strconv"
)

type FileMapType struct {
	Type         string
	Filekey      string
	FileMetadata string
	FileHash     string
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], nil
}

func (r *Resp) readInteger() (x int, err error) {
	line, err := r.readLine()
	if err != nil {
		return 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

func (r *Resp) Read() (FileMapType, error) {
	v := FileMapType{}

	for i := 0; i < 3; i++ {
		len, err := r.readInteger()
		if err != nil {
			return v, err
		}

		str := make([]byte, len)

		r.reader.Read(str)

		if i == 0 {
			v.Type = string(str)
		} else if i == 1 {
			v.Filekey = string(str)
		} else {
			if v.Type == "METADATA" {
				v.FileMetadata = string(str)
			} else if v.Type == "HASH" {
				v.FileHash = string(str)
			}
		}

		r.readLine()
	}

	return v, nil
}

func (v FileMapType) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, strconv.Itoa(len(v.Type))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Type...)
	bytes = append(bytes, '\r', '\n')

	bytes = append(bytes, strconv.Itoa(len(v.Filekey))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Filekey...)
	bytes = append(bytes, '\r', '\n')

	if v.Type == "METADATA" {
		bytes = append(bytes, strconv.Itoa(len(v.FileMetadata))...)
		bytes = append(bytes, '\r', '\n')
		bytes = append(bytes, v.FileMetadata...)
		bytes = append(bytes, '\r', '\n')
	} else if v.Type == "HASH" {
		bytes = append(bytes, strconv.Itoa(len(v.FileHash))...)
		bytes = append(bytes, '\r', '\n')
		bytes = append(bytes, v.FileHash...)
		bytes = append(bytes, '\r', '\n')
	} else {
		return []byte{}
	}

	return bytes
}

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (w *Writer) Write(v FileMapType) error {
	var bytes = v.Marshal()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
