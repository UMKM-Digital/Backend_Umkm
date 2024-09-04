package kategoriprodukrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type KategoriProdukRepository interface {
	CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error)
	GetKategoriProduk(umkmID uuid.UUID, filters string, limit int, page int) ([]domain.KategoriProduk, int, error)
}
