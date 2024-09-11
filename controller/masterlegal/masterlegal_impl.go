package masterlegalcontroller

import (
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	masterdokumenlegalservice "umkm/service/masterdokumenlegal"

	"github.com/labstack/echo/v4"
)

type MasterLegalControllerImpl struct {
	masterLegalService masterdokumenlegalservice.MasterDokumenLegal
}

func NewKategeoriProdukController(masterLegalService masterdokumenlegalservice.MasterDokumenLegal) *MasterLegalControllerImpl {
	return &MasterLegalControllerImpl{
		masterLegalService: masterLegalService,
	}
}

func (controller *MasterLegalControllerImpl) Create(c echo.Context) error {
	masterlegal := new(web.CreateMasterDokumenLegal)

	if err := c.Bind(masterlegal); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSaveKategoripruduk := controller.masterLegalService.CreatedRequest(*masterlegal)
	if errSaveKategoripruduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategoripruduk.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Create kategori produk Success", result))
}

func (controller *MasterLegalControllerImpl) GetMasterLegalList(c echo.Context) error {
	filters, limit, page := helper.ExtractFilter(c.QueryParams())

    masterlegal, err := controller.masterLegalService.GetMasterLegalList(filters, limit, page)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "gagal melihat dokumen master legal list",
			"code":    http.StatusBadRequest,
		})
    }

    response := map[string]interface{}{
        "code":   http.StatusOK,
        "status": true,
		"message": "list  dokumen master lagal",
        "data":   masterlegal, 
    }

    return c.JSON(http.StatusOK, response)
}

func (controller *MasterLegalControllerImpl) Delete( c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	if errDeleteKategoriProduk := controller.masterLegalService.DeleteMasterLegalId(id); errDeleteKategoriProduk != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteKategoriProduk.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Penghapusan master legal berhasil", nil))
}

func(controller *MasterLegalControllerImpl) GetIdMasterLegalId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getMasterLegal, errGetMasterLegal := controller.masterLegalService.GetMasterLegalid(id)

	if errGetMasterLegal != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetMasterLegal.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat kategori produk", getMasterLegal))
}