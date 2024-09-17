package beritaservice

import (
	"context"
	"mime/multipart"
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
)

type Berita interface {
	CreatedBerita(Berita web.CreatedBerita, file *multipart.FileHeader, userID int) (map[string]interface{}, error)
	GetBeritaList(ctx context.Context, limit int, page int) ([]entity.BeritaFilterEntity, int, int, int, *int, *int, error)
	DeleteBerita(id int) error
	GetBeritaByid(id int) (entity.BeritaFilterEntity, error)
	UpdateBerita(request web.UpdtaedBerita, Id int,file *multipart.FileHeader) (map[string]interface{}, error)
}