package homepageservice

import (
	entity "umkm/model/entity/homepage"
	"umkm/model/web/homepage"
)

type Testimonal interface {
	CreateTestimonial(testimonal web.CreateTestimonial) (map[string]interface{}, error)
	GetTestimonial() ([]entity.TesttimonialEntity, error)
	DeleteTestimonial(id int) error
	GetTestimonialid(id int) (entity.TesttimonialEntity, error)
	UpdateTestimonial(request web.UpdateTestimonial, Id int) (map[string]interface{}, error)
	GetTestimonialActive() ([]entity.TesttimonialEntity, error)
}