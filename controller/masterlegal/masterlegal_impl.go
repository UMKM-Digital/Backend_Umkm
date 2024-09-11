package masterlegalcontroller

import (
	"net/http"
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