package umkmservice

import (
	"context"
	"mime/multipart"
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type Umkm interface {
	// CreateUmkm( umkm web.UmkmRequest, UserId int, files map[string]*multipart.FileHeader) (map[string]interface{}, error)
	CreateUmkm(umkm web.UmkmRequest, userID int, files map[string]*multipart.FileHeader, dokumenFiles []*multipart.FileHeader, dokumenIDs []string) (map[string]interface{}, error)
	GetUmkmListByUserId(ctx context.Context, userId int, filters string, limit int, page int) ([]entity.UmkmFilterEntity, int, int, int, *int, *int, error)
	GetUmkmFilter(ctx context.Context, userID int, filters map[string]string, allowedFilters []string) ([]entity.UmkmFilterEntity, error)
	GetUmkmListWeb(ctx context.Context, userId int)([]entity.UmkmEntityList, error)
	GetUmkmId(id uuid.UUID)(entity.UmkmEntity, error)
	UpdateUmkmId(request web.Updateumkm, umkmid uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error)
	GetUmkmList(filters string, limit int, page int, kategori_umkm string, sortOrder string)([]entity.UmkmEntityWebList,int, int, int, *int, *int, error)
	// GetUmkmDetailList(id uuid.UUID)([]entity.UmkmDetailEntity, error)
	GetUmkmDetailList(id uuid.UUID, limit int, page int)([]entity.UmkmDetailEntity,int, int, int, *int, *int, error)
	Delete(id uuid.UUID) error
	GetUmkmActive() ([]entity.UmkmActive, error)
	UpdateUmkmActive(request web.UpdateActiveUmkm, Id uuid.UUID) (map[string]interface{}, error)
	GetTestimonialActive() ([]entity.UmkmActive, error) 
}