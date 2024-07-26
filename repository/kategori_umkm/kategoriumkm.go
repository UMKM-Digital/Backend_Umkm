package repokategoriumkm

import "umkm/model/domain"

type CreateCategoryUmkm interface {
	CreateRequest(categoryumkm domain.Kategori_Umkm) (domain.Kategori_Umkm, error)
	GetKategoriUmkm() ([]domain.Kategori_Umkm, error)
	GetKategoriUmkmId(idKategori int) (domain.Kategori_Umkm, error)
	UpdateKategoriId(idKategori int, kategori domain.Kategori_Umkm) (domain.Kategori_Umkm, error)
}