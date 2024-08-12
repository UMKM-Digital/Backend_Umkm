package umkmservice

import (
	"context"
	"umkm/model/entity"
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm( umkm web.UmkmRequest, UserId int) (map[string]interface{}, error)
	GetUmkmListByUserId(ctx context.Context, userId int) ([]entity.UmkmEntity, error)
}