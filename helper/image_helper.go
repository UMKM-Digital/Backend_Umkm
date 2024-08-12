package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func HandleFileUpload(file *multipart.FileHeader, uploadDir string) (string, error) {
    src, err := file.Open()
    if err != nil {
        return "", fmt.Errorf("failed to open file %s: %w", file.Filename, err)
    }
    defer src.Close()

    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        return "", fmt.Errorf("failed to create directory %s: %w", uploadDir, err)
    }

    filename := filepath.Base(file.Filename)
    dst, err := os.Create(filepath.Join(uploadDir, filename))
    if err != nil {
        return "", fmt.Errorf("failed to create file %s: %w", filename, err)
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return "", fmt.Errorf("failed to copy file %s to %s: %w", filename, dst.Name(), err)
    }

    return fmt.Sprintf("/uploads/%s", filename), nil
}

func SaveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}