package beritacontroller

import (
	"net/http"
	"strconv"
	"umkm/helper"
	"umkm/model"
	web "umkm/model/web/homepage"
	beritaservice "umkm/service/homepage/berita"

	"github.com/labstack/echo/v4"
)

type BeritaControllerImpl struct {
	berita beritaservice.BeritaServiceImpl
}

func NewBeritaController(berita beritaservice.BeritaServiceImpl) *BeritaControllerImpl {
	return &BeritaControllerImpl{
		berita: berita,
	}
}

func (controller *BeritaControllerImpl) Create(c echo.Context) error {
    // Membuat struct untuk input berita
	berita := new(web.CreatedBerita)
	berita.Title = c.FormValue("title")
	berita.Content = c.FormValue("content")

    // Mendapatkan file gambar yang diunggah
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to get the uploaded file", nil))
	}

    // Mendapatkan userID dari token JWT menggunakan helper
	userID, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "Failed to get user ID", nil))
	}

    // Memanggil service untuk menyimpan berita beserta userID (AuthorId)
	response, err := controller.berita.CreatedBerita(*berita, file, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Berita berhasil dibuat", response))
}

func(controller *BeritaControllerImpl) LIst(c echo.Context) error{
	limitStr := c.QueryParam("limit")
    pageStr := c.QueryParam("page")

   
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 {
        limit = 10 
    }

    page, err := strconv.Atoi(pageStr)
    if err != nil || page <= 0 {
        page = 1
    }

    // Panggil method GetBeritaList dengan limit dan page
    beritaList, totalCount, currentPage, totalPages, nextPage, prevPage, err := controller.berita.GetBeritaList(c.Request().Context(), limit, page)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", err.Error(), model.Pagination{}, nil))
    }

    // Buat pagination response
    pagination := model.Pagination{
        CurrentPage:  currentPage,
        NextPage:     nextPage,
        PrevPage:     prevPage,
        TotalPages:   totalPages,
        TotalRecords: totalCount,
    }

    // Kembalikan response dengan data UMKM list dan pagination
    return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "berhasil", pagination, beritaList))
}

func(controller *BeritaControllerImpl) Delete( c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

    if errTestimonial := controller.berita.DeleteBerita(id); errTestimonial != nil{
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errTestimonial.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete berita Success", nil))
}

func(controller *BeritaControllerImpl) GetId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getBerita, errGetBerita := controller.berita.GetBeritaByid(id)

	if errGetBerita != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetBerita.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getBerita))
}

func (controller *BeritaControllerImpl) DeleteTestimonial(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    if errTestimonial := controller.berita.DeleteBerita(id); errTestimonial != nil{
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errTestimonial.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete Kategori Success", nil))
}

func (controller *BeritaControllerImpl) GetTestimonialId(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	getTestimoni, errGetTestimoni := controller.berita.GetBeritaByid(id)

	if errGetTestimoni != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errGetTestimoni.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getTestimoni))
}

func (controller *BeritaControllerImpl) Update(c echo.Context) error {
    // Parse ID dari parameter URL
    id, _ := strconv.Atoi(c.Param("id"))

    // Ambil nilai dari form-data
    title := c.FormValue("title")
    content := c.FormValue("content")
    
    // Ambil file dari form-data jika ada
    file, err := c.FormFile("image")
    if err != nil && err != http.ErrMissingFile {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
    }

    // Buat objek request manual
    request := web.UpdtaedBerita{
        Title:   title,
        Content: content,
    }

    // Panggil fungsi UpdateTestimonial dari service
    beritaUpdate, errberitaUpdate := controller.berita.UpdateBerita(request, id, file)
    if errberitaUpdate != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errberitaUpdate.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", beritaUpdate))
}