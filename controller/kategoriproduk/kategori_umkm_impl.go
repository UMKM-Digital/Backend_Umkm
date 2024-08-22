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
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSaveKategoripruduk := controller.kategoriprodukService.CreateKategori(*kategoriproduk)
	if errSaveKategoripruduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategoripruduk.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Create kategori produk Success", result))
}

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
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "gagal melihat kategori list",
			"code":    http.StatusBadRequest,
		})
    }

    response := map[string]interface{}{
        "code":   http.StatusOK,
        "status": true,
		"message": "list kategori produk",
        "data":   kategoriProduk, 
    }

    return c.JSON(http.StatusOK, response)
}

