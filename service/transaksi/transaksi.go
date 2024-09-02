package transaksiservice

import (
	"umkm/model/entity"
	"umkm/model/web"

	"github.com/google/uuid"
)

type Transaksi interface {
	CreateTransaksi(umkm web.CreateTransaksi) (map[string]interface{}, error)
	GetKategoriUmkmId(id int)(entity.TransaksiEntity, error)
	GetTransaksiFilter(umkmID uuid.UUID, filtersTanggal map[string]string, allowedfiltersTanggal []string, filters string, limit int, page int, status int) (*TransaksiFilterResult, error)
	GetTransaksiByYear(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error) 
	GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, error)
	GetTransaksiByDate(umkmID string, year int, month int, page int, limit int, filter string) (map[string]interface{}, error)
}
