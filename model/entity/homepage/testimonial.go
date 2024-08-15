package entity

import domain "umkm/model/domain/homepage"

type TesttimonialEntity struct {
	
	Name string `json:"name"`
	Quotes string `json:"quote"`
}

func ToTestimonialEntity(testimony domain.Testimonal) TesttimonialEntity {
	return TesttimonialEntity{
		Name: testimony.Name,
		Quotes: testimony.Quotes,
	}
}

func ToKategoriProdukEntities(testimonal []domain.Testimonal) []TesttimonialEntity {
	var testimonaliEntities []TesttimonialEntity
	for _, kategori := range testimonal {
		testimonaliEntities = append(testimonaliEntities, ToTestimonialEntity(kategori))
	}
	return testimonaliEntities
}