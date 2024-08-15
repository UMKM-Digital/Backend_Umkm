package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/shopspring/decimal"
)

type TransasksiFilterEntity struct {
	Id int `json:"id"`
	Name string `json:"name_client"`
	NoInvoice string `json:"no_invoice"`
	Tanggal time.Time `json:"tanggal"`
	TotalJml decimal.Decimal `json:"total_jml"`
	Status int `json:"status"`
}

func ToTransaksiListEntity(transaksi domain.Transaksi) TransasksiFilterEntity {
	return TransasksiFilterEntity{
		Id: transaksi.IdTransaksi,
		Name: transaksi.Nameclient,
		NoInvoice: transaksi.NoInvoice,
		Tanggal: transaksi.Tanggal,
		TotalJml: transaksi.TotalJml,
		Status: transaksi.Status,
	}
}

func ToTransaksiFilterEntities(TransaksiFilterList []domain.Transaksi) []TransasksiFilterEntity {
    var TransaksiFilterEntities []TransasksiFilterEntity
    for _, Transaksifilter := range TransaksiFilterList {
        TransaksiFilterEntities = append(TransaksiFilterEntities, ToTransaksiListEntity(Transaksifilter))
    }
    return TransaksiFilterEntities
}
