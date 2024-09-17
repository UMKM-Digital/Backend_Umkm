package masterlegalcontroller

import (
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	masterdokumenlegalservice "umkm/service/masterdokumenlegal"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MasterLegalControllerImpl struct {
	masterLegalService masterdokumenlegalservice.MasterDokumenLegal
}

func NewKategeoriProdukController(masterLegalService masterdokumenlegalservice.MasterDokumenLegal) *MasterLegalControllerImpl {
	return &MasterLegalControllerImpl{
		masterLegalService: masterLegalService,
	}
}

func (controller *MasterLegalControllerImpl) Create(c echo.Context) error {
	masterlegal := new(web.CreateMasterDokumenLegal)

	if err := c.Bind(masterlegal); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	result, errSavemasterlegalpruduk := controller.masterLegalService.CreatedRequest(*masterlegal)
	if errSavemasterlegalpruduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSavemasterlegalpruduk.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Create masterlegal produk Success", result))
}

func (controller *MasterLegalControllerImpl) GetMasterLegalList(c echo.Context) error {
	filters, limit, page := helper.ExtractFilter(c.QueryParams())

    masterlegal, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.masterLegalService.GetMasterLegalList(filters, limit, page)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

    pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalCount,
    }

    return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK,  "true", "berhasil melihat seluruh list kategori umkm",pagination,masterlegal))
}

func (controller *MasterLegalControllerImpl) Delete( c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	if errDeletemasterlegalProduk := controller.masterLegalService.DeleteMasterLegalId(id); errDeletemasterlegalProduk != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeletemasterlegalProduk.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Penghapusan master legal berhasil", nil))
}

func(controller *MasterLegalControllerImpl) GetIdMasterLegalId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getMasterLegal, errGetMasterLegal := controller.masterLegalService.GetMasterLegalid(id)

	if errGetMasterLegal != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetMasterLegal.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil melihat masterlegal produk", getMasterLegal))
}

func(controller *MasterLegalControllerImpl) UpdateMasterLegalId(c echo.Context) error{
	masterlegal := new(web.UpdateMasterDokumenLegal)
    id, _ := strconv.Atoi(c.Param("id"))

    if err := c.Bind(masterlegal); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
    }

    userMasterLegal, errUserMasterLegal := controller.masterLegalService.UpdateMasterLegal(*masterlegal, id)

    if errUserMasterLegal != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUserMasterLegal.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data masterlegal berhasil diperbaharui", userMasterLegal))
}

func (controller *MasterLegalControllerImpl) List(c echo.Context) error {
    umkmIDStr := c.Param("umkm_id")
    umkmID, err := uuid.Parse(umkmIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid UMKM ID")
    }

    filters, limit, page := helper.ExtractFilter(c.QueryParams())

    dokumenStatusList, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.masterLegalService.GetDokumenUmkmStatus(umkmID, filters, limit, page)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve documents")
    }

    pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalCount,
    }

    return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "Successfully retrieved documents", pagination, dokumenStatusList))
}
