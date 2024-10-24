package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Omset struct {
	IdOmste   int       `gorm:"column:id;primaryKey;autoIncrement"`
	UmkmId    uuid.UUID `gorm:"column:umkm_id"`
	Bulan      string   `gorm:"column:bulan"`
	Nominal   decimal.Decimal       `gorm:"column:nominal"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Umkm    UMKM `gorm:"foreignKey:UmkmId"`
}

func (Omset) TableName() string {
	return "omzets"
}