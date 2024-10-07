package web

import (
	"encoding/json"

	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type WebProduk struct {
	UmkmId uuid.UUID `validate:"required" json:"umkm_id"`
	Name   string `validate:"required" json:"nama"`
	GambarId json.RawMessage `validate:"required" json:"gambar_id"`
	Harga int `validate:"required" json:"harga"`
	Satuan string `validate:"required" json:"satuan"`
	MinPesanan int `validate:"required" json:"min_pesanan"`
	Deskripsi string `validate:"required" json:"deskripsi"`
	KategoriProduk json.RawMessage `validate:"required" json:"kategori_produk_id"`
}


type UpdatedProduk struct {
	Name           string          `validate:"required" json:"nama"`
	GambarIDs      json.RawMessage `validate:"required" json:"gambar_id"`
	Harga          int             `validate:"required" json:"harga"`
	Satuan         string             `validate:"required" json:"satuan"`
	MinPesanan     int             `validate:"required" json:"min_pesanan"`
	Deskripsi      string          `validate:"required" json:"deskripsi"`
	KategoriProduk  json.RawMessage `validate:"required" json:"kategori_produk_id"`
	  Index       int      `json:"index"`
}

type UpdatePorudkActive struct{
	Active int    `json:"active"`
}
