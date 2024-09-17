package beritarepo

import (
	"context"
	"umkm/model/domain"
)

type BeritaRepo interface {
	CreateRequest(Berita domain.Berita) (domain.Berita, error)
	GetBeritaList(ctx context.Context, limit int, page int) ([]domain.Berita, int, int, int, *int, *int, error)
	DelBerita(id int) error
	GetBeritaByid(id int) (domain.Berita, error)
	UpdateBeritaId(id int, berita domain.Berita) (domain.Berita, error) 
}
