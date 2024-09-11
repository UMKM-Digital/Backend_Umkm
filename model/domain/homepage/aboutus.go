package domain

import "time"

type AboutUs struct {
	Id         int       `gorm:"column:id;primaryKey;autoIncrement"`
	Image  string    `gorm:"column:image"`
	Description  string    `gorm:"column:description"` // Remove extra space
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (AboutUs) TableName() string {
	return "homepage.about"
}
