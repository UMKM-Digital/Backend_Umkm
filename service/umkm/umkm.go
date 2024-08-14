package umkmservice

import (
	"context"
	"mime/multipart"
	"umkm/model/entity"
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm( umkm web.UmkmRequest, UserId int, files map[string]*multipart.FileHeader) (map[string]interface{}, error)
	GetUmkmListByUserId(ctx context.Context, userId int) ([]entity.UmkmEntity, error)
}