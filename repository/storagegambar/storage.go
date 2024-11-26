package storagegambarrepo

import "net/http"

type StorageGambarRepo interface {
	StorageGambar(r *http.Request) ([]string, error)
}