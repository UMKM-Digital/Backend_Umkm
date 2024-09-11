package entity

import "umkm/model/domain"

type masterlegalEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	IsWajib int `json:"is_wajib"`
}

func TomasterlegalEntity(masterlegal domain.MasterDokumenLegal) masterlegalEntity {
	return masterlegalEntity{
		Id: masterlegal.IdMasterDokumenLegal,
		Name: masterlegal.Name,
		IsWajib: masterlegal.Iswajib,
	}
}

func TomasterlegalEntities(masterlegalList []domain.MasterDokumenLegal) []masterlegalEntity {
    var masterlegalEntities []masterlegalEntity
    for _, masterlegal := range masterlegalList {
        masterlegalEntities = append(masterlegalEntities, TomasterlegalEntity(masterlegal))
    }
    return masterlegalEntities
}