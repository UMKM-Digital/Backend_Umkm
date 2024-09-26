package domain

import (
	"time"

	"github.com/google/uuid"
)

type Produk struct {
	IdUmkm uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	UmkmId uuid.UUID `gorm:"column:umkm_id"`
	Nama string `gorm:"column:nama"`
	Gamabr   JSONB       `gorm:"column:gambar_id"`
	Harga int `gorm:"column:harga"`
	Satuan int `gorm:"column:satuan"`
	Min_pesanan int `gorm:"column:min_pesanan"`
	KategoriProduk JSONB `gorm:"column:kategori_produk_id"`
	Deskripsi string `gorm:"column:deskripsi"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Umkm    UMKM `gorm:"foreignKey:UmkmId"`
}

func (Produk) TableName() string {
    return "produk"
}