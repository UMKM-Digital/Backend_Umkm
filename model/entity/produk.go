package entity

import (
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
	}
}

func ToProdukEntities(produklist []domain.Produk) []ProdukEntity {
    var produkEntities []ProdukEntity
    for _, produk := range produklist {
        produkEntities = append(produkEntities, ToProdukEntity(produk))
    }
    return produkEntities
}
