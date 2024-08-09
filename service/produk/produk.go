package produkservice

import "umkm/model/web"

type Produk interface {
	CreateProduk(produk web.WebProduk) (map[string]interface{}, error)
}