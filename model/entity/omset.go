package entity

import (
	"umkm/model/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OmsetEntity struct {
    Id   int    `json:"id"`
	UmkmId uuid.UUID `json:"umkm_id"`
	Bulan string `json:"bulan"`
	Nominal decimal.Decimal `json:"nominal"`
	Status string `json:"status"`
}

func ToOmsetEntityList(omset domain.Omset) OmsetEntity {
	return OmsetEntity{
		Id:   omset.IdOmste,
	    UmkmId: omset.UmkmId,
		Bulan: omset.Bulan,
		Nominal: omset.Nominal,
		Status: "true",
	}
}

// Fungsi untuk mengonversi daftar domain.UMKM ke daftar OmsetEntity (versi sederhana)
func ToOmsetListEntities(omsetList []domain.Omset) []OmsetEntity {
	var omsetListEntities []OmsetEntity
	for _, omset := range omsetList {
		omsetListEntities = append(omsetListEntities, ToOmsetEntityList(omset))
	}
	return omsetListEntities
}