package transaksicontroller

import (
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	transaksiservice "umkm/service/transaksi"

	"net/http"

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


func (controller *TransaksiControllerImpl) GetTransaksiFilterList(c echo.Context) error {
    umkmIDStr := c.Param("umkm_id")
    dateStr := c.Param("date") // Mengambil tanggal dari query parameter
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
    pageParam := c.QueryParam("page")
    limitParam := c.QueryParam("limit")
    filter := c.QueryParam("filter")

    userId, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

    // Parse page dan limit
    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        page = 1 // Default ke halaman 1 jika parameter halaman tidak valid
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

    // Panggil service untuk mendapatkan transaksi dengan pagination
    transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := controller.transaksiservice.GetTransaksiByYear(userId, page, limit, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Siapkan data pagination
    pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalRecords,
    }

    // Buat response dengan data pagination dan transaksi
    response := model.ResponseToClientpagi(http.StatusOK, "true", "Berhasil menampilkan transaksi berdasarkan tahun", pagination, transactions)

    return c.JSON(http.StatusOK, response)
}

    func (controller *TransaksiControllerImpl) GetTransaksiByMonth(c echo.Context) error {
        userId, err := helper.GetAuthId(c)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "message": "Unauthorized",
            })
        }

        umkmIDStr := c.QueryParam("umkm_id")
        yearParam := c.QueryParam("year")
        pageParam := c.QueryParam("page")
        limitParam := c.QueryParam("limit")
        filter := c.QueryParam("filter")

        var umkmID uuid.UUID
        // Jika umkm_id tidak disediakan, ambil semua UMKM yang dimiliki oleh user
        if umkmIDStr != "" {
            // Parse UUID untuk UMKM ID
            var err error
            umkmID, err = uuid.Parse(umkmIDStr)
            if err != nil {
                return c.JSON(http.StatusBadRequest, map[string]interface{}{
                    "status":  false,
                    "message": "Invalid UMKM ID",
                })
            }
        }

        // Parse tahun
        year, err := strconv.Atoi(yearParam)
        if err != nil && yearParam != "" {
            return c.JSON(http.StatusBadRequest, map[string]interface{}{
                "status":  false,
                "message": "Invalid year parameter",
            })
        }

        // Parse page dan limit
        page, err := strconv.Atoi(pageParam)
        if err != nil || page <= 0 {
            page = 1  // Default ke halaman 1 jika parameter halaman tidak valid
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

        // Panggil service untuk mendapatkan transaksi dan data pagination
        transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := controller.transaksiservice.GetTransaksiByMonth(umkmID, userId, year, page, limit, filter)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
        }

        // Siapkan data pagination
        pagination := model.Pagination{
            CurrentPage:  currentPage,
            NextPage:     nextPage,
            PrevPage:     prevPage,
            TotalPages:   totalPages,
            TotalRecords: totalRecords,
        }

        // Buat respons dengan pagination dan data transaksi
        response := model.ResponseToClientpagi(http.StatusOK, "true", "Berhasil menampilkan transaksi berdasarkan bulan", pagination, transactions)

        return c.JSON(http.StatusOK, response)
    }


func (controller *TransaksiControllerImpl) GetTransaksiByDate(c echo.Context) error {
    userId, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

    umkmIDStr := c.QueryParam("umkm_id")
    yearParam := c.QueryParam("year")
    monthParam := c.QueryParam("month")
    pageParam := c.QueryParam("page")
    limitParam := c.QueryParam("limit")
    filter := c.QueryParam("filter")

    var umkmID uuid.UUID
    // Jika umkm_id tidak disediakan, ambil semua UMKM yang dimiliki oleh user
    if umkmIDStr != "" {
        var err error
        umkmID, err = uuid.Parse(umkmIDStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]interface{}{
                "status":  false,
                "message": "Invalid UMKM ID",
            })
        }
    }

    // Parse year dan month
    year, err := strconv.Atoi(yearParam)
    if err != nil && yearParam != "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid year parameter",
        })
    }

    month, err := strconv.Atoi(monthParam)
    if err != nil && monthParam != "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status":  false,
            "message": "Invalid month parameter",
        })
    }

    // Parse page dan limit
    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        page = 1 // Default ke halaman 1 jika parameter halaman tidak valid
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

    // Panggil service untuk mendapatkan transaksi dan data pagination
    transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := controller.transaksiservice.GetTransaksiByDate( umkmID, userId, year, month, page, limit, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Siapkan data pagination
    pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalRecords,
    }

    // Buat respons dengan pagination dan data transaksi
    response := model.ResponseToClientpagi(http.StatusOK, "true", "Berhasil menampilkan transaksi berdasarkan tanggal", pagination, transactions)

    return c.JSON(http.StatusOK, response)
}
