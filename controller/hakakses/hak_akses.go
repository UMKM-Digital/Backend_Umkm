package hakaksescontroller

import "github.com/labstack/echo/v4"

type MasterLelgalkController interface {
	UpdateMasterLegalId(c echo.Context) error
}
