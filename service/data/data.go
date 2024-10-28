package dataservice

import (
	"context"
	datarepo "umkm/repository/data"
)

type AuthUserService interface {
	CountAtas() (map[string]interface{}, error)
	GrafikKategoriBySektor(ctx context.Context, sektorUsahaID int, kecamatan, kelurahan string, tahun int) ([]datarepo.KategoriCount, error)
	TotalUmkmKriteriaUsahaPerBulan(tahun int) (map[string]map[string]int64, error)
	TotalUmkmBinaan()(map[string]interface{}, error)
	TotalOmzetBulanIni()(map[string]interface{}, error)
	DataUmkm(id int)(map[string]interface{}, error)
	DataOmzetUmkm(id int, tahun int) (map[string]interface{}, error)
}
