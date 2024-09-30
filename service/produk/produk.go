package produkservice

import (
	"mime/multipart"
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type Produk interface {
	CreateProduk(produk web.WebProduk, files map[string]*multipart.FileHeader) (map[string]interface{}, error)
	DeleteProduk(id uuid.UUID) error
	GetProdukId(id uuid.UUID)(entity.ProdukEntity, error)
	GetProdukList(Produkid uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]entity.ProdukList, int, int, int, *int, *int, error) 
	UpdateProduk(request web.UpdatedProduk, id uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error) 
	GetProduk(limit int, page int, filters string, kategoriproduk string, sort string) ([]entity.ProdukWebEntity, int, int, int, *int, *int, error) 
	GetProdukWebId(id uuid.UUID)(entity.ProdukWebIdEntity, error)
}