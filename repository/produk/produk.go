package produkrepo

import "umkm/model/domain"

type CreateProduk interface {
	CreateRequest(produk domain.Produk)(domain.Produk, error)
}