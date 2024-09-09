package aboutuscontroller

import (
	"net/http"
	"strconv"
	"umkm/model"
	web "umkm/model/web/homepage"
	aboutusservice "umkm/service/homepage/aboutus"

	"github.com/labstack/echo/v4"
)

type AboutUsControllerImpl struct {
	aboutus aboutusservice.AboutUsServiceImpl
}

func NewAboutUsController(aboutus aboutusservice.AboutUsServiceImpl) *AboutUsControllerImpl {
	return &AboutUsControllerImpl{
		aboutus: aboutus,
	}
}

func (controller *AboutUsControllerImpl) Create(c echo.Context) error {
	aboutus := new(web.CreateAboutUs)
	aboutus.Description = c.FormValue("name")
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file", nil))
	}

	// Memanggil service untuk menyimpan brand logo
	response, err := controller.aboutus.CreateAboutUs(*aboutus, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "aboutus berhasil dibuat", response))
}

func (controller *AboutUsControllerImpl) GetAboutUs(c echo.Context) error {
	getAboutUs, errGetAboutUs := controller.aboutus.GetAboutUs()

	if errGetAboutUs != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetAboutUs.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getAboutUs))
}

func (controller *AboutUsControllerImpl) GetAboutId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getAboutUs, errGetAboutUs := controller.aboutus.GetAboutUsid(id)

	if errGetAboutUs != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetAboutUs.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getAboutUs))
}

func (controller *AboutUsControllerImpl) UpdateAboutUs(c echo.Context) error {
	// Parse ID dari parameter URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Ambil nilai dari form-data
	name := c.FormValue("description")

	// Ambil file dari form-data jika ada
	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
	}

	// Buat objek request manual
	request := web.UpdateAboutUs{
		Description:   name,
	}

	// Panggil fungsi UpdateTestimonial dari service
	aboutusUpdate, erraboutusUpdate := controller.aboutus.UpdateAboutUs(request, id, file)
	if erraboutusUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, erraboutusUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", aboutusUpdate))
}
