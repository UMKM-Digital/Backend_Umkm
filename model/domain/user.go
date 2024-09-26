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
   CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime"`
    HakAkses   []HakAkses `gorm:"foreignKey:user_id"`
    Berita   []Berita `gorm:"foreignKey:author"`
}

func (Users) TableName() string {
    return "users"
}
