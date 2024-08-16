package umkmcontroller

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	// "umkm/helper"
	"umkm/helper"
	"umkm/model"
	"umkm/model/web"
	umkmservice "umkm/service/umkm"

	"github.com/labstack/echo/v4"
)

type UmkmControllerImpl struct {
    umkmservice umkmservice.Umkm
}

func NewUmkmController(umkm umkmservice.Umkm) *UmkmControllerImpl {
    return &UmkmControllerImpl{
        umkmservice: umkm,
       
    }
}

func (controller *UmkmControllerImpl) Create(c echo.Context) error {
    // Bind form data
    umkm := new(web.UmkmRequest)
    umkm.Name = c.FormValue("name")
    umkm.NoNpwp = c.FormValue("no_npwp")
    umkm.Nama_Penanggung_Jawab = c.FormValue("nama_penanggung_jawab")
    umkm.No_Kontak = c.FormValue("no_kontak")
    umkm.Lokasi = c.FormValue("lokasi")

    // Handle JSON fields
    kategoriUmkmId := c.FormValue("kategori_umkm_id")
    if kategoriUmkmId != "" {
        umkm.Kategori_Umkm_Id = json.RawMessage(kategoriUmkmId)
    }

    informasiJamBuka := c.FormValue("informasi_jambuka")
    if informasiJamBuka != "" {
        umkm.Informasi_JamBuka = json.RawMessage(informasiJamBuka)
    }

    maps := c.FormValue("maps")
    if maps != "" {
        umkm.Maps = json.RawMessage(maps)
    }

    // Handle image URLs
    // var gambarURLs []string
    // if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
    //     return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to parse form", nil))
    // }

    // files := c.Request().MultipartForm.File["images"]
    // for _, file := range files {
    //     url, err := helper.HandleFileUpload(file, "uploads")
    //     if err != nil {
    //         return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to upload file", nil))
    //     }
    //     gambarURLs = append(gambarURLs, url)
    // }

    // gambarURLsJSON, err := json.Marshal(gambarURLs)
    	// Handle image upload
	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to parse form", nil))
	}

	files := c.Request().MultipartForm.File["images"]
	fileHeaders := make(map[string]*multipart.FileHeader)
	for _, file := range files {
		fileHeaders[file.Filename] = file
	}

	umkm.Gambar = json.RawMessage([]byte("[]")) 

    // if err != nil {
    //     return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "Failed to marshal image URLs", nil))
    // }
    // umkm.Gambar = json.RawMessage(gambarURLsJSON)

    // Log data for debugging
    fmt.Printf("Form Data: %+v\n", umkm)

    // Get authenticated user ID
    userID, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "Failed to get user ID", nil))
    }

    // Call service to create UMKM
    result, errSaveKategori := controller.umkmservice.CreateUmkm(*umkm, userID, fileHeaders)
    if errSaveKategori != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveKategori.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Create UMKM Success", result))
}

    func (controller *UmkmControllerImpl) GetUmkmList(c echo.Context) error {
        userId, err := helper.GetAuthId(c)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
        }
    
        umkmList, err := controller.umkmservice.GetUmkmListByUserId(c.Request().Context(), userId)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
        }
    
        return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", umkmList))
    }

    func (controller *UmkmControllerImpl) GetUmkmFilter(c echo.Context) error {
        userId, err := helper.GetAuthId(c)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
        }
    
        filters := map[string]string{"name": c.QueryParam("name")}
        allowedFilters := []string{"name"}
    
        umkmList, err := controller.umkmservice.GetUmkmFilter(c.Request().Context(), userId, filters, allowedFilters)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
        }
    
        return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", umkmList))
    }
    