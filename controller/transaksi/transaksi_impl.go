package transaksicontroller

import (
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	transaksiservice "umkm/service/transaksi"

	"net/http"

	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransaksiControllerImpl struct {
	transaksiservice transaksiservice.Transaksi
	db                *gorm.DB
}

func NewTransaksiController(transaksi transaksiservice.Transaksi, db    *gorm.DB) *TransaksiControllerImpl {
	return &TransaksiControllerImpl{
		transaksiservice: transaksi,
		db: db,
	}
}

// controller/transaksi/transaksi_controller_impl.go
func (controller *TransaksiControllerImpl) Create(c echo.Context) error {
	transaksi := new(web.CreateTransaksi)

	if err := c.Bind(transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	if err := c.Validate(transaksi); err != nil {
		return err
	}

	result, errSavetransaksi := controller.transaksiservice.CreateTransaksi(*transaksi)

	if errSavetransaksi != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSavetransaksi.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Create Success", result))
}

func (controller *TransaksiControllerImpl) GetKategoriId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getTransaksi, errGetTransaksi := controller.transaksiservice.GetKategoriUmkmId(id)

	if errGetTransaksi != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetTransaksi.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id transaksi", getTransaksi))
}

// //filter
// func (controller *TransaksiControllerImpl) GetTransaksiFilterList(c echo.Context) error {
//     umkmIDStr := c.Param("umkm_id")
//     fmt.Println("Received UMKM ID:", umkmIDStr) // Debug log

//     if umkmIDStr == "" {
//         return echo.NewHTTPError(http.StatusBadRequest, "UMKM ID cannot be empty")
//     }

//     umkmID, err := uuid.Parse(umkmIDStr)
//     if err != nil {
//         fmt.Println("Error parsing UMKM ID:", err) // Debug log
//         return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
//     }

//     kategoriProduk, err := controller.transaksiservice.GetTransaksiFilter(umkmID)
//     if err != nil {
//         return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get kategori produk")
//     }

//     response := map[string]interface{}{
//         "code":   http.StatusOK,
//         "status": "success",
//         "data":   kategoriProduk,
//     }

//     return c.JSON(http.StatusOK, response)
// }

func (controller *TransaksiControllerImpl) GetTransaksiFilterList(c echo.Context) error {
    umkmIDStr := c.Param("umkm_id")
    dateStr := c.QueryParam("date") // Mengambil tanggal dari query parameter
    filters, status, limit, page := helper.ExtractFilterSort(c.QueryParams())

    umkmID, err := uuid.Parse(umkmIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
    }

    // Preparing filters for tanggal
    filtersTanggal := map[string]string{"tanggal": dateStr}
    allowedFilters := []string{"tanggal"}

    // Call service to get transaksi list and pagination data
    transaksiResult, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.transaksiservice.GetTransaksiFilter(umkmID, filtersTanggal, allowedFilters, filters, limit, page, status)
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
    response := model.ResponseToClientpagi(http.StatusOK, "true", "Berhasil melihat seluruh transaksi berdasarkan umkm_id", pagination, transaksiResult)

    return c.JSON(http.StatusOK, response)
}






func (controller *TransaksiControllerImpl) GetTransaksiByYear(c echo.Context) error {
    umkmID := c.Param("umkm_id")
    pageParam := c.QueryParam("page")
    limitParam := c.QueryParam("limit")
    filter := c.QueryParam("filter")

    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        page = 1
    }

    var limit int
    if limitParam == "all" {
        limit = -1
    } else {
        limit, err = strconv.Atoi(limitParam)
        if err != nil || limit <= 0 {
            return c.JSON(http.StatusBadRequest, map[string]interface{}{
                "status":  false,
                "message": "Invalid limit parameter",
            })
        }
    }

    transaksiPerTahun, err := controller.transaksiservice.GetTransaksiByYear(umkmID, page, limit, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status":  false,
            "message": "Failed to retrieve transactions",
            "data":    transaksiPerTahun,
        })
    }

    response := map[string]interface{}{
        "code":    http.StatusOK,
        "status":  true,
        "message": "Menampilkan transaksi berdasarkan tahun",
        "data":    transaksiPerTahun,
    }

    return c.JSON(http.StatusOK, response)
}


func (controller *TransaksiControllerImpl) GetTransaksiByMounth(c echo.Context) error {
    umkmID := c.Param("umkm_id")
    yearParam := c.QueryParam("year")
    pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
    filter := c.QueryParam("filter")

    
    year, err := strconv.Atoi(yearParam)
    if err != nil || year <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid year parameter",
        })
    }

	page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        page = 1  // Default to page 1 if invalid page parameter
    }

var limit int
if limitParam == "all" {
    limit = -1
} else {
    limit, err = strconv.Atoi(limitParam)
    if err != nil || limit <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid limit parameter",
        })
    }
}

    // Memanggil service untuk mendapatkan jumlah transaksi per bulan di tahun tertentu
    transaksiPerTahun, err := controller.transaksiservice.GetTransaksiByMonth(umkmID, year, page,limit, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status":  false,
            "message": "Failed to retrieve transactions",
            "data":    transaksiPerTahun,
        })
    }

    response := map[string]interface{}{
        "code":    http.StatusOK,
        "status":  true,
        "message": "Menampilkan transaksi berdasarkan tahun",
        "data":    transaksiPerTahun,
    }
    // Mengembalikan hasil dalam format JSON
    return c.JSON(http.StatusOK, response)
}


func (controller *TransaksiControllerImpl) GetTransaksiByDate(c echo.Context) error {
    umkmID := c.Param("umkm_id")
    yearParam := c.QueryParam("year")
    monthParam := c.QueryParam("month")
    pageParam := c.QueryParam("page")
    limitParam := c.QueryParam("limit")
    filter := c.QueryParam("filter")

    year, err := strconv.Atoi(yearParam)
    if err != nil || year <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid year parameter",
        })
    }

    month, err := strconv.Atoi(monthParam)
    if err != nil || month <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid month parameter",
        })
    }

    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        page = 1  // Default to page 1 if invalid page parameter
    }

    var limit int
    if limitParam == "all" {
        limit = -1
    } else {
        limit, err = strconv.Atoi(limitParam)
        if err != nil || limit <= 0 {
            return c.JSON(http.StatusBadRequest, map[string]interface{}{
                "status":  false,
                "message": "Invalid limit parameter",
            })
        }
    }

    transaksiPerTahun, err := controller.transaksiservice.GetTransaksiByDate(umkmID, year, month, page, limit, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status":  false,
            "message": fmt.Sprintf("Failed to retrieve transactions: %v", err),
        })
    }

    response := map[string]interface{}{
        "code":    http.StatusOK,
        "status":  true,
        "message": "Menampilkan transaksi berdasarkan tahun",
        "data":    transaksiPerTahun,
    }
    return c.JSON(http.StatusOK, response)
}

