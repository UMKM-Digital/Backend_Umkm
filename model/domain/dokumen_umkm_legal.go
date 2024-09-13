package domain

import (
	"time"

	"github.com/google/uuid"
)

type UmkmDokumen struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	DokumenId  int       `gorm:"column:dokumen_id"`
	UmkmId     uuid.UUID `gorm:"column:umkm_id;type:uuid"`
	DokumenUpload     JSONB     `gorm:"column:dok_upload"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DokumenMaster MasterDokumenLegal     `gorm:"foreignKey:DokumenId;references:IdMasterDokumenLegal"`
    UMKM       UMKM      `gorm:"foreignKey:UmkmId;references:IdUmkm"`
}

func (UmkmDokumen) TableName() string {
    return "umkm_dokumen_legal"
}

