package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/shopspring/decimal"
	  "strings"
	  "encoding/json"
)

type KategoriProduk struct{
	// Id string `json:"id"`
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
        var kategoriProduk domain.Kategori_Produk

        // Convert JSONB map to JSON bytes
        jsonBytes, err := json.Marshal(transaksi.IdKategoriProduk)
        if err != nil {
            return TransaksiEntity{}
        }

        // Unmarshal JSON bytes to domain.Kategori_Produk
        err = json.Unmarshal(jsonBytes, &kategoriProduk)
        if err != nil {
            return TransaksiEntity{}
        }
	
        // Convert to desired format
        // idArray := kategoriProduk.ID
        namaArray := strings.Split(kategoriProduk.Nama[0], ",")

        for i := range namaArray {
            if i < len(namaArray) {
                kategori = append(kategori, KategoriProduk{
                    // Id:   strings.TrimSpace(idArray[i]),
                    Nama: strings.TrimSpace(namaArray[i]),
                })
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
