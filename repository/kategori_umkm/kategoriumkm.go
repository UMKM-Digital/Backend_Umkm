package repokategoriumkm

import "umkm/model/domain"

type CreateCategoryUmkm interface {
	CreateRequest(categoryumkm domain.Kategori_Umkm) (domain.Kategori_Umkm, error)
	GetKategoriUmkm(filters string, limit int, page int) ([]domain.Kategori_Umkm, int, int, int, *int, *int, error)
	GetKategoriUmkmId(idKategori int) (domain.Kategori_Umkm, error)
	UpdateKategoriId(idKategori int, kategori domain.Kategori_Umkm) (domain.Kategori_Umkm, error)
	DeleteKategoriUmkmId(id int) error
	GetKategoriUmkmBySektor(id int)([]domain.Kategori_Umkm, error)
}