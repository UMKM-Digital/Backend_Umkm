package testimonialrepo

import "umkm/model/domain/homepage"

type Testimonal interface {
	CreateTestimonial(testimonal domain.Testimonal)(domain.Testimonal, error)
	GetTestimonial()([]domain.Testimonal, error)
}