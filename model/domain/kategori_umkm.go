package domain

import(
	"time"
)

type Kategori_Umkm struct {
    IdKategori     int       `gorm:"column:id;primaryKey;autoIncrement"`
    Name       string    `gorm:"column:name"`
    CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Kategori_Umkm) TableName() string {
    return "kategori_umkm"
}
