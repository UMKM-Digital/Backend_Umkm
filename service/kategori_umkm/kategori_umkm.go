package kategoriumkmservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type KategoriUmkm interface {
	CreateKategori(kategori web.CreateCategoriUmkm) (map[string]interface{}, error)
	GetKategoriUmkmList(filters string, limit int, page int) (map[string]interface{}, error)
	GetKategoriUmkmId(id int) (entity.KategoriEntity, error)
	UpdateKategori(request web.UpdateCategoriUmkm, pathId int) (map[string]interface{}, error)
	DeleteKategoriUmkmId(id int) error
}