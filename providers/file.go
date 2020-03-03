package providers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileProvider struct {
	file *os.File
}

func NewFileProvider() *FileProvider {
	return &FileProvider{}
}

func (fp *FileProvider) TempDir(prefix string) (string, error) {
	return ioutil.TempDir(os.TempDir(), "fitz")
}

func (fp *FileProvider) Create(tmpDir string, pageNumber int) error {
	osFile, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", pageNumber)))
	if err != nil {
		return err
	}

	fp.file = osFile
	return nil
}

func (fp *FileProvider) WriteString(s string) error {
	_, err := fp.file.WriteString(s)
	return err
}

func (fp *FileProvider) Close() error {
	return fp.file.Close()
}
