	package brandlogoservice

	import (
		"mime/multipart"
		entity "umkm/model/entity/homepage/brandlogo"
		web "umkm/model/web/homepage"
	)

	type Brandlogo interface {
		CreateBrandlogo(brandlogo web.CreatedBrandLogo, file *multipart.FileHeader) (map[string]interface{}, error)
		GetBrandlogoList() ([]entity.BrandLogoEntity, error)
		DeleteBrandLogo(id int) error
		GetBrandLogoid(id int) (entity.BrandLogoEntity, error)
	}