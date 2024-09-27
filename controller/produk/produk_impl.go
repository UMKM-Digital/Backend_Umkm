package produkcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	// "umkm/helper"
	"umkm/helper"
	"umkm/model"
	"umkm/model/entity"
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

	files := c.Request().MultipartForm.File["gambar"]
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

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id produk", getProduk))
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

	if result == nil {
		result = []entity.ProdukList{}
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
    // Ambil ID produk dari URL dan parsing
    IdProduk := c.Param("id")
    id, err := uuid.Parse(IdProduk)
    if err != nil {
        log.Printf("Error parsing UUID: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UUID format", nil))
    }
    log.Printf("Received request to update product with ID: %s", id)

    // Proses data produk lainnya
    name := c.FormValue("nama")
    deskripsistr := c.FormValue("deskripsi")
	kategoriProdukID := c.FormValue("kategori_umkm_id")
    hargastr := c.FormValue("harga")
    harga, err := strconv.Atoi(hargastr)
    if err != nil {
        log.Printf("Error converting harga: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid harga format", nil))
    }
    log.Printf("Parsed harga: %d", harga)

    satuanstr := c.FormValue("satuan")
    satuan, err := strconv.Atoi(satuanstr)
    if err != nil {
        log.Printf("Error converting satuan: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid satuan format", nil))
    }
    log.Printf("Parsed satuan: %d", satuan)

    minpesananstr := c.FormValue("minpesanan")
    minpesanan, err := strconv.Atoi(minpesananstr)
    if err != nil {
        log.Printf("Error converting minpesanan: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid minpesanan format", nil))
    }
    log.Printf("Parsed minpesanan: %d", minpesanan)

    // Parse multipart form to handle file uploads
    err = c.Request().ParseMultipartForm(32 << 20) // Limit size 32MB
    if err != nil {
        log.Printf("Error parsing multipart form: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to parse form", nil))
    }

    // Handle gambar files
    gambarFiles := make(map[int]*multipart.FileHeader) // Use int for ID
    i := 1
    for {
        key := fmt.Sprintf("gambar[%d]", i)
        fileHeaders, ok := c.Request().MultipartForm.File[key]
        if !ok {
            break
        }
        for _, fileHeader := range fileHeaders {
            gambarIDStr := c.Request().FormValue(fmt.Sprintf("gambarID[%d]", i)) // Use c.Request().FormValue for gambarID
            log.Printf("Form value for gambarID[%d]: %s", i, gambarIDStr) // Debug log for gambarID
            gambarID, err := strconv.Atoi(gambarIDStr)
            if err != nil {
                log.Printf("Error converting gambarID: %v", err)
                return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid gambarID format", nil))
            }
            log.Printf("Received file: %s with size %d bytes and ID: %d", fileHeader.Filename, fileHeader.Size, gambarID)
            gambarFiles[gambarID] = fileHeader // Simpan file berdasarkan ID gambar
        }
        i++
    }
    log.Printf("Number of gambar files received: %d", len(gambarFiles))


    // Buat request untuk update produk
    request := web.UpdatedProduk{
        Name:           name,
        Harga:          harga,
        Satuan:         satuan,
        MinPesanan:     minpesanan,
        Deskripsi:      deskripsistr,
		KategoriProduk: json.RawMessage(kategoriProdukID),
    }

    // Update produk menggunakan service
    updatedProduk, errUpdate := controller.Produk.UpdateProduk(request, id, gambarFiles)
    if errUpdate != nil {
        log.Printf("Error updating product: %v", errUpdate)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUpdate.Error(), nil))
    }

    log.Printf("Product updated successfully with ID: %s", id)
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Data berhasil diupdate", updatedProduk))
}
func (controller *ProdukControllerImpl) GetProdukListWeb(c echo.Context) error {
	// Gunakan helper untuk mengekstrak filters, limit, dan page
	filters, limit, page := helper.ExtractFilter(c.QueryParams())

	// Jika limit atau page tidak valid, berikan nilai default
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	// Ambil kategori dan sort dari query parameters secara langsung
	kategoriproduk := c.QueryParam("kategori")
	sort := c.QueryParam("sort")

	// Panggil service dengan parameter tambahan
	getProduk, totalCount, currentPage, totalPages, nextPage, prevPage, errGetProduk := controller.Produk.GetProduk(limit, page, filters, kategoriproduk, sort)

	if errGetProduk != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", errGetProduk.Error(), model.Pagination{}, nil))
	}

	// Set up pagination
	pagination := model.Pagination{
		CurrentPage:  currentPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalPages:   totalPages,
		TotalRecords: totalCount,
	}

	return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "berhasil", pagination, getProduk))
}


func(controller *ProdukControllerImpl) GetProdukWebId(c echo.Context) error{
	IdProduk := c.Param("id")
	id, _ := uuid.Parse(IdProduk)

	getProduk, errGetProduk := controller.Produk.GetProdukWebId(id)

	if errGetProduk != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id produk", getProduk))
}