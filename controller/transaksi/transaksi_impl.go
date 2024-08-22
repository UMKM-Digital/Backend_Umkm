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

func NewUmkmController(transaksi transaksiservice.Transaksi, db    *gorm.DB) *TransaksiControllerImpl {
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

