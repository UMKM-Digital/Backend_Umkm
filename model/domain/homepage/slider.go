package domain

import "time"

type Slider struct {
	Id         int       `gorm:"column:id;primaryKey;autoIncrement"`
	SlideDesc  string    `gorm:"column:slide_desc"`
	SlideTitle string    `gorm:"column:slide_title"`
	Active     int `gorm:"column:active"`
	Gambar string `gorm:"column:gambar"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"` // Auto-fill saat record dibuat
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func(Slider) TableName() string{
	return "homepage.sleder"
}