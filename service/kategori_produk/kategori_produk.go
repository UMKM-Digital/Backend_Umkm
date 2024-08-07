package kategoriprodukservice

import (
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type KategoriUmkm interface {
	CreateKategori(kategoriproduk web.CreateCategoriProduk) (map[string]interface{}, error)
	GetKategoriProdukList(umkmID uuid.UUID) ([]entity.KategoriProdukEntity, error)
}
