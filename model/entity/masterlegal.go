package entity

import (
	"umkm/model/domain"
)


type MasterlegalEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	IsWajib int `json:"is_wajib"`
}

func ToMasterlegalEntity(masterlegal domain.MasterDokumenLegal) MasterlegalEntity {
	return MasterlegalEntity{
		Id: masterlegal.IdMasterDokumenLegal,
		Name: masterlegal.Name,
		IsWajib: masterlegal.Iswajib,
	}
}

func TomasterlegalEntities(masterlegalList []domain.MasterDokumenLegal) []MasterlegalEntity {
    var masterlegalEntities []MasterlegalEntity
    for _, masterlegal := range masterlegalList {
        masterlegalEntities = append(masterlegalEntities, ToMasterlegalEntity(masterlegal))
    }
    return masterlegalEntities
}