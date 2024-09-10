package sliderservice

import (
	"mime/multipart"
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
)

type Slider interface {
	CreateSlider(slider web.CreatedSlider, file *multipart.FileHeader) (map[string]interface{}, error)
	GetSlider() ([]entity.SliderEntity, error)
	GetSliderid(id int) (entity.SliderEntity, error)
	DeleteId(id int) error
	UpdateTestimonial(request web.UpdateSlider, Id int, file *multipart.FileHeader) (map[string]interface{}, error)
}
