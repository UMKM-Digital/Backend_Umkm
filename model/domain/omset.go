package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Omset struct {
	IdOmste   int       `gorm:"column:id;primaryKey;autoIncrement"`
	UmkmId    uuid.UUID `gorm:"column:umkm_id"`
	Tahun 	   string   `gorm:"column:tahun"`
	Bulan      string   `gorm:"column:bulan"`
	JumlahOmset   decimal.Decimal       `gorm:"column:jumlah_omset"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Umkm    UMKM `gorm:"foreignKey:UmkmId"`
}

func (Omset) TableName() string {
	return "omset"
}