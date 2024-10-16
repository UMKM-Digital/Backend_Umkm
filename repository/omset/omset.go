package omsetrepo

import "umkm/model/domain"

type OmsetRepo interface {
	CreateRequest(produk domain.Omset) (domain.Omset, error)
}