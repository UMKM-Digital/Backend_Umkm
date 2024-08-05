package transaksiservice

import "umkm/model/web"

type Transaksi interface {
	CreateTransaksi(umkm web.CreateTransaksi) (map[string]interface{}, error)
}
