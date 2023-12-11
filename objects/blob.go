package objects

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/leonardodelira/go-git-clone/hash"
)

func CreateBlob(file *os.File) (*os.File, error) {
	reader := bufio.NewReader(file)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	hash := hash.New(content)
	folder := hash[0:2]
	filename := hash[2:]

	path := filepath.Join(".", ".fit", "objects", folder)
	if err = os.Mkdir(path, 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	file, err = os.OpenFile(filepath.Join(path, filename), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	compressor := gzip.NewWriter(&b)
	if _, err := compressor.Write(content); err != nil {
		return nil, err
	}
	compressor.Close()
	_, err = io.WriteString(file, b.String())

	if err != nil {
		return nil, err
	}

	return file, err
}
