package kategoriumkmservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type KategoriUmkm interface {
	CreateKategori(kategori web.CreateCategoriUmkm) (map[string]interface{}, error)
	GetKategoriUmkmList() ([]entity.KategoriEntity, error)
	GetKategoriUmkmId(id int) (entity.KategoriEntity, error)
	UpdateKategori(request web.UpdateCategoriUmkm, pathId int) (map[string]interface{}, error)
	DeleteKategoriUmkmId(id int) error
}