package produkcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	produkservice "umkm/service/produk"

	
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProdukControllerImpl struct {
	Produk produkservice.Produk
}

func NewKategeoriUmkmController(Produk produkservice.Produk) *ProdukControllerImpl {
	return &ProdukControllerImpl{
		Produk: Produk,
	}
}

type Produk struct {
    UmkmId uuid.UUID
}

func(controller *ProdukControllerImpl) CreateProduk(c echo.Context) error {

	produk := new(web.WebProduk)
   // Konversi umkm_id dari string ke uuid.UUID
	umkmIDStr := c.FormValue("umkm_id")
	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid UMKM ID format",
		})
	}
	//harga
	produkHargastr := c.FormValue("harga")
	produkHarga, err := strconv.Atoi(produkHargastr)
	//satuan
	produkstauanstr := c.FormValue("satuan")
	produksatuan, err := strconv.Atoi(produkstauanstr)
	//min_pesanan
	produkminpesananstr := c.FormValue("min_pesanan")
	produkminpesanan, err := strconv.Atoi(produkminpesananstr)
	produk.UmkmId = umkmID
	produk.Harga = produkHarga
	produk.Satuan = produksatuan
	produk.MinPesanan = produkminpesanan
	produk.Name = c.FormValue("nama")
	// produk.Harga = c.FormValue("harga")
	// produk.Satuan = c.FormValue("satuan")
	// produk.MinPesanan = c.FormValue("min_pesanan")
	produk.Deskripsi = c.FormValue("deskripsi")

	//uid
	
	//gambar
	var gambarURLs []string
    if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to parse form", nil))
    }

    files := c.Request().MultipartForm.File["images"]
    for _, file := range files {
        url, err := helper.HandleFileUpload(file, "uploads")
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to upload file", nil))
        }
        gambarURLs = append(gambarURLs, url)
    }

    gambarURLsJSON, err := json.Marshal(gambarURLs)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to marshal image URLs", nil))
    }
    produk.GambarId = json.RawMessage(gambarURLsJSON)

	//
	kategorid := c.FormValue("kategori_produk_id")
    if kategorid != "" {
        produk.KategoriProduk = json.RawMessage(kategorid)
    }

	result, errSaveProduk := controller.Produk.CreateProduk(*produk)
    if errSaveProduk != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveProduk.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create Produk Success", result))
}
	
	

