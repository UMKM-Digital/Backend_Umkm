package domain

import (
    "github.com/google/uuid"
    "time"
)

type HakAkses struct {
    IdHakakses     int       `gorm:"column:id;primaryKey;autoIncrement"`
    UserId         int       `gorm:"column:user_id"`
    UmkmId         uuid.UUID `gorm:"column:umkm_id;type:uuid"`
    Status         int       `gorm:"column:status"`
    CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
	User       Users     `gorm:"foreignKey:UserId;references:IdUser"`
    UMKM       UMKM      `gorm:"foreignKey:UmkmId;references:IdUmkm"`
}

func (HakAkses) TableName() string {
    return "hak_akses"
}