package aboutusrepo

import domain "umkm/model/domain/homepage"

type AboutUs interface {
	CreatedAboutUs(AboutUs domain.AboutUs) (domain.AboutUs, error)
	GetAboutUs() ([]domain.AboutUs, error)
	FindByAboutId(id int) (domain.AboutUs, error)
	UpdateAboutUsId(id int, AboutUs domain.AboutUs) (domain.AboutUs, error) 
}
