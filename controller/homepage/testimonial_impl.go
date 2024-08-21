package homepagecontroller

import (
	"log"
	"net/http"
	"strconv"
	"umkm/helper"
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


func (controller *TestimonalControllerImpl) GetTestimonial(c echo.Context) error {
	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonial()

	if errGetTestimoni != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getTestimoni))
}

func (controller *TestimonalControllerImpl) DeleteTestimonial(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    if errTestimonial := controller.testimonal.DeleteTestimonial(id); errTestimonial != nil{
        return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errTestimonial.Error(), nil))
    }

    return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Delete Kategori Success", nil))
}

func (controller *TestimonalControllerImpl) GetTestimonialId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonialid(id)

	if errGetTestimoni != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getTestimoni))
}
func (conntroller *TestimonalControllerImpl) UpdateTestimonial(c echo.Context) error{
	testimonal := new(web.UpdateTestimonial)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(testimonal); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	testimonalUpdate, errTestimonalUpdate := conntroller.testimonal.UpdateTestimonial(*testimonal, id)

	if errTestimonalUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errTestimonalUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "data berhasil diupdate", testimonalUpdate))
}

func (controller *TestimonalControllerImpl) GetTestimonialActive(c echo.Context) error {
	getTestimoni, errGetTestimoni := controller.testimonal.GetTestimonialActive()

	if errGetTestimoni != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetTestimoni.Error(), nil))
	}

	log.Println("Controller received testimonials:", getTestimoni) 

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getTestimoni))
}
