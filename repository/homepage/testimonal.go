package testimonialrepo

import domain "umkm/model/domain/homepage"

type Testimonal interface {
	CreateTestimonial(testimonal domain.Testimonal) (domain.Testimonal, error)
	GetTestimonial() ([]domain.Testimonal, error)
	DelTransaksi(id int) error
	GetTransaksiByid(id int) (domain.Testimonal, error)
	UpdateTransaksiId(id int, kategori domain.Testimonal) (domain.Testimonal, error)
	GetTestimonialActive(active int)([]domain.Testimonal, error)
}