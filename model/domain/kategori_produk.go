package domain

import (
    "time"
)

type KategoriProduk struct {
    IdProduk    int       `gorm:"column:id;primaryKey;autoIncrement"`
    Nama        string    `gorm:"column:nama"`
    CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (KategoriProduk) TableName() string {
    return "kategori_produk"
}
