package storagegambarservice

import "net/http"

type StorageGambarProduk interface {
	SaveImages(r *http.Request) ([]string, error)
}