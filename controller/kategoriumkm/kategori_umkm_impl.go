package kategoriumkmcontroller

import (
	"net/http"
	"umkm/model"
	"umkm/model/web"
	kategoriumkmservice "umkm/service/kategori_umkm"

	"github.com/labstack/echo/v4"
)

type KategoriUmkmControllerImpl struct {
	kategoriService kategoriumkmservice.KategoriUmkm
}

func NewKategeoriUmkmController(kategoriService kategoriumkmservice.KategoriUmkm) *KategoriUmkmControllerImpl {
	return &KategoriUmkmControllerImpl{
		kategoriService: kategoriService,
	}
}

func (controller *KategoriUmkmControllerImpl) Create(c echo.Context) error {
	kategoriumkm := new(web.CreateCategoriUmkm)

	if err := c.Bind(kategoriumkm); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	result, errSaveKategori := controller.kategoriService.CreateKategori(*kategoriumkm)
	if errSaveKategori != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveKategori.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create kategori Umkm Success", result))
}

//melihat isi kategori
func (controller *KategoriUmkmControllerImpl) GetKategoriList(c echo.Context) error {
	getKategoriUmkm, errGetKategoriUmkm := controller.kategoriService.GetKategoriUmkmList()

	if errGetKategoriUmkm != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetKategoriUmkm.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getKategoriUmkm))
}
