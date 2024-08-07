package transaksirepo

import (
	"umkm/model/domain"
)

type TransaksiRepo interface {
	CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error)
	GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error)
	// GetRequestTransaksiByTahun(tanggal time.Time)(domain.Transaksi, error)
}