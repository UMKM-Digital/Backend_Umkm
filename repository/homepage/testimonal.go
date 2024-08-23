package testimonialrepo

import domain "umkm/model/domain/homepage"

type Testimonal interface {
	CreateTestimonial(testimonal domain.Testimonal) (domain.Testimonal, error)
	GetTestimonial() ([]domain.Testimonal, error)
	DelTestimonial(id int) error
	GetTransaksiByid(id int) (domain.Testimonal, error)
	UpdateTestimonialId(id int, testimonal domain.Testimonal) (domain.Testimonal, error)
	GetTestimonialActive(active int)([]domain.Testimonal, error)
	UpdateActiveId(idTestimonial int, active int) error
}