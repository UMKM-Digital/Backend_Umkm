package hakaksesrepo

import "umkm/model/domain"

type CreateHakakses interface {
	CreateHakAkses(hakAkses *domain.HakAkses) error
}