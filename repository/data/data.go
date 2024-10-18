package datarepo

import (
	"context"
)

type DataUserRepo interface {
	CountProductByCategoryWithPercentage() (int64, error)
	CountWaitingVerify() (int64, error)
	CountUmkmBina() (int64, error)
	CountUmkmTertolak() (int64, error)
	TotalUmkm() (int64, error)
	TotalMikro() (int64, error)
	TotalMenengah() (int64, error)
	TotalKecil() (int64, error)
	TotalSektorJasa() (int64, error)
	TotalSektorProduksi() (int64, error)
	TotalSektorPerdagangan() (int64, error)
	TotalEkonomiKreatif() (int64, error)
	GrafikKategoriBySektorUsaha(ctx context.Context, sektorUsahaID int, kecamatan, kelurahan string, tahun int) ([]KategoriCount, error)
}