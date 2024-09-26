package domain

import "time"

type Brandlogo struct {
	Id         int       `gorm:"column:id;primaryKey;autoIncrement"`
	BrandName  string    `gorm:"column:brand_name"`
	BrandLogo  string    `gorm:"column:brand_logo"` // Remove extra space
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Brandlogo) TableName() string {
	return "homepage.brandlogo"
}
