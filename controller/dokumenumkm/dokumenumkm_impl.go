package dokumenumkmcontroller

import (
	"log"
	"mime/multipart"
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
	umkmIDStr := c.Param("umkm_id")
	umkmID, err := uuid.Parse(umkmIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UMKM ID format", nil))
	}

	// Konversi dokumen_id dari string ke int
	dokumenIdStr := c.Param("dokumen_id")
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

	files := form.File["file"]  // Ambil semua file yang diunggah dengan key "files"
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

func (controller *DokumenUmkmControllerImpl) GetDokumenId(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    umkmkIDStr := c.Param("umkm_id") // Pastikan ini hanya mengambil UUID

    umkmid, err := uuid.Parse(umkmkIDStr)
    if err != nil {
        log.Println("Error parsing umkm_id:", err)
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid umkm ID")
    }

    log.Printf("ID Param: %d, UMKM_ID Param: %s\n", id, umkmkIDStr)

    dokumenUmkm, errdokumenUmkm := controller.dokumenumkmService.GetDokumenUmkmId(id, umkmid)

    if errdokumenUmkm != nil {
        return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, errdokumenUmkm.Error(), nil))
    }

    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", dokumenUmkm))
}




func (controller *DokumenUmkmControllerImpl) UpdateProduk(c echo.Context) error {
    // Ambil ID produk dan UMKM
    idStr := c.Param("id")
    umkmidStr := c.Param("umkm_id")

    log.Printf("Starting update for document with product ID: %s and UMKM ID: %s", idStr, umkmidStr)

    // Parsing UMKM ID (UUID)
    umkmid, err := uuid.Parse(umkmidStr)
    if err != nil {
        log.Printf("Error parsing UMKM UUID: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid UUID format", nil))
    }

    // Parsing Product ID (Integer)
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Printf("Error converting product ID %s to integer: %v", idStr, err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid product ID format", nil))
    }

    log.Printf("Parsed UMKM ID: %s and Product ID: %d successfully", umkmid, id)

    // Parse multipart form untuk menangani upload file
    log.Printf("Attempting to parse multipart form for file upload")
    err = c.Request().ParseMultipartForm(100 << 20) // Limit size 32MB
    if err != nil {
        log.Printf("Error parsing multipart form: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Failed to parse form", nil))
    }
    
    log.Printf("Multipart form parsed successfully")

    // Handle dokumen files
    dokumenFiles := c.Request().MultipartForm.File["dokumen"] // Pastikan "dokumen" sesuai dengan field yang dikirimkan
    if len(dokumenFiles) == 0 {
        log.Printf("No document files were found in the form data")
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "No document files found", nil))
    }

    var dokumenList []*multipart.FileHeader
    log.Printf("Found %d document file(s)", len(dokumenFiles))

    // Log untuk setiap file yang diterima
    for _, fileHeader := range dokumenFiles {
        log.Printf("Processing document file: %s, Size: %d bytes", fileHeader.Filename, fileHeader.Size)
        dokumenList = append(dokumenList, fileHeader)
    }

    log.Printf("Total %d document files added to processing list", len(dokumenList))

    // Update dokumen menggunakan service
    log.Printf("Attempting to update document with product ID: %d and UMKM ID: %s", id, umkmid)
    updatedDokumen, errUpdate := controller.dokumenumkmService.UpdateDokumenUmkm(id, umkmid, dokumenList)
    if errUpdate != nil {
        log.Printf("Error updating document for product ID: %d, UMKM ID: %s: %v", id, umkmid, errUpdate)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errUpdate.Error(), nil))
    }

    log.Printf("Document update successful for product ID: %d", id)
    return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Dokumen berhasil diupdate", updatedDokumen))
}
