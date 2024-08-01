package umkmservice

import (
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm(kategori web.UmkmRequest) (map[string]interface{}, error)
}