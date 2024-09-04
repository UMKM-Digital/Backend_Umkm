package produkcontroller

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	// "umkm/helper"
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

func NewProdukController(Produk produkservice.Produk) *ProdukControllerImpl {
	return &ProdukControllerImpl{
		Produk: Produk,
	}
}

func (controller *ProdukControllerImpl) CreateProduk(c echo.Context) error {
	produk := new(web.WebProduk)

	// Konversi umkm_id dari string ke uuid.UUID
	umkmIDStr := c.FormValue("umkm_id")
	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UMKM ID format", nil))
	}

	produkHargastr := c.FormValue("harga")
	produkHarga, err := strconv.Atoi(produkHargastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest,false, "Invalid harga format", nil))
	}

	produkstauanstr := c.FormValue("satuan")
	produksatuan, err := strconv.Atoi(produkstauanstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid satuan format", nil))
	}

	produkminpesananstr := c.FormValue("min_pesanan")
	produkminpesanan, err := strconv.Atoi(produkminpesananstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid min_pesanan format", nil))
	}

	kategorid := c.FormValue("kategori_produk_id")
    if kategorid != "" {
        produk.KategoriProduk = json.RawMessage(kategorid)
    }


	produk.UmkmId = umkmID
	produk.Harga = produkHarga
	produk.Satuan = produksatuan
	produk.MinPesanan = produkminpesanan
	produk.Name = c.FormValue("nama")
	produk.Deskripsi = c.FormValue("deskripsi")

	// Handle image upload
	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "Failed to parse form", nil))
	}

	files := c.Request().MultipartForm.File["images"]
	fileHeaders := make(map[string]*multipart.FileHeader)
	for _, file := range files {
		fileHeaders[file.Filename] = file
	}

	produk.GambarId = json.RawMessage([]byte("[]")) // Default empty JSON array
	result, errSaveProduk := controller.Produk.CreateProduk(*produk, fileHeaders)
	if errSaveProduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "membut produk berhasil", result))
}

func (controller *ProdukControllerImpl) DeleteProdukId(c echo.Context) error {
	// Ambil ID dari URL dan konversi ke UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid ID format", nil))
	}

	if errDeleteProduk := controller.Produk.DeleteProduk(id); errDeleteProduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete Produk Success", nil))
}

// func (controller *ProdukControllerImpl) GetProdukId(c echo.Context) error{
// 	IdProduk := c.Param("id")
// 	id, err := uuid.Parse(IdProduk)

// 	if err != nil{
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "Id tidak ada", nil))
// 	}

// 	if errGetPorudk := controller.Produk.GetProdukId(id); errGetPorudk != nil{
// 		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errGetPorudk.Error(), nil))
// 	}
// 	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK,"id ditemukan", nil))
// }


func (controller *ProdukControllerImpl) GetprodukList(c echo.Context) error {
	produkIDStr := c.Param("umkm_id")
	kategoriProdukID := c.QueryParam("kategori")

	filters, limit, page := helper.ExtractFilter(c.QueryParams())

	if produkIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Produk ID cannot be empty")
	}

	produkId, err := uuid.Parse(produkIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Produk ID")
	}

	// Panggil service untuk mendapatkan daftar produk dan total count
	result, err := controller.Produk.GetProdukList(produkId, filters, limit, page, kategoriProdukID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get produk list")
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"message": "Berhasil melihat seluruh produk berdasarkan umkm_id",
		"data":    result,
	}

	return c.JSON(http.StatusOK, response)
}
