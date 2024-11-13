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
	GetUmkmIdsByUserId(userId int) ([]uuid.UUID, error)
	GetUmkmId(umkmid uuid.UUID) (domain.HakAkses, error)
	AcceptBulkStatus(umkmids []uuid.UUID, hakakses domain.HakAkses) error
	CheckUmkmStatus(umkmId uuid.UUID) (bool, error) 
}