package kategoriumkmservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type KategoriUmkm interface {
	CreateKategori(kategori web.CreateCategoriUmkm) (map[string]interface{}, error)
	GetKategoriUmkmList() ([]entity.KategoriEntity, error)
}