package transaksirepo

import (
	"umkm/model/domain"

	"github.com/google/uuid"
	
)

type TransaksiRepo interface {
	CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error)
	GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error)
	GetFilterTransaksi(umkmID uuid.UUID, filters string, filterTanggal string, limit int, page int, status string) ([]domain.Transaksi, int, int, int, *int, *int, error)
	// GetFilterTransaksiWebTahun(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error)
	GetFilterTransaksiWebTahunByUserID(userID int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) 
	// GetTransaksiByMonth(umkmID uuid.UUID, year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error)
	GetTransaksiByMonth(userId int, umkmID uuid.UUID, year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) 
	GetTransaksiByDate(userId int, umkmID uuid.UUID, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error)
	DeleteTransaksiUmkmId(id uuid.UUID) error
}