package produkcontroller

import (
	"encoding/json"
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

// func (controller *ProdukControllerImpl) CreateProduk(c echo.Context) error {
// 	produk := new(web.WebProduk)

// 	// Konversi umkm_id dari string ke uuid.UUID
// 	umkmIDStr := c.FormValue("umkm_id")
// 	umkmID, err := uuid.Parse(umkmIDStr)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UMKM ID format", nil))
// 	}

// 	produkHargastr := c.FormValue("harga")
// 	produkHarga, err := strconv.Atoi(produkHargastr)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest,false, "Invalid harga format", nil))
// 	}


// 	produkminpesananstr := c.FormValue("min_pesanan")
// 	produkminpesanan, err := strconv.Atoi(produkminpesananstr)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid min_pesanan format", nil))
// 	}

// 	kategorid := c.FormValue("kategori_produk_id")
//     if kategorid != "" {
//         produk.KategoriProduk = json.RawMessage(kategorid)
//     }


// 	produk.UmkmId = umkmID
// 	produk.Harga = produkHarga
// 	produk.Satuan = c.FormValue("satuan")
// 	produk.MinPesanan = produkminpesanan
// 	produk.Name = c.FormValue("nama")
// 	produk.Deskripsi = c.FormValue("deskripsi")

// 	// Handle image upload
// 	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
// 		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "Failed to parse form", nil))
// 	}

// 	files := c.Request().MultipartForm.File["gambar"]
// 	fileHeaders := make(map[string]*multipart.FileHeader)
// 	for _, file := range files {
// 		fileHeaders[file.Filename] = file
// 	}

// 	produk.GambarId = json.RawMessage([]byte("[]")) // Default empty JSON array
// 	result, errSaveProduk := controller.Produk.CreateProduk(*produk, fileHeaders)
// 	if errSaveProduk != nil {
// 		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveProduk.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "membut produk berhasil", result))
// }


func (controller *ProdukControllerImpl) CreateProduk(c echo.Context) error {
	produk := new(web.WebProduk)

	// Konversi umkm_id dari string ke uuid.UUID
	umkmIDStr := c.FormValue("umkm_id")
	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UMKM ID format", nil))
	}

	// Konversi harga dari string ke integer
	produkHargastr := c.FormValue("harga")
	produkHarga, err := strconv.Atoi(produkHargastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid harga format", nil))
	}

	// Konversi min_pesanan dari string ke integer
	produkMinPesananStr := c.FormValue("min_pesanan")
	produkMinPesanan, err := strconv.Atoi(produkMinPesananStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid min_pesanan format", nil))
	}

	// Konversi kategori_produk_id ke json.RawMessage
	kategoriID := c.FormValue("kategori_produk_id")
	if kategoriID != "" {
		produk.KategoriProduk = json.RawMessage(kategoriID)
	}

	// Set field produk lainnya
	produk.UmkmId = umkmID
	produk.Harga = produkHarga
	produk.Satuan = c.FormValue("satuan")
	produk.MinPesanan = produkMinPesanan
	produk.Name = c.FormValue("nama")
	produk.Deskripsi = c.FormValue("deskripsi")

	// Handle image upload
	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "Failed to parse form", nil))
	}

	// Ambil semua file gambar yang diunggah
	files := c.Request().MultipartForm.File["gambar"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "No images uploaded", nil))
	}

	// Simpan file sesuai urutan yang diupload
	fileHeaders := make([]*multipart.FileHeader, 0) // Menggunakan slice untuk urutan file
	for _, file := range files {
		fileHeaders = append(fileHeaders, file) // Simpan dalam urutan aslinya
	}

	// Call service untuk menyimpan produk dan gambar
	result, errSaveProduk := controller.Produk.CreateProduk(*produk, fileHeaders)
	if errSaveProduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Produk berhasil dibuat", result))
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
	Sort := c.QueryParam("sort")

	filters, limit, page := helper.ExtractFilter(c.QueryParams())

	if produkIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Produk ID cannot be empty")
	}

	produkId, err := uuid.Parse(produkIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Produk ID")
	}

	// Call service to get produk list and pagination data
	result, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.Produk.GetProdukList(produkId, filters, limit, page, kategoriProdukID, Sort)
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
    deskripsi := c.FormValue("deskripsi")
    kategoriJSON := c.FormValue("kategori") // Ambil kategori sebagai JSON
    satuan := c.FormValue("satuan") // Ambil kategori sebagai JSON
    hargaStr := c.FormValue("harga")
    harga, err := strconv.Atoi(hargaStr)
    if err != nil {
        log.Printf("Error converting harga: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid harga format", nil))
    }

    // satuanStr := c.FormValue("satuan")
    // satuan, err := strconv.Atoi(satuanStr)
    // if err != nil {
    //     log.Printf("Error converting satuan: %v", err)
    //     return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid satuan format", nil))
    // }

    minPesananStr := c.FormValue("minpesanan")
    minPesanan, err := strconv.Atoi(minPesananStr)
    if err != nil {
        log.Printf("Error converting minpesanan: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid minpesanan format", nil))
    }

    // Parse multipart form to handle file uploads
    err = c.Request().ParseMultipartForm(32 << 20) // Limit size 32MB
    if err != nil {
        log.Printf("Error parsing multipart form: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to parse form", nil))
    }

    // Handle gambar files
    // Ambil gambar yang baru diunggah
    var gambarFiles []*multipart.FileHeader
    if fileHeaders, ok := c.Request().MultipartForm.File["gambar"]; ok && len(fileHeaders) > 0 {
        gambarFiles = fileHeaders // Ambil semua file gambar
    }

    // Buat request untuk update produk
    request := web.UpdatedProduk{
        Name:           name,
        Harga:          harga,
        Satuan:         satuan,
        MinPesanan:     minPesanan,
        Deskripsi:      deskripsi,
        KategoriProduk: json.RawMessage(kategoriJSON),
    }

    // Update produk menggunakan service
    updatedProduk, errUpdate := controller.Produk.UpdateProduk(request, id, gambarFiles)
    if errUpdate != nil {
        log.Printf("Error updating product: %v", errUpdate)
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errUpdate.Error(), nil))
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


func (controller *ProdukControllerImpl) GetProdukByLogin(c echo.Context) error {
    // Ambil user ID dari token atau session (contoh menggunakan JWT)
    userId, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

    // Ambil semua produk yang terkait dengan UMKM yang dimiliki oleh user
    produkEntities, err := controller.Produk.GetProdukByUser(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "Gagal mengambil produk",
        })
    }

    return c.JSON(http.StatusOK, produkEntities)
}


func(controller *ProdukControllerImpl) GetProdukBaru(c echo.Context) error{
	IdUmkm := c.Param("id")
	id, _ := uuid.Parse(IdUmkm)

	getProduk, errGetProduk := controller.Produk.GetProdukBaru(id)

	if errGetProduk != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id produk", getProduk))
}

func (controller *ProdukControllerImpl) GetTopProuduk(c echo.Context) error {
	IdUmkm := c.QueryParam("id")
	id, _ := uuid.Parse(IdUmkm)

	getTestimoni, errGetTestimoni := controller.Produk.GetTopProduk(id)

	if errGetTestimoni != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getTestimoni))
}

func (controller *ProdukControllerImpl) UpdateTopProduk(c echo.Context) error{
	produk := new(web.UpdatePorudkActive)
    IdUmkm := c.Param("id")
	id, _ := uuid.Parse(IdUmkm)

	if err := c.Bind(produk); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    produkUpdate, errprodukUpdate := controller.Produk.UpdateProdukActive(*produk, id)

    if errprodukUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errprodukUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", produkUpdate))
}

func (controller *ProdukControllerImpl) GetProdukActive(c echo.Context) error {
    getProduk, errGetProduk := controller.Produk.GetTopProdukActive()
    if errGetProduk != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetProduk.Error(), nil))
    }
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getProduk))
}

//login google

