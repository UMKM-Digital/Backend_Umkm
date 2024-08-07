package kategoriprodukcontroller

import (
	"net/http"
	"umkm/model"
	"umkm/model/web"
	kategoriprodukservice "umkm/service/kategori_produk"

	"github.com/google/uuid"
	"fmt"
	"github.com/labstack/echo/v4"
)

type KategoriProdukControllerImpl struct {
	kategoriprodukService kategoriprodukservice.KategoriProdukServiceImpl
}

func NewKategeoriProdukController(kategoriprodukService kategoriprodukservice.KategoriProdukServiceImpl) *KategoriProdukControllerImpl {
	return &KategoriProdukControllerImpl{
		kategoriprodukService: kategoriprodukService,
	}
}

func (controller *KategoriProdukControllerImpl) Create(c echo.Context) error {
	kategoriproduk := new(web.CreateCategoriProduk)

	if err := c.Bind(kategoriproduk); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	result, errSaveKategoripruduk := controller.kategoriprodukService.CreateKategori(*kategoriproduk)
	if errSaveKategoripruduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveKategoripruduk.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create kategori Umkm Success", result))
}

// func (controller *KategoriProdukControllerImpl) GetKategoriList(c echo.Context) error {
// 	// getKategoriProduk, errGetKategoriUmkm := controller.kategoriprodukService.GetKategoriProdukList()

// 	// if errGetKategoriUmkm != nil {
// 	// 	return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetKategoriUmkm.Error(), nil))
// 	// }

// 	// return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getKategoriProduk))
// 	umkmIDStr := c.QueryParam("umkm_id")
//     umkmID, err := uuid.Parse(umkmIDStr)
//     if err != nil {
//         return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
//     }

//     kategoriProduk, err := controller.kategoriprodukService.GetKategoriProdukList(umkmID)
//     if err != nil {
//         return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get kategori produk")
//     }

//     response := map[string]interface{}{
//         "code":   200,
//         "status": "success",
//         "data":   kategoriProduk,
//     }

//     return c.JSON(http.StatusOK, response)
// }
func (controller *KategoriProdukControllerImpl) GetKategoriList(c echo.Context) error {
    umkmIDStr := c.Param("umkm_id")
    fmt.Println("Received UMKM ID:", umkmIDStr) // Debug log

    if umkmIDStr == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "UMKM ID cannot be empty")
    }

    umkmID, err := uuid.Parse(umkmIDStr)
    if err != nil {
        fmt.Println("Error parsing UMKM ID:", err) // Debug log
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
    }

    kategoriProduk, err := controller.kategoriprodukService.GetKategoriProdukList(umkmID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get kategori produk")
    }

    response := map[string]interface{}{
        "code":   http.StatusOK,
        "status": "success",
        "data":   kategoriProduk,
    }

    return c.JSON(http.StatusOK, response)
}
