package kategoriumkmcontroller

import (
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSaveKategori := controller.kategoriService.CreateKategori(*kategoriumkm)
	if errSaveKategori != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategori.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "pembuatan kategori Umkm berhasil", result))
}

//melihat isi kategori
func (controller *KategoriUmkmControllerImpl) GetKategoriList(c echo.Context) error {
	getKategoriUmkm, errGetKategoriUmkm := controller.kategoriService.GetKategoriUmkmList()

	if errGetKategoriUmkm != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetKategoriUmkm.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat seluruh list kategori umkm", getKategoriUmkm))
}

// //melihat isi kategori
func (controller *KategoriUmkmControllerImpl) GetKategoriId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.kategoriService.GetKategoriUmkmId(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat kategori ummkm", getUser))
}

// //update kategori
func (controller *KategoriUmkmControllerImpl) UpdateKategoriId(c echo.Context) error {
    kategori := new(web.UpdateCategoriUmkm)
    id, _ := strconv.Atoi(c.Param("id"))

    if err := c.Bind(kategori); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    userUpdate, errUserUpdate := controller.kategoriService.UpdateKategori(*kategori, id)

    if errUserUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUserUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data kategoriumkm berhasil diperbaharui", userUpdate))
}


// //hapus kategori
func (controller *KategoriUmkmControllerImpl) DeleteKategoriId(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    if errDeleteKategori := controller.kategoriService.DeleteKategoriUmkmId(id); errDeleteKategori != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteKategori.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Penghapusan Kategori umkm berhasil", nil))
}


