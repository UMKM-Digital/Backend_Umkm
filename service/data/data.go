package dataservice

import (
	"context"
	datarepo "umkm/repository/data"
)

type AuthUserService interface {
	CountAtas() (map[string]interface{}, error)
	GrafikKategoriBySektor(ctx context.Context, sektorUsahaID int, kecamatan, kelurahan string, tahun int) ([]datarepo.KategoriCount, error)
}
