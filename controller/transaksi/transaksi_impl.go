package transaksicontroller

import (
	"umkm/model"
	"umkm/model/web"
	transaksiservice "umkm/service/transaksi"

	"net/http"

	"github.com/labstack/echo/v4"
)

type TransaksiControllerImpl struct {
	transaksiservice transaksiservice.Transaksi
}

func NewUmkmController(transaksi transaksiservice.Transaksi) *TransaksiControllerImpl {
	return &TransaksiControllerImpl{
		transaksiservice: transaksi,
	}
}

// controller/transaksi/transaksi_controller_impl.go
func (controller *TransaksiControllerImpl) Create(c echo.Context) error {
	transaksi := new(web.CreateTransaksi)

	if err := c.Bind(transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(transaksi); err != nil {
		return err
	}

	result, errSavetransaksi := controller.transaksiservice.CreateTransaksi(*transaksi)

	if errSavetransaksi != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSavetransaksi.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Register Success", result))
}
