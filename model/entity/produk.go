package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type ProdukEntity struct {
	IdProduk uuid.UUID `json:"id"`
	Name string `json:"nama"`
	Images domain.JSONB `json:"gambar_id"`
	Harga int `json:"harga"`
	KategdoriProduk domain.JSONB `json:"kategori_produk_id"`
	Satuan int `json:"satuan"`
	MinPesanan int `json:"min_pesanan"`
	Deskripsi string `json:"deskripsi"`
	Created time.Time `json:"created_at"`
	Update time.Time `josn:"updated_at"`
}

func ToProdukEntity(produk domain.Produk) ProdukEntity {
	return ProdukEntity{
	 IdProduk: produk.IdUmkm,
	 Name: produk.Nama,
	 Images: produk.Gamabr,
	 Deskripsi:  produk.Deskripsi,
	 Harga: produk.Harga,
	 Satuan: produk.Satuan,
	 MinPesanan: produk.Min_pesanan,
	 KategdoriProduk: produk.KategoriProduk,
	 Created: produk.Created_at,
	 Update: produk.Updated_at,
	}
}


type ProdukList struct{
	IdProduk uuid.UUID `json:"id"`
	Name string `json:"nama"`
	Images domain.JSONB `json:"gambar_id"`
	KategdoriProduk domain.JSONB `json:"kategori_produk_id"`
}

func ToProdukList(produk domain.Produk) ProdukList{
	return ProdukList{
		IdProduk: produk.IdUmkm,
		Name: produk.Nama,
		Images: produk.Gamabr,
		KategdoriProduk: produk.KategoriProduk,
	}
}

func ToProdukEntities(produklist []domain.Produk) []ProdukList {
    var produkEntities []ProdukList
    for _, produk := range produklist {
        produkEntities = append(produkEntities, ToProdukList(produk))
    }
    return produkEntities
}


type ProdukWebEntity struct {
	IdProduk   uuid.UUID `json:"id"`
	Gambar domain.JSONB `json:"gambar_id"`
	Name       string    `json:"nama"`
	Harga int `json:"harga"`
	NameUmkm   string    `json:"name"`
}

func ToProdukWebEntity(produk domain.Produk) ProdukWebEntity {
	
	return ProdukWebEntity{
		IdProduk: produk.IdUmkm,
		Gambar: produk.Gamabr,
		Name:     produk.Nama,
		Harga: produk.Harga,
		NameUmkm: produk.Umkm.Name, // Assuming the Umkm relationship is populated
	}
}

func ToProdukWebEntities(produkList []domain.Produk) []ProdukWebEntity {
	var produkEntities []ProdukWebEntity
	for _, produk := range produkList {
		produkEntities = append(produkEntities, ToProdukWebEntity(produk))
	}
	return produkEntities
}




type ProdukWebIdEntity struct {
	IdProduk   uuid.UUID `json:"id"`
	Gambar domain.JSONB `json:"gambar_id"`
	Name       string    `json:"nama"`
	KategdoriProduk domain.JSONB `json:"kategori_produk_id"`
	Harga int `json:"harga"`
	NameUmkm   string    `json:"name"`
	Satuan int `json:"satuan"`
	MinPesanan int `json:"min_pesanan"`
	Deskripsi string `json:"deskripsi"`
	NoKontak             string    `json:"no_kontak"`
	IdUmkm uuid.UUID `json:"id_umkm"`
}

func ToProdukWebIdEntity(produk domain.Produk) ProdukWebIdEntity {
	
	return ProdukWebIdEntity{
		IdProduk: produk.IdUmkm,
		Gambar: produk.Gamabr,
		Name:     produk.Nama,
		Harga: produk.Harga, // Assuming the Umkm relationship is populated
		NoKontak: produk.Umkm.NoKontak,
		Satuan: produk.Satuan,
		MinPesanan: produk.Min_pesanan,
		Deskripsi: produk.Deskripsi,
		KategdoriProduk: produk.KategoriProduk,
		NameUmkm: produk.Umkm.Name,
		IdUmkm: produk.Umkm.IdUmkm,
	}
}

func ToProdukWebIdEntities(produkList []domain.Produk) []ProdukWebIdEntity {
	var produkEntities []ProdukWebIdEntity
	for _, produk := range produkList {
		produkEntities = append(produkEntities, ToProdukWebIdEntity(produk))
	}
	return produkEntities
}