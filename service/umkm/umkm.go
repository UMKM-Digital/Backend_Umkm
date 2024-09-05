package umkmservice

import (
	"context"
	"mime/multipart"
	"umkm/model/entity"
	"umkm/model/web"
)

type Umkm interface {
	CreateUmkm( umkm web.UmkmRequest, UserId int, files map[string]*multipart.FileHeader) (map[string]interface{}, error)
	GetUmkmListByUserId(ctx context.Context, userId int,filters string, limit int, page int) (map[string]interface{}, error)
	GetUmkmFilter(ctx context.Context, userID int, filters map[string]string, allowedFilters []string) ([]entity.UmkmFilterEntity, error)
	GetUmkmListWeb(ctx context.Context, userId int)([]entity.UmkmEntityList, error)
}