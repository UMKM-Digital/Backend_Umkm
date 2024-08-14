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
	// GetProdukId(id uuid.UUID)(entity.ProdukEntity, error)
	 GetProdukList(Produkid uuid.UUID) ([]entity.ProdukEntity, error)
}