package produkrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type CreateProduk interface {
	CreateRequest(produk domain.Produk)(domain.Produk, error)
	DeleteProdukId(id uuid.UUID) error
	FindById(id uuid.UUID) (domain.Produk, error)
}
