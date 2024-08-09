package umkmservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm( umkm web.UmkmRequest, UserId int) (map[string]interface{}, error)
	GetUmkmList()([]entity.UmkmEntity, error)
}