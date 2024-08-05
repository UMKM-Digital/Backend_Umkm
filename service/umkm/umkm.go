package umkmservice

import (
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm( umkm web.UmkmRequest) (map[string]interface{}, error)
}