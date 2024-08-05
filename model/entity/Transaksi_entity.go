package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/shopspring/decimal"
)

type TransaksiEntity struct {
	Id               int    `json:"id"`
	NoInvoice        string `json:no_invoice`
	Tanggal          time.Time `json:tanggal`
	Nameclient       string `json:name_client`
	IdKategoriProduk domain.JSONB ` json:"id_kategori_produk"`
	TotalBelanja decimal.Decimal `json:total_jml`
}
func ToTransaksiEntity(transaksi domain.Transaksi) TransaksiEntity {
	return TransaksiEntity{
		Id: transaksi.IdTransaksi,
		NoInvoice: transaksi.NoInvoice,
		Tanggal: transaksi.Tanggal,
		Nameclient: transaksi.Nameclient,
		IdKategoriProduk: transaksi.IdKategoriProduk,
		TotalBelanja: transaksi.TotalJml,
	}
}

func ToTransaksiEntities(transaksilist []domain.Transaksi) []TransaksiEntity {
    var kategoriEntities []TransaksiEntity
    for _, kategori := range transaksilist {
        kategoriEntities = append(kategoriEntities, ToTransaksiEntity(kategori))
    }
    return kategoriEntities
}