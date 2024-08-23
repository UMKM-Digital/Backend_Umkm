package domain

import "time"

type Testimonal struct {
	Id         int    `gorm:"column:id;primaryKey;autoIncrement"`
	Quotes     string `gorm:"column:quote"`
	Name       string `gorm:"column:name"`
	Active     int `gorm:"column:active"`
	GambarTesti string `gorm:"column:gambar_testi"`
	Created_at time.Time`gorm:"column:created_at"`
}

func (Testimonal) TableName() string {
	return "homepage.testimonial"
}