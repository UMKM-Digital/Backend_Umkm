package storagegambarrepo

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type StorageGambar interface {
	StorageGambar(r *http.Request) ([]string, error)
}

type StorageGambarRepoImpl struct{}

func NewStorageGambarRepo() StorageGambar {
	return &StorageGambarRepoImpl{}
}

func (repo *StorageGambarRepoImpl) StorageGambar(r *http.Request) ([]string, error) {
	// Parsing form data untuk mendukung file upload
	err := r.ParseMultipartForm(10 << 20) // Maksimal ukuran file 10MB
	if err != nil {
		return nil, errors.New("failed to parse multipart form")
	}

	// Ambil semua file dari form field "file"
	formFiles := r.MultipartForm.File["file"]
	if len(formFiles) == 0 {
		return nil, errors.New("no files provided in the request")
	}

	// Direktori untuk menyimpan file
	storagePath := "./uploads/Produk/"
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create storage directory")
		}
	}

	var filePaths []string

	// Iterasi setiap file dan simpan
	for _, fileHeader := range formFiles {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, errors.New("failed to open file")
		}
		defer file.Close()

		filePath := filepath.Join(storagePath, fileHeader.Filename)

		// Pastikan path menggunakan slash ('/')
		filePath = filepath.ToSlash(filePath)

		out, err := os.Create(filePath)
		if err != nil {
			return nil, errors.New("failed to create file")
		}
		defer out.Close()

		if _, err = io.Copy(out, file); err != nil {
			return nil, errors.New("failed to save file")
		}

		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}
