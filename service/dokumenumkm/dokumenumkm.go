package dokumenumkmservice

import (
	"mime/multipart"
	"umkm/model/web"
)

type DokumenUmkmService interface {
	CreateDokumenUmkm(produk web.CreateUmkmDokumenLegal, files []*multipart.FileHeader) (map[string]interface{}, error)
}