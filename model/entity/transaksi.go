package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/shopspring/decimal"
	  "strings"
	
)

type KategoriProduk struct{
	Nama string `json:"nama"`
}

type TransaksiEntity struct {
	Id               int             `json:"id"`
	NoInvoice        string          `json:"no_invoice"`
	Tanggal          time.Time       `json:"tanggal"`
	Nameclient       string          `json:"name_client"`
	IdKategoriProduk []KategoriProduk    ` json:"id_kategori_produk"`
	TotalBelanja     decimal.Decimal `json:"total_jml"`
	 TiketValidasi  string    `json:"tiket_validasi"`

}

func ToTransaksiEntity(transaksi domain.Transaksi) TransaksiEntity {
	var kategori []KategoriProduk

	
	if transaksi.IdKategoriProduk != nil {
      

        // Assuming IdKategoriProduk["nama"] is a []interface{} or []string
        if namaArray, ok := transaksi.IdKategoriProduk["nama"].([]interface{}); ok {
            for _, nama := range namaArray {
                if namaStr, ok := nama.(string); ok {
                    kategori = append(kategori, KategoriProduk{
                        Nama: strings.TrimSpace(namaStr),
                    })
                } 
            }
        } 
    }

	return TransaksiEntity{
		Id:               transaksi.IdTransaksi,
		NoInvoice:        transaksi.NoInvoice,
		Tanggal:          transaksi.Tanggal,
		Nameclient:       transaksi.Nameclient,
		IdKategoriProduk: kategori,
		TotalBelanja:     transaksi.TotalJml,
		TiketValidasi: transaksi.TiketValidasi,
	}
}

// func ToTransaksiEntities(transaksilist []domain.Transaksi) []TransaksiEntity {
// 	var kategoriEntities []TransaksiEntity
// 	for _, kategori := range transaksilist {
// 		kategoriEntities = append(kategoriEntities, ToTransaksiEntity(kategori))
// 	}
// 	return kategoriEntities
// }
