package entity

import "umkm/model/domain"

type KategoriProdukEntity struct {
	Id   int    `json:"key"`
	Name string `json:"value"`
}

func ToKategoriProdukEntity(kategoriproduk domain.KategoriProduk) KategoriProdukEntity {
	return KategoriProdukEntity{
		Id: kategoriproduk.IdProduk,
		Name: kategoriproduk.Nama,
	}
}

func ToKategoriProdukEntities(kategoriList []domain.KategoriProduk) []KategoriProdukEntity {
    var kategoriEntities []KategoriProdukEntity
    for _, kategori := range kategoriList {
        kategoriEntities = append(kategoriEntities, ToKategoriProdukEntity(kategori))
    }
    return kategoriEntities
}