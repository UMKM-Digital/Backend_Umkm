package sektorusahacontroller

import (
	"net/http"
	"umkm/model"
	web "umkm/model/web/master"
	sektorusahaservice "umkm/service/sektorusaha"

	"github.com/labstack/echo/v4"
)

type SektorUsahaControllerImpl struct {
	sektorusahaService sektorusahaservice.SektorUsaha
}

func NewSektorUsahaController(sektorusahaService sektorusahaservice.SektorUsaha) *SektorUsahaControllerImpl {
	return &SektorUsahaControllerImpl{
		sektorusahaService: sektorusahaService,
	}
}

func (controller *SektorUsahaControllerImpl) Create(c echo.Context) error {
	kategoriumkm := new(web.CreateSektorUsaha)

	if err := c.Bind(kategoriumkm); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSaveKategori := controller.sektorusahaService.CreateSektorUsaha(*kategoriumkm)
	if errSaveKategori != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategori.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "pembuatan kategori sektor usaha berhasil", result))
}

// melihat isi kategori
func (controller *SektorUsahaControllerImpl) GetSektorUsaha(c echo.Context) error {
	getKategoriUmkm, errGetKategoriUmkm := controller.sektorusahaService.GetSektorUsaha()

	if errGetKategoriUmkm != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetKategoriUmkm.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat jenis sektorusaha", getKategoriUmkm))
}