	package brandlogoservice

	import (
		"mime/multipart"
		entity "umkm/model/entity/homepage/brandlogo"
		web "umkm/model/web/homepage"
	)

	type Brandlogo interface {
		CreateBrandlogo(brandlogo web.CreatedBrandLogo, file *multipart.FileHeader) (map[string]interface{}, error)
		GetBrandlogoList() ([]entity.BrandLogoEntity, error)
		// DeleteTestimonial(id int) error
		// GetTestimonialid(id int) (entity.TesttimonialEntity, error)
		// UpdateTestimonial(request web.UpdateTestimonial, Id int) (map[string]interface{}, error)
	}