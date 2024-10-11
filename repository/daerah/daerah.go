package daerahrepo

import domain "umkm/model/domain/master"

type Daerah interface {
	GetProvinsi() ([]domain.Provinsi, error)
	GetKabupaten(id string) ([]domain.Kabupaten, error)
	GetKecamatan(id string) ([]domain.Kecamatan, error)
	GetKelurahan(id string) ([]domain.Keluarahan, error)
}