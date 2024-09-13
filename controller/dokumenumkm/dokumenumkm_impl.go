package dokumenumkmcontroller

import (
	"net/http"
	"strconv"
	"umkm/model"
	"umkm/model/web"
	dokumenumkmservice "umkm/service/dokumenumkm"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DokumenUmkmControllerImpl struct {
	dokumenumkmService dokumenumkmservice.DokumenUmkmService
}

func NewDokumenUmkmController(dokumenumkmService dokumenumkmservice.DokumenUmkmService) *DokumenUmkmControllerImpl{
	return &DokumenUmkmControllerImpl{
		dokumenumkmService: dokumenumkmService,
	}
}

func (controller *DokumenUmkmControllerImpl) Create(c echo.Context) error {
	dokumenLegal := new(web.CreateUmkmDokumenLegal)

	// Konversi umkm_id dari string ke uuid.UUID
	umkmIDStr := c.FormValue("umkm_id")
	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UMKM ID format", nil))
	}

	// Konversi dokumen_id dari string ke int
	dokumenIdStr := c.FormValue("dokumen_id")
	dokumenid, err := strconv.Atoi(dokumenIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid Dokumen ID format", nil))
	}

	dokumenLegal.UmkmId = umkmID
	dokumenLegal.DokumenId = dokumenid

	// Ambil file yang diunggah
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to parse multipart form", nil))
	}

	files := form.File["files"]  // Ambil semua file yang diunggah dengan key "files"
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "No files uploaded", nil))
	}

	// Panggil service untuk menyimpan dokumen
	response, err := controller.dokumenumkmService.CreateDokumenUmkm(*dokumenLegal, files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, "Failed to create document", nil))
	}

	// Kembalikan respons yang sukses
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Document created successfully", response))
}
