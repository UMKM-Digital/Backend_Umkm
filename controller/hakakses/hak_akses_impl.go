package hakaksescontroller


import (
	"net/http"
	"umkm/model"
	"umkm/model/web"
	hakaksesservice "umkm/service/hak_akses"
	"github.com/labstack/echo/v4"
)

type HakAksesControllerImpl struct {
	hakaksesService hakaksesservice.HakAkses
}

func NewHakAksesController(hakaksesService hakaksesservice.HakAkses) *HakAksesControllerImpl {
	return &HakAksesControllerImpl{
		hakaksesService: hakaksesService,
	}
}


func(controller *HakAksesControllerImpl) UpdateHakAksesIds(c echo.Context) error {
    var request web.HakAksesUpdate
    if err := c.Bind(&request); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid request format", nil))
    }

    updatedHakAkses, err := controller.hakaksesService.UpdateBulkHakAkses(request)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Hak akses berhasil diperbaharui", updatedHakAkses))
}
