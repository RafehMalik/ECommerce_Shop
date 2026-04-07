package providors

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalUploadProvidor struct {
	basePah string
}

func NewLocalUploadProvidor(path string) *LocalUploadProvidor {
	return &LocalUploadProvidor{
		basePah: path,
	}
}

func (p *LocalUploadProvidor) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	fullPath := filepath.Join(p.basePah, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return "", err
	}
	//open source
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	//open destination
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}
	return fmt.Sprintf("/Uploads/%s", path), nil
}

func (p *LocalUploadProvidor) DeleteFile(path string) error {
	fullPath := filepath.Join(p.basePah, path)
	return os.Remove(fullPath)
}
