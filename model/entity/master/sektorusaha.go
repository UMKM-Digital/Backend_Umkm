package entity

import domain "umkm/model/domain/master"

type SektorUsahaEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToSektorUsahaEntity(sektorusaha domain.SektorUsaha) SektorUsahaEntity {
	return SektorUsahaEntity{
		Id:   sektorusaha.Id,
		Name: sektorusaha.Nama,
	}
}

func ToKategoriEntities(sektorusahaList []domain.SektorUsaha) []SektorUsahaEntity {
	var sektorusahaEntities []SektorUsahaEntity
	for _, kategori := range sektorusahaList {
		sektorusahaEntities = append(sektorusahaEntities, ToSektorUsahaEntity(kategori))
	}
	return sektorusahaEntities
}