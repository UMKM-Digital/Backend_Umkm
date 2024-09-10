package entity

import (
	"time"
	domain "umkm/model/domain/homepage"
)

type SliderEntity struct {
	ID         int       `json:"id"`
	SlideTitle string    `json:"slide_title"`
	SlideDesc  string    `json:"slide_desc"`
	Active     int       `json:"active"`
	Gambar     string    `json:"gambar"`
	Created    time.Time `json:"created_at"`
	Update    time.Time `json:"updated_at"`
}

func ToSliderEntity(slider domain.Slider) SliderEntity {
	return SliderEntity{
		ID:          slider.Id,
		SlideTitle:        slider.SlideTitle,
		SlideDesc:      slider.SlideDesc,
		Active:      slider.Active,
		Gambar: slider.Gambar,
		Created:     slider.CreatedAt,
		Update: slider.UpdatedAt,
	}
}

func ToSliderEntities(slider []domain.Slider) []SliderEntity {
	var sliderEntities []SliderEntity
	for _, slider := range slider {
		sliderEntities = append(sliderEntities, ToSliderEntity(slider))
	}
	return sliderEntities
}