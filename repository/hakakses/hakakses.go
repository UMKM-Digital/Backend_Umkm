package hakaksesrepo

import (
	"context"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type CreateHakakses interface {
	CreateHakAkses(hakAkses *domain.HakAkses) error
	GetHakAksesByUserId(ctx context.Context, userId int) ([]domain.HakAkses, error)
	DeleteUmkmId(id uuid.UUID) error
}