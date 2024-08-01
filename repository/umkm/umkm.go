package umkmrepo

import "umkm/model/domain"

type CreateUmkm interface {
	CreateRequest(umkm domain.UMKM)(domain.UMKM, error)
}