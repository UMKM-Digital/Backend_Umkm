package entity

import (
	"time"
	domain "umkm/model/domain/homepage"
)

type TesttimonialEntity struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Quotes string `json:"quote"`
	Active int `json:"active"`
	Created time.Time `json:"created_at"`
}

func ToTestimonialEntity(testimony domain.Testimonal) TesttimonialEntity {
	return TesttimonialEntity{
		ID: testimony.Id,
		Name: testimony.Name,
		Quotes: testimony.Quotes,
		Active: testimony.Active,
		Created: testimony.Created_at,
	}
}

func ToKategoriProdukEntities(testimonal []domain.Testimonal) []TesttimonialEntity {
	var testimonaliEntities []TesttimonialEntity
	for _, kategori := range testimonal {
		testimonaliEntities = append(testimonaliEntities, ToTestimonialEntity(kategori))
	}
	return testimonaliEntities
}