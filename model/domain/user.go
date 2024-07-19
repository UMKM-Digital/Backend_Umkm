package domain

import "time"

type Users struct {
    IdUser     int       `gorm:"column:id;primaryKey;autoIncrement"`
    Username   string    `gorm:"column:username"`
    Email      string    `gorm:"column:email"`
    Password   string    `gorm:"column:password"`
	Role       string    `gorm:"column:role"`
    Created_at time.Time `gorm:"column:created_at"`
    Updated_at time.Time `gorm:"column:updated_at"`
}

func (Users) TableName() string {
    return "users"
}
