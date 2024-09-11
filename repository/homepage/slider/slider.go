package sliderrepo

import domain "umkm/model/domain/homepage"

type Slider interface {
	Created(slider domain.Slider) (domain.Slider, error)
	GetSlider() ([]domain.Slider, error)
	GetSliderId(id int) (domain.Slider, error)
	DelSlider(id int) error
	UpdateSliderId(id int, slider domain.Slider) (domain.Slider, error)
	UpdateActiveId(idSlider int, active int) error
	GetSliderActive(active int)([]domain.Slider, error)
}