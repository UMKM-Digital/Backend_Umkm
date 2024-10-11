package daerahservice

import entity "umkm/model/entity/master"

type Daerah interface {
	GetDaerah() ([]entity.DaerahEntity, error)
	GetKabupaten(id string) ([]entity.KabupatenEntity, error)
	GetKecamatan(id string) ([]entity.KecamatanEntity,error)
	GetKelurahan(id string) ([]entity.KeluarahanEntity, error)
}