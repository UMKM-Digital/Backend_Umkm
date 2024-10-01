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
	Satuan string `json:"satuan"`
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
	 Created: produk.CreatedAt,
	 Update: produk.UpdatedAt,
	}
}


type ProdukList struct{
	IdProduk uuid.UUID `json:"id"`
	Name string `json:"nama"`
	Images domain.JSONB `json:"gambar_id"`
	KategdoriProduk domain.JSONB `json:"kategori_produk_id"`
	Harga int `json:"harga"`
}

func ToProdukList(produk domain.Produk) ProdukList{
	return ProdukList{
		IdProduk: produk.IdUmkm,
		Name: produk.Nama,
		Images: produk.Gamabr,
		KategdoriProduk: produk.KategoriProduk,
		Harga: produk.Harga,
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
	KategoriProduk domain.JSONB `json:"kategori_produk"`
}

func ToProdukWebEntity(produk domain.Produk) ProdukWebEntity {
	
	return ProdukWebEntity{
		IdProduk: produk.IdUmkm,
		Gambar: produk.Gamabr,
		Name:     produk.Nama,
		Harga: produk.Harga,
		NameUmkm: produk.Umkm.Name, // Assuming the Umkm relationship is populated
		KategoriProduk: produk.KategoriProduk,
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
	Satuan string `json:"satuan"`
	MinPesanan int `json:"min_pesanan"`
	Deskripsi string `json:"deskripsi"`
	NoKontak             string    `json:"no_kontak"`
	IdUmkm uuid.UUID `json:"id_umkm"`
	DeskripsiUmkm string `json:"deskripsi_umkm"`
	GambarUmkm domain.JSONB `json:"gambar_umkm"`
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
		DeskripsiUmkm: produk.Umkm.Deskripsi,
		GambarUmkm: produk.Umkm.Images,
	}
}

func ToProdukWebIdEntities(produkList []domain.Produk) []ProdukWebIdEntity {
	var produkEntities []ProdukWebIdEntity
	for _, produk := range produkList {
		produkEntities = append(produkEntities, ToProdukWebIdEntity(produk))
	}
	return produkEntities
}


// //detailpunya fadlan

// type ProdukEntityDetailMobile struct{
// 	Id uuid.UUID `json:"id"`
// 	Nama string `json:"nama"`
// 	Gambar domain.JSONB `json:"gambar_id"`
// 	Harga int `json:"harga"`
// 	Deskripsi string `json:"deskripsi"`
// 	KategdoriProduk domain.JSONB `json:"kategori_produk_id"`
// }


// func ToProdukIdEntity(produk domain.Produk) ProdukEntityDetailMobile {
	
// 	return ProdukEntityDetailMobile{
// 		Id: produk.IdUmkm,
// 		Gambar: produk.Gamabr,
// 		Nama:     produk.Nama,
// 		Harga: produk.Harga, // Assuming the Umkm relationship is populated
// 		Deskripsi: produk.Deskripsi,
// 		KategdoriProduk: produk.KategoriProduk,
// 	}
// }

// func ToProdukIdEntities(produkList []domain.Produk) []ProdukEntityDetailMobile {
// 	var produkEntities []ProdukEntityDetailMobile
// 	for _, produk := range produkList {
// 		produkEntities = append(produkEntities, ToProdukIdEntity(produk))
// 	}
// 	return produkEntities
// }