package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadImage(fieldName string, r *http.Request) (string, error) {
	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	uploadPath := filepath.Join("data/image", handler.Filename)
	destination, err := os.Create(uploadPath)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		return "", err
	}
	return "/data/image/" + handler.Filename, nil
}
