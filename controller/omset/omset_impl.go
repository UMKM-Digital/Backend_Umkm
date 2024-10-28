package omsetcontroller

import (
	"net/http"
	"strconv"
	"umkm/model"
	"umkm/model/web"
	omsetservice "umkm/service/omset"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OmsetControllerImpl struct {
	omsetservice omsetservice.OmsetService
}

func NewOmsetController(omsetservice omsetservice.OmsetService) *OmsetControllerImpl{
	return &OmsetControllerImpl{
		omsetservice: omsetservice,
	}
}

func (controller *OmsetControllerImpl) CreateOmsetcontroller(c echo.Context) error {
	omset := new(web.Omset)
	umkmIdParam := c.Param("umkm_id")
	umkmId, err := uuid.Parse(umkmIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid umkm_id", nil))
	}

	// Set umkm_id to omset struct
	omset.UmkmId = umkmId

	if err := c.Bind(omset); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSaveKategoripruduk := controller.omsetservice.CreateOmsetService(*omset)
	if errSaveKategoripruduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategoripruduk.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Create omset Success", result))
}

func(controller *OmsetControllerImpl) LisOmsetController(c echo.Context) error{
	umkmIdParam := c.Param("umkm_id")
	umkmId, err := uuid.Parse(umkmIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid umkm_id", nil))
	}

	tahunParam := c.QueryParam("tahun") // Membaca tahun dari query parameter
	if tahunParam == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Tahun is required", nil))
	}

	result, err := controller.omsetservice.ListOmsetService(umkmId, tahunParam)
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}



func(controller *OmsetControllerImpl) GetOmsetController(c echo.Context) error{
	idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid id parameter", nil))
    }

	getOmset, err := controller.omsetservice.GetOmsetServiceId(id)

	if err != nil{
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat omset", getOmset))
}

func (controller *OmsetControllerImpl) UpdateOmset(c echo.Context) error {
	Omset := new(web.UpdateOmset)
	if err := c.Bind(Omset); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid request data", nil))
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid id parameter", nil))
	}

	getOmset, err := controller.omsetservice.UpdateOmset(*Omset, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil memperbarui omset", getOmset))
}
