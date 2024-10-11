package daerahcontoller

import (
	"net/http"
	"umkm/model"
	daerahservice "umkm/service/daerah"

	"github.com/labstack/echo/v4"
)

type DaerahControllerImpl struct {
	daerahservice daerahservice.Daerah
}

func NewDaerahController(daerahservice daerahservice.Daerah) *DaerahControllerImpl {
	return &DaerahControllerImpl{
		daerahservice: daerahservice,
	}
}


func (controller *DaerahControllerImpl) GetDaerah(c echo.Context) error{
	getSlider, errGetSlider := controller.daerahservice.GetDaerah()

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}

func (controller *DaerahControllerImpl) GetKabupaten(c echo.Context) error{
	id := c.Param("id")
	getSlider, errGetSlider := controller.daerahservice.GetKabupaten(id)

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}

func (controller *DaerahControllerImpl) GetKecamatan(c echo.Context) error{
	id := c.Param("id") 
	getSlider, errGetSlider := controller.daerahservice.GetKecamatan(id)

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}

func (controller *DaerahControllerImpl) GetKelurahan(c echo.Context) error{
	id := c.Param("id")

	getSlider, errGetSlider := controller.daerahservice.GetKelurahan(id)

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}