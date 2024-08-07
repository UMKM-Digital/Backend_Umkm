package domain

import "time"

type Users struct {
    IdUser     int       `gorm:"column:id;primaryKey;autoIncrement"`
    Username   string    `gorm:"column:username"`
    Email      string    `gorm:"column:email"`
    Password   string    `gorm:"column:password"`
	Role       string    `gorm:"column:role"`
	No_Phone   string    `gorm:"column:no_phone"`
	Picture    string    `gorm:"column:picture"`
    Created_at time.Time `gorm:"column:created_at"`
    Updated_at time.Time `gorm:"column:updated_at"`
    HakAkses   []HakAkses `gorm:"foreignKey:user_id"`
}

func (Users) TableName() string {
    return "users"
}
