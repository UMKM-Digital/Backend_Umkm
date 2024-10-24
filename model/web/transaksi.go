	package web
	
	import (
		"github.com/google/uuid"
		// "github.com/shopspring/decimal"
		"encoding/json"
	)

	type CreateTransaksi struct {
		UmkmId         uuid.UUID `validate:"required" json:"umkm_id"`
		// No_Invoice     string `validate:"required" json:"no_invoice"`
		Tanggal        string `validate:"required" json:"tanggal"`
		NamaClient     string `validate:"required" json:"name_client"`
		IDKategoriProduk json.RawMessage `validate:"required" json:"id_kategori_produk"`
		TotalJml        float64  `validate:"required" json:"total_jml"`
		Keteranagan    string `validate:"required" json:"Keterangan"`
		Status         int    `validate:"required" json:"status"`
		NoHp         string    `json:"no_hp"`
		// TiketValidasi  string `validate:"required" json:"tiket_validasi"`
	}