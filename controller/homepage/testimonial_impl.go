package homepagecontroller

import (
	"log"
	"net/http"
	"strconv"
	"umkm/model"
	web "umkm/model/web/homepage"
	homepageservice "umkm/service/homepage"

	"github.com/labstack/echo/v4"
)

type TestimonalControllerImpl struct {
	testimonal homepageservice.TestimonalServiceImpl
}

func NewTestimonialController(testimonal homepageservice.TestimonalServiceImpl) *TestimonalControllerImpl {
	return &TestimonalControllerImpl{
		testimonal: testimonal,
	}
}

func (controller *TestimonalControllerImpl) Create(c echo.Context) error {
	testimonal := new(web.CreateTestimonial)
	testimonal.Name = c.FormValue("name")
	testimonal.Quotes = c.FormValue("quote")
	// Mendapatkan file gambar yang diunggah
	file, err := c.FormFile("gambar_testi")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file", nil))
	}

	// Memanggil service untuk menyimpan brand logo
	response, err := controller.testimonal.CreateTestimonial(*testimonal, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "testimonial berhasil dibuat", response))
}


func (controller *TestimonalControllerImpl) GetTestimonial(c echo.Context) error {
	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonial()

	if errGetTestimoni != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getTestimoni))
}

func (controller *TestimonalControllerImpl) DeleteTestimonial(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    if errTestimonial := controller.testimonal.DeleteTestimonial(id); errTestimonial != nil{
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errTestimonial.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete Kategori Success", nil))
}

func (controller *TestimonalControllerImpl) GetTestimonialId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonialid(id)

	if errGetTestimoni != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getTestimoni))
}

func (controller *TestimonalControllerImpl) UpdateTestimonial(c echo.Context) error {
    // Parse ID dari parameter URL
    id, _ := strconv.Atoi(c.Param("id"))

    // Ambil nilai dari form-data
    name := c.FormValue("name")
    quotes := c.FormValue("quotes")
    
    // Ambil file dari form-data jika ada
    file, err := c.FormFile("gambar")
    if err != nil && err != http.ErrMissingFile {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
    }

    // Buat objek request manual
    request := web.UpdateTestimonial{
        Name:   name,
        Quotes: quotes,
    }

    // Panggil fungsi UpdateTestimonial dari service
    testimonalUpdate, errTestimonalUpdate := controller.testimonal.UpdateTestimonial(request, id, file)
    if errTestimonalUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errTestimonalUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", testimonalUpdate))
}

func (controller *TestimonalControllerImpl) GetTestimonialActive(c echo.Context) error {
    getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonialActive()
    if errGetTestimoni != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetTestimoni.Error(), nil))
    }

    log.Println("Controller received testimonials:", getTestimoni)

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getTestimoni))
}