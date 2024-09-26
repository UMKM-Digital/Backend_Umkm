package domain

import (
    "time"

    "github.com/google/uuid"
)

type KategoriProduk struct {
    IdProduk    int       `gorm:"column:id;primaryKey;autoIncrement"`
    Umkm        uuid.UUID `gorm:"column:umkm_id"`
    Nama        string    `gorm:"column:nama"`
    CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
    UMKM        UMKM      `gorm:"foreignKey:umkm_id"` // Ensure UMKM is a struct or interface type
}

func (KategoriProduk) TableName() string {
    return "kategori_produk"
}
