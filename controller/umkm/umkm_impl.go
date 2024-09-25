package umkmcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	// "umkm/helper"
	"umkm/helper"
	"umkm/model"
	"umkm/model/entity"
	"umkm/model/web"
	umkmservice "umkm/service/umkm"

	"github.com/google/uuid"
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

	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "Failed to parse form", nil))
	}

	files := c.Request().MultipartForm.File["images"]
	fileHeaders := make(map[string]*multipart.FileHeader)
	for _, file := range files {
		fileHeaders[file.Filename] = file
	}

	umkm.Gambar = json.RawMessage([]byte("[]")) 

    
    fmt.Printf("Form Data: %+v\n", umkm)

    // Get authenticated user ID
    userID, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "Failed to get user ID", nil))
    }

    // Call service to create UMKM
    result, errSaveKategori := controller.umkmservice.CreateUmkm(*umkm, userID, fileHeaders)
    if errSaveKategori != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errSaveKategori.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Pembuatan Umkm Berhasil", result))
}

// //umkm list
func (controller *UmkmControllerImpl) GetUmkmList(c echo.Context) error {
	userId, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", err.Error(), model.Pagination{}, nil))
	}

	filters, limit, page := helper.ExtractFilter(c.QueryParams())
	umkmList, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.umkmservice.GetUmkmListByUserId(c.Request().Context(), userId, filters, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", err.Error(), model.Pagination{}, nil))
	}

	// Jika umkmList kosong, set sebagai array kosong agar tidak null
	if umkmList == nil {
		umkmList = []entity.UmkmFilterEntity{}
	}

	pagination := model.Pagination{
		CurrentPage:  currentPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalPages:   totalPages,
		TotalRecords: totalCount,
	}
	
	return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "berhasil", pagination, umkmList))
}




//filter umkm name
func (controller *UmkmControllerImpl) GetUmkmFilter(c echo.Context) error {
        userId, err := helper.GetAuthId(c)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
        }
    
        filters := map[string]string{"name": c.QueryParam("name")}
        allowedFilters := []string{"name"}
    
        umkmList, err := controller.umkmservice.GetUmkmFilter(c.Request().Context(), userId, filters, allowedFilters)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
        }
    
        return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success mendapatkan umkm", umkmList))
    }
    
func (controller *UmkmControllerImpl) GetUmkmListWeb(c echo.Context) error{
    userId, err := helper.GetAuthId(c)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    umkmList, err := controller.umkmservice.GetUmkmListWeb(c.Request().Context(), userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success melihat list Umkm", umkmList))
}


func (controller *UmkmControllerImpl) GetUmkmId(c echo.Context) error{
	IdUmkm := c.Param("id")
	id, _ := uuid.Parse(IdUmkm)

	getProduk, errGetProduk := controller.umkmservice.GetUmkmId(id)

	if errGetProduk != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil mengambil id transaksi", getProduk))
}

func (controller *UmkmControllerImpl) UpdateUmkm(c echo.Context) error {
    // Parsing UMKM ID (UUID) dari URL parameter
    umkmidStr := c.Param("umkm_id")
    umkmid, err := uuid.Parse(umkmidStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UUID format", nil))
    }

    log.Printf("Parsed UMKM ID: %s successfully", umkmid)

    // Ambil nilai dari form-data
    name := c.FormValue("name")
    noNpwp := c.FormValue("no_npwp")
    namaPenanggungJawab := c.FormValue("nama_penanggung_jawab")
    noKontak := c.FormValue("no_kontak")
    lokasi := c.FormValue("lokasi")
    kategoriUmkmID := c.FormValue("kategori_umkm_id")
    informasiJamBuka := c.FormValue("informasi_jam_buka")
    maps := c.FormValue("maps")

    log.Printf("Form values - Name: %s, NoNpwp: %s, KategoriUmkmId: %s, informasijambuka: %s, ", name, noNpwp, kategoriUmkmID, informasiJamBuka)

    // Ambil file dari form-data jika ada
    files := []*multipart.FileHeader{}
    if file, err := c.FormFile("gambar"); err == nil {
        files = append(files, file)
    } else if err != http.ErrMissingFile {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
    }

    // Buat objek request manual
    request := web.Updateumkm{
        Name:                name,
        NoNpwp:              noNpwp,
        Nama_Penanggung_Jawab: namaPenanggungJawab,
        No_Kontak:          noKontak,
        Lokasi:             lokasi,
        Kategori_Umkm_Id:   json.RawMessage(kategoriUmkmID),
        Informasi_JamBuka: json.RawMessage(informasiJamBuka),
        Maps:               json.RawMessage(maps),
    }

    // Memanggil service untuk update UMKM
    result, err := controller.umkmservice.UpdateUmkmId(request, umkmid, files)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Response sukses
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil di update", result))
}
