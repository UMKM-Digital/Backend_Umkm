package kategoriprodukservice

import (
	"umkm/model/entity"
	"umkm/model/web"


)

type KategoriUmkm interface {
	CreateKategori(kategoriproduk web.CreateCategoriProduk) (map[string]interface{}, error)
	GetKategoriProdukList(filters string, limit int, page int) ([]entity.KategoriProdukEntity, int, int, int, *int, *int, error)
	GetKategoriProdukId(id int) (entity.KategoriProduk, error)
	UpdateKategoriProduk(request web.UpdateCategoriProduk, pathId int) (map[string]interface{}, error)
	DeleteKategoriProdukId(idproduk int) error
}
