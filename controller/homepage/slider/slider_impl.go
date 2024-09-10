package slidercontroller

import (
	"net/http"
	"strconv"
	"umkm/model"
	web "umkm/model/web/homepage"
	sliderservice "umkm/service/homepage/slider"

	"github.com/labstack/echo/v4"
)

type SliderControllerImpl struct {
	slider sliderservice.SliderServiceImpl
}

func NewTestimonialController(slider sliderservice.SliderServiceImpl) *SliderControllerImpl {
	return &SliderControllerImpl{
		slider: slider,
	}
}

func (controller *SliderControllerImpl) Create(c echo.Context) error {
	slider := new(web.CreatedSlider)
	slider.SlideTitle = c.FormValue("slide_title")
	slider.SlideDesc = c.FormValue("slide_desc")
	// Mendapatkan file gambar yang diunggah
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file", nil))
	}

	// Memanggil service untuk menyimpan brand logo
	response, err := controller.slider.CreateSlider(*slider, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "testimonial berhasil dibuat", response))
}

func (controller *SliderControllerImpl) List(c echo.Context) error{
	getSlider, errGetSlider := controller.slider.GetSlider()

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}

func (controller *SliderControllerImpl) GetSlideId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getSlider, errGetSlider := controller.slider.GetSliderid(id)

	if errGetSlider != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}