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
	TotalUmkmKriteriaUsahaPerBulan(tahun int) (map[string]map[string]int64, error) 
	TotalUmkmBulan()(int64, error)
	TotalUmkmBulanLalu()(int64, error)
	TotalUmkmTahun() (int64, error)
	TotalUmkmTahunLalu() (int64, error) 
	PersentasiKenaikanUmkm() (float64, error)
	PersentasiKenaikanUmkmTahun() (float64, error)
	TotalOmzetBulanIni() (float64, error)
	TotalOmzetBulanLalu() (float64, error)
	TotalomzestTahunIni() (float64, error)
	TotalOmzetTahunLalu() (float64, error)
	Persentasiomzetbulan() (float64, error)
	Persentasiomzettahun() (float64, error)
}