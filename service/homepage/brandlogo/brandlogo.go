package brandlogoservice

import (
	"mime/multipart"
	web "umkm/model/web/homepage"
)

type Brandlogo interface {
	CreateBrandlogo(brandlogo web.CreatedBrandLogo, file *multipart.FileHeader) (map[string]interface{}, error)
	// GetTestimonial() ([]entity.TesttimonialEntity, error)
	// DeleteTestimonial(id int) error
	// GetTestimonialid(id int) (entity.TesttimonialEntity, error)
	// UpdateTestimonial(request web.UpdateTestimonial, Id int) (map[string]interface{}, error)
}