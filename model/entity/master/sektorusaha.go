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
//betnukusaha
type BentukUsahaEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToBentukUsahaEntity(bentukusaha domain.BentukUsaha) BentukUsahaEntity {
	return BentukUsahaEntity{
		Id:   bentukusaha.Id,
		Name: bentukusaha.Nama,
	}
}

func ToBentukUsahaEntities(bentukusahaList []domain.BentukUsaha) []BentukUsahaEntity {
	var bentukusahaEntities []BentukUsahaEntity
	for _, bentukusaha := range bentukusahaList {
		bentukusahaEntities = append(bentukusahaEntities, ToBentukUsahaEntity(bentukusaha))
	}
	return bentukusahaEntities
}
//statustempat usaha
type StatusTempatUsahaEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToStatusTempatUsahaEntity(bentukusaha domain.StatusTempatUsaha) StatusTempatUsahaEntity {
	return StatusTempatUsahaEntity{
		Id:   bentukusaha.Id,
		Name: bentukusaha.Nama,
	}
}

func ToStatusTempatUsahaEntities(statususahaList []domain.StatusTempatUsaha) []StatusTempatUsahaEntity {
	var statustempatuusahaEntities []StatusTempatUsahaEntity
	for _, statustempatusaha := range statususahaList {
		statustempatuusahaEntities = append(statustempatuusahaEntities, ToStatusTempatUsahaEntity(statustempatusaha))
	}
	return statustempatuusahaEntities
}