package domain

import(
	"time"
)

type Kategori_Umkm struct {
    IdKtegori     int       `gorm:"column:id;primaryKey;autoIncrement"`
    Name       string    `gorm:"column:name"`
    Created_at time.Time `gorm:"column:created_at"`
    Updated_at time.Time `gorm:"column:updated_at"`
}

func (Kategori_Umkm) TableName() string {
    return "kategori_umkm"
}
