package umkmrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type CreateUmkm interface {
	CreateRequest(umkm domain.UMKM)(domain.UMKM, error)
	// GetUmkmList()([]domain.UMKM, error)
	GetUmkmListByIds(ctx context.Context, umkmIDs []uuid.UUID, filters string, limit int, page int) ([]domain.UMKM, int, int, int, *int, *int, error)
	GetUmkmFilterName(ctx context.Context,  umkmIDs []uuid.UUID)([]domain.UMKM, error)
	GetUmkmListWeb(ctx context.Context, umkmIds []uuid.UUID)([]domain.UMKM, error)
	GetUmkmID(id uuid.UUID)(domain.UMKM, error)
	UpdateUmkmId(id uuid.UUID, umkm domain.UMKM)(domain.UMKM, error)
	GetUmkmList(filters string, limit int, page int, kategori_umkm string, sortOrder string) ([]domain.UMKM, int, int, int, *int, *int, error) 
	// GetUmkmListDetailId(id uuid.UUID) ([]domain.UMKM, error)
	GetUmkmListDetailPaginated(id uuid.UUID, limit int, page int) ([]domain.UMKM, int, int, int, *int, *int, error) 
	DeleteUmkmId(id uuid.UUID) error
	FindById(umkmId uuid.UUID) (domain.UMKM, error) 
}