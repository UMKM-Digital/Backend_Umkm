package transaksirepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
	
)

type TransaksiRepo interface {
	CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error)
	GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error)
	GetFilterTransaksi(umkmID uuid.UUID) ([]domain.Transaksi, error)
	GetFilterTransaksiWebTahun(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error)
	GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, error) 
	GetTransaksiByDate(umkmID string, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, error) 
}