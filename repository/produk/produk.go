package produkrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type CreateProduk interface {
	CreateRequest(produk domain.Produk)(domain.Produk, error)
	DeleteProdukId(id uuid.UUID) error
	FindById(id uuid.UUID) (domain.Produk, error)
	// ProdukById(id uuid.UUID) (domain.Produk, error)
	GetProduk(ProdukId uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]domain.Produk, int, int, int, *int, *int, error)
	UpdatedProduk(ProdukId uuid.UUID, produk domain.Produk) (domain.Produk, error)
}