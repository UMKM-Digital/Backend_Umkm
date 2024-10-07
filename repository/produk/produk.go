package produkrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type CreateProduk interface {
	CreateRequest(produk domain.Produk)(domain.Produk, error)
	DeleteProdukId(id uuid.UUID) error
	FindById(id uuid.UUID) (domain.Produk, error)
	// ProdukById(id uuid.UUID) (domain.Produk, error)
	GetProduk(ProdukId uuid.UUID, filters string, limit int, page int, kategori_produk_id string, sort string) ([]domain.Produk, int, int, int, *int, *int, error)
	UpdatedProduk(ProdukId uuid.UUID, produk domain.Produk) (domain.Produk, error)
	GetProductsByUmkmIds(ctx context.Context, umkmIDs []uuid.UUID) ([]domain.Produk, error)
	GetProdukList(limit int, page int, filters string, kategoriproduk string, sort string) ([]domain.Produk, int, int, int, *int, *int, error)
	FindWebId(id uuid.UUID) (domain.Produk, error)
	DeleteProdukUmkmId(id uuid.UUID) error
	GetProdukByUmkmId(id uuid.UUID) ([]domain.Produk, error) 
	GetProdukByUmkmLogin(umkmIds []uuid.UUID) ([]domain.Produk, error)
	GetProdukBaru(umkmId uuid.UUID)([]domain.Produk, error)
	GetTopProduk(idUmkm uuid.UUID) ([]domain.Produk, error) 
	UpdateTopProduk(idproduk uuid.UUID, active int) error
	GetProdukActive(active int)([]domain.Produk, error)
}