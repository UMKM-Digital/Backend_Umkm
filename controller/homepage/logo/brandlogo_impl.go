package brandlogo

import (
	"net/http"
	"umkm/helper"
	"umkm/model"
	web "umkm/model/web/homepage"
	brandlogoservice "umkm/service/homepage/brandlogo"

	"github.com/labstack/echo/v4"
)

type BrandLogoControllerImpl struct {
	brandlogoService brandlogoservice.Brandlogo
}

func NewBrandLogoController(brandlogoService brandlogoservice.Brandlogo) *BrandLogoControllerImpl {
	return &BrandLogoControllerImpl{
		brandlogoService: brandlogoService,
	}
}

func (controller *BrandLogoControllerImpl) Create(c echo.Context) error {
	// Mendapatkan nilai untuk nama brand dari form
	brandlogo := new(web.CreatedBrandLogo)
	brandlogo.BrandName = c.FormValue("brand_name")

	// Mendapatkan file gambar yang diunggah
	file, err := c.FormFile("brand_logo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToJsonBrandLogo(http.StatusBadRequest, "Failed to get the uploaded file", nil))
	}

	// Memanggil service untuk menyimpan brand logo
	response, err := controller.brandlogoService.CreateBrandlogo(*brandlogo, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToJsonBrandLogo(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToJsonBrandLogo(http.StatusOK, "Brand logo successfully created", response))
}

func (controller *BrandLogoControllerImpl) GetBrandLogoList(c echo.Context) error {
	GetBrandLogo, errGetBrandLogo := controller.brandlogoService.GetBrandlogoList()

	if errGetBrandLogo != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetBrandLogo.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", GetBrandLogo))
}