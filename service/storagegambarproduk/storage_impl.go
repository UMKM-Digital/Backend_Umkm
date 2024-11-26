package storagegambarservice

import (
	"net/http"
	storagegambarrepo "umkm/repository/storagegambar"
)

type StorageGambar interface {
	SaveImages(r *http.Request) ([]string, error) // Ubah untuk mendukung array
}

type StorageGambarProdukImpl struct {
	repo storagegambarrepo.StorageGambar // Pastikan menggunakan interface yang benar
}

func NewStorageGambarService(repo storagegambarrepo.StorageGambar) StorageGambar {
	return &StorageGambarProdukImpl{repo: repo}
}

func (s *StorageGambarProdukImpl) SaveImages(r *http.Request) ([]string, error) {
	return s.repo.StorageGambar(r) // Panggil fungsi repositori yang mengembalikan banyak file
}
