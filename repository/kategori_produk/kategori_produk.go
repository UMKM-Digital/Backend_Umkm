package kategoriprodukrepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type KategoriProdukRepository interface {
	CreateKategoriProduk(kategoriproduk domain.KategoriProduk) (domain.KategoriProduk, error)
	GetKategoriUmkm(umkmID uuid.UUID)([]domain.KategoriProduk, error)
}
