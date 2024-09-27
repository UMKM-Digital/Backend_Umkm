package kategoriprodukcontroller

import (
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	kategoriprodukservice "umkm/service/kategori_produk"
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
	filters, limit, page := helper.ExtractFilter(c.QueryParams())


    kategoriProduk, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.kategoriprodukService.GetKategoriProdukList( filters, limit, page)
   
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalCount,
    }

	return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "berhasil melihat seluruh list kategori umkm", pagination, kategoriProduk))
}

func (controller *KategoriProdukControllerImpl) GetKategoriId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.kategoriprodukService.GetKategoriProdukId(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat kategori produk", getUser))
}

func (controller *KategoriProdukControllerImpl) UpdateKategoriProduk(c echo.Context) error{
	kategori := new(web.UpdateCategoriProduk)
    id, _ := strconv.Atoi(c.Param("id"))

    if err := c.Bind(kategori); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    userUpdate, errUserUpdate := controller.kategoriprodukService.UpdateKategoriProduk(*kategori, id)

    if errUserUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUserUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data kategoriproduk berhasil diperbaharui", userUpdate))
}

func (controller *KategoriProdukControllerImpl)Delete( c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	if errDeleteKategoriProduk := controller.kategoriprodukService.DeleteKategoriProdukId(id); errDeleteKategoriProduk != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteKategoriProduk.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Penghapusan Kategori produk berhasil", nil))
}