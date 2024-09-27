package kategoriprodukrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type KategoriProdukRepository interface {
	CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error)
	GetKategoriProduk( filters string, limit int, page int) ([]domain.KategoriProduk, int,int, int, *int, *int, error)
	GetKategoriProdukId(idproduk int) (domain.KategoriProduk, error)
	UpdateKategoriId(idProduk int, kategori domain.KategoriProduk) (domain.KategoriProduk, error)
	DeleteKategoriProdukId(idproduk int) error
	DeleteKategoriUmkmId(id uuid.UUID) error
}
