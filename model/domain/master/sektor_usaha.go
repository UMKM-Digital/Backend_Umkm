package domain

import (
	"time"

)

type SektorUsaha struct {
	Id int       `gorm:"column:id;primaryKey;autoIncrement"`
	Nama string `gorm:"column:urai"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"` // Auto-fill saat record dibuat
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (SektorUsaha) TableName() string {
    return "master.sektor_usaha"
}

