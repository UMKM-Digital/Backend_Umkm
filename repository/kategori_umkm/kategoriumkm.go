package repokategoriumkm

import "umkm/model/domain"

type CreateCategoryUmkm interface {
	CreateRequest(categoryumkm domain.Kategori_Umkm) (domain.Kategori_Umkm, error)
	GetKategoriUmkm() ([]domain.Kategori_Umkm, error)
}