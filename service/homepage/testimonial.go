package homepageservice

import (
	entity "umkm/model/entity/homepage"
	"umkm/model/web/homepage"
)

type Testimonal interface {
	CreateTestimonial(testimonal web.CreateTestimonial) (map[string]interface{}, error)
	GetTestimonial() ([]entity.TesttimonialEntity, error)
	DeleteTestimonial (id int) error
}