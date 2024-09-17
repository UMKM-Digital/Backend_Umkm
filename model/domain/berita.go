package domain

import (
	"time"
)

type Berita struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string    `gorm:"column:title"`
	Image     string    `gorm:"column:image"` // Remove extra space
	Content   string    `gorm:"column:content"`
	AuthorId  int       `gorm:"column:author"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"` // Auto-fill saat record dibuat
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	User      Users     `gorm:"foreignKey:AuthorId;references:IdUser"`
}

func (Berita) TableName() string {
	return "homepage.berita"
}
