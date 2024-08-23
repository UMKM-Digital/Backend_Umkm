package brandrepo

import domain "umkm/model/domain/homepage"

type Brandlogo interface {
	CreatedBrandLogo(brandlogo domain.Brandlogo) (domain.Brandlogo, error)
	GetBrandLogo() ([]domain.Brandlogo, error)
	DeleteLogoId(id int) error
	FindById(id int) (domain.Brandlogo, error)
	UpdateBrandLogoId(id int, brandlogo domain.Brandlogo) (domain.Brandlogo, error) 
}
