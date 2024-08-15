 package homepagecontroller

import (
	"net/http"
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

	if err := c.Bind(testimonal); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	result, errSaveTestimonial := controller.testimonal.CreateTestimonial(*testimonal)
if errSaveTestimonial != nil {
    return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveTestimonial.Error(), nil))
}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create Testimonial Success", result))
}


func (controller *TestimonalControllerImpl) GetKategoriList(c echo.Context) error {
	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonial()

	if errGetTestimoni != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getTestimoni))
}