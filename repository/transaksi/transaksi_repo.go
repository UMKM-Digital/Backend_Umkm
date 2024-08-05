package transaksirepo

import "umkm/model/domain"

type TransaksiRepo interface {
	CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error)
}