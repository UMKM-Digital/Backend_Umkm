package domain

import (
    "github.com/google/uuid"
    "time"
)

// Definisikan tipe untuk enum status
type StatusEnum string

const (
    Menunggu   StatusEnum = "menunggu"
    Disetujui  StatusEnum = "disetujui"
    Ditolak    StatusEnum = "ditolak"
)

type HakAkses struct {
    IdHakakses     int         `gorm:"column:id;primaryKey;autoIncrement"`
    UserId         int         `gorm:"column:user_id"`
    UmkmId         uuid.UUID   `gorm:"column:umkm_id;type:uuid"`
    Status         StatusEnum  `gorm:"column:status;type:enum"` // Ubah ke tipe enum
    Pembina         string  `gorm:"column:pembina"` // Ubah ke tipe enum
    Pesan         string  `gorm:"column:pesan"` // Ubah ke tipe enum
    CreatedAt      time.Time   `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt      time.Time   `gorm:"column:updated_at;autoUpdateTime"`
    User           Users       `gorm:"foreignKey:UserId;references:IdUser"`
    UMKM           UMKM        `gorm:"foreignKey:UmkmId;references:IdUmkm"`
}

func (HakAkses) TableName() string {
    return "hak_akses"
}
