package transaksiservice

import (
	"umkm/model/entity"
	"umkm/model/web"
)

type Transaksi interface {
	CreateTransaksi(umkm web.CreateTransaksi) (map[string]interface{}, error)
	GetKategoriUmkmId(id int)(entity.TransaksiEntity, error)
}
