package produkservice

import (
	"encoding/json"
	"errors"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/web"
	produkrepo "umkm/repository/produk"
)

type ProdukServiceImpl struct {
	produkrepository produkrepo.CreateProduk
}

func NewUmkmService(produkrepository produkrepo.CreateProduk) *ProdukServiceImpl {
	return &ProdukServiceImpl{
		produkrepository: produkrepository,
	}
}

func(service *ProdukServiceImpl) CreateProduk(produk web.WebProduk) (map[string]interface{}, error){
	var Gambar domain.JSONB
    if len(produk.GambarId) > 0 {
        var imgURLs []string
        if err := json.Unmarshal(produk.GambarId, &imgURLs); err != nil {
            return nil, errors.New("invalid type for Images")
        }
        Gambar = domain.JSONB{"urls": imgURLs}
    } else {
        Gambar = domain.JSONB{"urls": []string{}}
    }

	KategoriProduk, err := helper.RawMessageToJSONB(produk.KategoriProduk)

    if err != nil {
        return nil, errors.New("invalid type for kategoriproduk")
    }

	newProduk := domain.Produk{
     UmkmId: produk.UmkmId,
	 Nama: produk.Name,
	 Gamabr: Gambar,
	 Harga: produk.Harga,
	 Satuan: produk.Satuan,
	 Min_pesanan: produk.MinPesanan,
	 KategoriProduk: KategoriProduk,
	 Deskripsi: produk.Deskripsi,
    }

	saveProduk, errSaveProduk := service.produkrepository.CreateRequest(newProduk)
    if errSaveProduk != nil {
        return nil, errSaveProduk
    }

	return map[string]interface{}{
		"umkm_id": saveProduk.UmkmId,
        "nama":                  saveProduk.Nama,
		"gambar": saveProduk.Gamabr,
        "harga":         saveProduk.Harga,
        "satuan":             saveProduk.Harga,
        "minimal_pesanan":                saveProduk.Satuan,
		"kategori_produk_id":        saveProduk.KategoriProduk,
		"deskripsi":                 saveProduk.Deskripsi,
	
    }, nil
}