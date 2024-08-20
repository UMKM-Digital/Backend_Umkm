package domain

import "time"

type Brandlogo struct {
	Id         int       `gorm:"column:id;primaryKey;autoIncrement"`
	BrandName  string    `gorm:"column:brand_name"`
	BrandLogo  string    `gorm:"column:brand_logo"` // Remove extra space
	Created_at time.Time `gorm:"column:created_at"`
}

func (Brandlogo) TableName() string {
	return "homepage.brandlogo"
}
