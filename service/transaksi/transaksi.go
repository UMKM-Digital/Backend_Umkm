package transaksiservice

import (
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type Transaksi interface {
	CreateTransaksi(umkm web.CreateTransaksi) (map[string]interface{}, error)
	GetKategoriUmkmId(id int)(entity.TransaksiEntity, error)
	GetTransaksiFilter(umkmID uuid.UUID, filtersTanggal map[string]string, allowedfiltersTanggal []string, filters string, limit int, page int, status string) ([]entity.TransasksiFilterEntity, int, int, int, *int, *int, error)
	GetTransaksiByYear(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error) 
	GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error)
	GetTransaksiByDate(umkmID uuid.UUID, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error)
}
