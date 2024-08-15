package domain

type Testimonal struct{
	Id int   `gorm:"column:id;primaryKey;autoIncrement"`
	Quotes string `gorm:"column:quote"`
	Name string 	`gorm:"column:name"`
}

func (Testimonal) TableName() string {
    return "homepage.testimonial"
}