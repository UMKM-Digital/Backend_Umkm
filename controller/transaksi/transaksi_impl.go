package transaksicontroller

import (
	"strconv"
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
	dateStr := c.Param("date") // Pastikan parameter tanggal sesuai

	// Debug log untuk memeriksa ID UMKM
	fmt.Println("Received UMKM ID:", umkmIDStr)

	if umkmIDStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "UMKM ID cannot be empty")
	}

	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		fmt.Println("Error parsing UMKM ID:", err) // Debug log
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
	}

	filters := map[string]string{"tanggal": dateStr}
	allowedFilters := []string{"tanggal"} // Sesuaikan filter yang diizinkan

	// Panggil metode GetTransaksiFilter dari service
	transaksiList, err := controller.transaksiservice.GetTransaksiFilter(umkmID, filters, allowedFilters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Invalid request",
			"code":    http.StatusInternalServerError,
		})
	}

	response := map[string]interface{}{
		"code":   http.StatusOK,
		"status": true,
		"message": "menampilkan transaksi berdsakan tanggal",
		"data":   transaksiList,
	}

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
    mountParam := c.QueryParam("month")
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
    mounth, err := strconv.Atoi(mountParam)
    if err != nil || mounth <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid mounth parameter",
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
    transaksiPerTahun, err := controller.transaksiservice.GetTransaksiByDate(umkmID, year, mounth, page, limit, filter)
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

