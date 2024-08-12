package hakaksesrepo

import (
	"context"
	"umkm/model/domain"
)

type CreateHakakses interface {
	CreateHakAkses(hakAkses *domain.HakAkses) error
	GetHakAksesByUserId(ctx context.Context, userId int) ([]domain.HakAkses, error)
}