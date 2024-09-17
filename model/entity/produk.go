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
