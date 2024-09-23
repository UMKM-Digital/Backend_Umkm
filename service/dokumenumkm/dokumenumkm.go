package dokumenumkmservice

import (
	"mime/multipart"
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type DokumenUmkmService interface {
	CreateDokumenUmkm(produk web.CreateUmkmDokumenLegal, files []*multipart.FileHeader) (map[string]interface{}, error)
	GetDokumenUmkmId(id int, umkm uuid.UUID)(entity.DokumenLegalEntity, error)
	UpdateDokumenUmkm(id int, umkmId uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error)
}