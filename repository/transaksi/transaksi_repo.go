package transaksirepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
)

type TransaksiRepo interface {
	CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error)
	GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error)
	GetFilterTransaksi(umkmID uuid.UUID) ([]domain.Transaksi, error)
	GetFilterTransaksiWebTahun(umkmId uuid.UUID, year int) ([]domain.Transaksi, error)
}