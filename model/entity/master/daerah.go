package entity

import domain "umkm/model/domain/master"

type DaerahEntity struct {
	Id   int    `json:"id_prov"`
	Name string `json:"name"`
}

func ToDaerahEntity(daerah domain.Provinsi) DaerahEntity {
	return DaerahEntity{
		Id:   daerah.KodeWilayah,
		Name: daerah.NamaWilayah,
	}
}

func ToDaerahEntities(daerahList []domain.Provinsi) []DaerahEntity {
	var daerahEntities []DaerahEntity
	for _, daerah := range daerahList {
		daerahEntities = append(daerahEntities, ToDaerahEntity(daerah))
	}
	return daerahEntities
}

//Kabupaten
type KabupatenEntity struct {
	IdProvinsi   string    `json:"id_prov"`
	IdKabupaten  string `json:"id_Kabupaten"` 
	Name string `json:"name"`
}

func ToKabupatenEntity(daerah domain.Kabupaten) KabupatenEntity {
	return KabupatenEntity{
		IdProvinsi:   daerah.IdProvi,
		IdKabupaten: daerah.KodeKabupaten,
		Name: daerah.NamaKabupaten,
	}
}

func ToDaerahKabupatenEntities(daerahList []domain.Kabupaten) []KabupatenEntity {
	var daerahEntities []KabupatenEntity
	for _, daerah := range daerahList {
		daerahEntities = append(daerahEntities, ToKabupatenEntity(daerah))
	}
	return daerahEntities
}

//kecamtan
type KecamatanEntity struct {
	IdKabupaten   string    `json:"id_kabupaten"`
	IdKecamatan  string `json:"id_kecamatan"` 
	Name string `json:"name"`
}

func ToKecamatanEntity(daerah domain.Kecamatan) KecamatanEntity {
	return KecamatanEntity{
		IdKabupaten:   daerah.IdKabupaten,
		IdKecamatan: daerah.KodeWilayah,
		Name: daerah.NamaWilayah,
	}
}

func ToDaerahKecamatanEntities(daerahList []domain.Kecamatan) []KecamatanEntity {
	var daerahEntities []KecamatanEntity
	for _, daerah := range daerahList {
		daerahEntities = append(daerahEntities, ToKecamatanEntity(daerah))
	}
	return daerahEntities
}

//
type KeluarahanEntity struct {
	IdKabupaten   string    `json:"id_kecamatan"`
	IdKecamatan  string `json:"id_kelurahan"` 
	Name string `json:"name"`
}

func ToKelurahanEntity(daerah domain.Keluarahan) KeluarahanEntity {
	return KeluarahanEntity{
		IdKabupaten:   daerah.KodeKecamatan,
		IdKecamatan: daerah.KodeKewilayah,
		Name: daerah.NamaWilayah,
	}
}

func ToDaerahKeluarahanEntities(daerahList []domain.Keluarahan) []KeluarahanEntity {
	var daerahEntities []KeluarahanEntity
	for _, daerah := range daerahList {
		daerahEntities = append(daerahEntities, ToKelurahanEntity(daerah))
	}
	return daerahEntities
}