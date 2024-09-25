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
}