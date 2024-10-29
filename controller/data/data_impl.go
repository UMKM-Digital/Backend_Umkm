package datacontroller

import (
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	dataservice "umkm/service/data"

	"github.com/labstack/echo/v4"
)

type DataControllerImpl struct {
	dataservice dataservice.AuthUserService
}

func NewUmkmController(dataservice dataservice.AuthUserService) *DataControllerImpl {
	return &DataControllerImpl{
		dataservice: dataservice,
	}
}


func (controller *DataControllerImpl) CountData(c echo.Context) error {
    // Memanggil service untuk menghitung jumlah pengguna berdasarkan gender
    result, err := controller.dataservice.CountAtas()
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}


func (controller *DataControllerImpl) GrafikKategoriBySektorHandler(c echo.Context) error {
    // Mengambil parameter sektor_usaha_id dari query string
    sektorUsahaIDParam := c.QueryParam("sektor_usaha_id")
    sektorUsahaID, err := strconv.Atoi(sektorUsahaIDParam)
    if err != nil || sektorUsahaID <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "code":    http.StatusBadRequest,
            "message": "Invalid sektor usaha ID",
            "data":    nil,
        })
    }
    tahun := c.QueryParam("tahun")
    tahunint, err := strconv.Atoi(tahun)
    if err != nil || tahunint <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "code":    http.StatusBadRequest,
            "message": "Invalid sektor usaha ID",
            "data":    nil,
        })
    }

     // Mengambil parameter kecamatan dan kelurahan dari query string
     kecamatan := c.QueryParam("kecamatan")
     kelurahan := c.QueryParam("kelurahan")

    // Panggil service untuk mendapatkan data kategori UMKM berdasarkan sektor
    result, err := controller.dataservice.GrafikKategoriBySektor(c.Request().Context(), sektorUsahaID, kecamatan, kelurahan, tahunint)
    if err != nil {
         return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}


func (controller *DataControllerImpl) TotalUmkmKriteriaUsahaPerBulanHandler(c echo.Context) error {
    // Ambil parameter tahun dari query string
    tahunParam := c.QueryParam("tahun")
    if tahunParam == "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Tahun tidak boleh kosong",
        })
    }

    // Konversi parameter tahun menjadi integer
    tahun, err := strconv.Atoi(tahunParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Tahun harus berupa angka",
        })
    }

    // Panggil service untuk mendapatkan data UMKM
    result, err := controller.dataservice.TotalUmkmKriteriaUsahaPerBulan(tahun)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Return hasil perhitungan dalam bentuk JSON
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}

func (controller *DataControllerImpl) CountUmkmBulan(c echo.Context) error {
    // Memanggil service untuk menghitung jumlah pengguna berdasarkan gender
    result, err := controller.dataservice.TotalUmkmBinaan()
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}

//omset
func (controller *DataControllerImpl) CountOmzets(c echo.Context) error {
    // Memanggil service untuk menghitung jumlah pengguna berdasarkan gender
    result, err := controller.dataservice.TotalOmzetBulanIni()
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}


func(controller *DataControllerImpl) CountPengggunaUmkm(c echo.Context) error{
    userId, err := helper.GetAuthId(c)


    result, err := controller.dataservice.DataUmkm(userId)
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    // Mengembalikan respons sukses dengan data yang diperoleh
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}

func (controller *DataControllerImpl) CountPenggunaOmzet(c echo.Context) error {
    userId, err := helper.GetAuthId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "User tidak terautentikasi",
        })
    }

    tahunParam := c.QueryParam("tahun")
    if tahunParam == "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Tahun tidak boleh kosong",
        })
    }

    // Konversi tahunParam dari string ke int
    tahun, err := strconv.Atoi(tahunParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Format tahun tidak valid",
        })
    }

    result, err := controller.dataservice.DataOmzetUmkm(userId, tahun)
    if err != nil {
        // Mengembalikan respons error jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", result))
}
