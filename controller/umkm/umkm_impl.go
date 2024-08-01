package umkmcontroller

import (
	"net/http"
	"umkm/model"
	"umkm/model/web"
	umkmservice "umkm/service/umkm"

	"github.com/labstack/echo/v4"
)

type UmkmControllerImpl struct {
	umkmservice umkmservice.Umkm
}

func NewUmkmController(umkm umkmservice.Umkm) *UmkmControllerImpl {
	return &UmkmControllerImpl{
		umkmservice: umkm,
	}
}

func (controller *UmkmControllerImpl) Create(c echo.Context) error {
	umkm := new(web.UmkmRequest)

	if err := c.Bind(umkm); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	result, errSaveKategori := controller.umkmservice.CreateUmkm(*umkm)
	if errSaveKategori != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveKategori.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create  Umkm Success", result))
}