package domain

import "time"

type MasterDokumenLegal struct {
	IdMasterDokumenLegal int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name                 string    `gorm:"column:nama"`
	Iswajib int `gorm:"column:is_wajib"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Dokumen            []UmkmDokumen `gorm:"foreignkey:DokumenId"`
}

func (MasterDokumenLegal) TableName() string {
    return "master_dokumen_legal"
}