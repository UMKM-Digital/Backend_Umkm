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

func (controller *SliderControllerImpl) DelSlideId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	if errSlider := controller.slider.DeleteId(id); errSlider != nil{
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSlider.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete Slider Success", nil))
}

func(controller *SliderControllerImpl) Update(c echo.Context) error{
	 // Parse ID dari parameter URL
	 id, _ := strconv.Atoi(c.Param("id"))

	 // Ambil nilai dari form-data
	 slidedsc := c.FormValue("slide_desc")
	 slidetitle := c.FormValue("slide_title")
	 
	 // Ambil file dari form-data jika ada
	 file, err := c.FormFile("gambar")
	 if err != nil && err != http.ErrMissingFile {
		 return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
	 }
 
	 // Buat objek request manual
	 request := web.UpdateSlider{
		SlideDesc: slidedsc,
		SlideTitle: slidetitle,
	 }
 
	 // Panggil fungsi UpdateTestimonial dari service
	 testimonalUpdate, errTestimonalUpdate := controller.slider.UpdateTestimonial(request, id, file)
	 if errTestimonalUpdate != nil {
		 return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errTestimonalUpdate.Error(), nil))
	 }
 
	 return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", testimonalUpdate))
}