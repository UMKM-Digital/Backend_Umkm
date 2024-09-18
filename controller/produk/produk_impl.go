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

func (controller *ProdukControllerImpl) GetProdukId(c echo.Context) error{
	IdProduk := c.Param("id")
	id, _ := uuid.Parse(IdProduk)

	getProduk, errGetProduk := controller.Produk.GetProdukId(id)

	if errGetProduk != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id transaksi", getProduk))
}


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

	// Call service to get produk list and pagination data
	result, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.Produk.GetProdukList(produkId, filters, limit, page, kategoriProdukID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	// Prepare pagination data
	pagination := model.Pagination{
		CurrentPage:  currentPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalPages:   totalPages,
		TotalRecords: totalCount,
	}

	// Create response
	response := model.ResponseToClientpagi(http.StatusOK, "true", "Berhasil melihat seluruh produk berdasarkan umkm_id", pagination, result)

	return c.JSON(http.StatusOK, response)
}

func (controller *ProdukControllerImpl) UpdateProduk(c echo.Context) error {
    IdProduk := c.Param("id")
    id, err := uuid.Parse(IdProduk)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UUID format", nil))
    }

    name := c.FormValue("nama")
    deskripsistr := c.FormValue("deskripsi")

    hargastr := c.FormValue("harga")
    harga, err := strconv.Atoi(hargastr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid harga format", nil))
    }

    satuanstr := c.FormValue("satuan")
    satuan, err := strconv.Atoi(satuanstr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid satuan format", nil))
    }

    minpesananstr := c.FormValue("minpesanan")
    minpesanan, err := strconv.Atoi(minpesananstr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid minpesanan format", nil))
    }

    // Ambil file gambar jika ada
    file, err := c.FormFile("gambar")
    if err != nil && err != http.ErrMissingFile {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get uploaded file", nil))
    }

    // Ambil kategori produk dari form value
    kategoriProdukJSON := c.FormValue("kategori_produk")

    // Ambil ID gambar yang akan dihapus
    indexHapusStr := c.FormValue("index_hapus")
    var indexHapus []string
    if err := json.Unmarshal([]byte(indexHapusStr), &indexHapus); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid index_hapus format", nil))
    }

    // Buat request untuk service
    request := web.UpdatedProduk{
        Name:           name,
        Harga:          harga,
        Satuan:         satuan,
        MinPesanan:     minpesanan,
        Deskripsi:      deskripsistr,
        KategoriProduk: json.RawMessage(kategoriProdukJSON),
    }

    // Panggil service untuk memperbarui produk
    updatedProduk, errUpdate := controller.Produk.UpdateProduk(request, id, file, indexHapus)
    if errUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Data berhasil diupdate", updatedProduk))
}
