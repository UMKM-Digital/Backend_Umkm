package aboutusservice

import (
	"mime/multipart"
	entity "umkm/model/entity/homepage/brandlogo"
	"umkm/model/web/homepage"
)

type AboutUs interface {
	CreateAboutUs(testimonal web.CreateAboutUs, file *multipart.FileHeader) (map[string]interface{}, error)
	GetAboutUs() ([]entity.AboutUsEntity, error)
	GetAboutUsid(id int) (entity.AboutUsEntity, error)
	UpdateAboutUs(request web.UpdateAboutUs, Id int, file *multipart.FileHeader) (map[string]interface{}, error)
}