package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type DokumenLegalEntity struct {
	Id   int    `json:"dokumen_id"`
	UmkmId   uuid.UUID    `json:"umkm_id"`
	DokUpload domain.JSONB `json:"dok_upload"`
	TanggalUpload time.Time `json:"tanggal_upload"`
}

func ToDokumenLegalEntity(dokumenlegal domain.UmkmDokumen) DokumenLegalEntity {
	return DokumenLegalEntity{
		Id:   dokumenlegal.Id,
		UmkmId: dokumenlegal.UmkmId,
		DokUpload: dokumenlegal.DokumenUpload,
		TanggalUpload: dokumenlegal.CreatedAt,
	}
}

func ToDokuemenLegalEntities(dokumenumkmList []domain.UmkmDokumen) []DokumenLegalEntity {
	var dokumenEntities []DokumenLegalEntity
	for _, dokumenumkm := range dokumenumkmList {
		dokumenEntities = append(dokumenEntities, ToDokumenLegalEntity(dokumenumkm))
	}
	return dokumenEntities
}