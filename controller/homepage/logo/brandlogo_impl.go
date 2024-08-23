package brandlogo

import (
	"net/http"
	"strconv"
	"umkm/model"

	// "umkm/model"
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
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file", nil))
	}

	// Memanggil service untuk menyimpan brand logo
	response, err := controller.brandlogoService.CreateBrandlogo(*brandlogo, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "brand logo berhasil dibuat", response))
}

func (controller *BrandLogoControllerImpl) GetBrandLogoList(c echo.Context) error {
	GetBrandLogo, errGetBrandLogo := controller.brandlogoService.GetBrandlogoList()

	if errGetBrandLogo != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetBrandLogo.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", GetBrandLogo))
}

func (controller *BrandLogoControllerImpl) DeleteProdukId(c echo.Context) error {
	// Ambil ID dari URL dan konversi ke integer
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid ID format", nil))
	}

	// Hapus produk berdasarkan ID
	if errDeleteBrandLogo := controller.brandlogoService.DeleteBrandLogo(id); errDeleteBrandLogo != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteBrandLogo.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete logo Success", nil))
}

