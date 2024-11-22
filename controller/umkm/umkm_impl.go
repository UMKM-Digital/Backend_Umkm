package umkmcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	// "umkm/helper"
	"umkm/helper"
	"umkm/model"
	"umkm/model/entity"
	"umkm/model/web"
	umkmservice "umkm/service/umkm"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
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
	umkm.Deskripsi = c.FormValue("deskripsi")
	umkm.Lokasi = c.FormValue("lokasi")
	umkm.KodeProv = c.FormValue("kode_prov")
	umkm.KodeKabupaten = c.FormValue("kode_kab")
	umkm.KodeKec = c.FormValue("kode_kec")
	umkm.KodeKel = c.FormValue("kode_kelurahan")
	umkm.Rt = c.FormValue("rt")
	umkm.Rw = c.FormValue("rw")
	umkm.KodePos = c.FormValue("kode_pos")
	umkm.NoNpwd = c.FormValue("no_npwd")
	umkm.BahanBakar = c.FormValue("bahan_bakar")
	umkm.Kapasitas = c.FormValue("bahan_bakar")
	umkm.JenisUsaha = c.FormValue("jenis_usaha")
	umkm.NoNib = c.FormValue("no_nib")
	umkm.KriteriaUsaha = c.FormValue("kriteria_usaha")
	umkm.SektorUsaha = c.FormValue("sektor_usaha")
	umkm.StatusTempatUsaha = c.FormValue("status_tempat_usaha")
	umkm.BentukUsaha = c.FormValue("bentuk_usaha")
	ekonomiKreatifStr := c.FormValue("ekonomi_kreatif")
	ekonomiKreatif, err := strconv.ParseBool(ekonomiKreatifStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for ekonomi_kreatif. Must be true or false.",
		})
	}

	umkm.EkonomiKreatif = ekonomiKreatif
	///
	nominalsendiristr := c.FormValue("nominal_sendiri")
	nominalsendiri, err := strconv.ParseFloat(nominalsendiristr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for nominal_sendiri. Must be a decimal number.",
		})
	}

	// Convert float64 to decimal.Decimal
	nominalsendiriDecimal := decimal.NewFromFloat(nominalsendiri)

	// Assign to umkm.NominalSendiri
	umkm.NominalSendiri = nominalsendiriDecimal
	/////
	///ini gaji
	gajistr := c.FormValue("gaji")
	gaji, err := strconv.ParseFloat(gajistr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for gaji. Must be a decimal number.",
		})
	}

	// Convert float64 to decimal.Decimal
	gajiDecimal := decimal.NewFromFloat(gaji)

	// Assign to umkm.NominalSendiri
	umkm.NominalSendiri = gajiDecimal
	/////

	nominalasetstr := c.FormValue("nominal_aset")
	nominalaset, err := strconv.ParseFloat(nominalasetstr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for nominal_sendiri. Must be a decimal number.",
		})
	}

	// Convert float64 to decimal.Decimal
	nominalasetDecimal := decimal.NewFromFloat(nominalaset)

	// Assign to umkm.NominalSendiri
	umkm.NominalAset = nominalasetDecimal

	TenagaKerjaPriastr := c.FormValue("tenaga_kerja_pria")
	tenagaKerjaPria, err := strconv.Atoi(TenagaKerjaPriastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for tenaga_kerja_pria. Must be an integer.",
		})
	}

	umkm.TenagaKerjaPria = tenagaKerjaPria

	//////
	TenagaKerjaWanitastr := c.FormValue("tenaga_kerja_wanita")
	TenagaKerjaWanita, err := strconv.Atoi(TenagaKerjaWanitastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for tenaga_kerja_pria. Must be an integer.",
		})
	}

	umkm.TenagaKerjaWanita = TenagaKerjaWanita
	////
	//////karyawanpria
	karyawanpriastr := c.FormValue("karyawan_pria")
	karyawanpria, err := strconv.Atoi(karyawanpriastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for tenaga_kerja_pria. Must be an integer.",
		})
	}

	umkm.KaryawanPria = karyawanpria
	////
	//////karyawanwanita
	karyawanwanitastr := c.FormValue("karyawan_pria")
	karyawanwanita, err := strconv.Atoi(karyawanwanitastr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for tenaga_kerja_pria. Must be an integer.",
		})
	}

	umkm.KaryawanWanita = karyawanwanita
	////
	
    tanggalMulaiUsahaStr := c.FormValue("tanggal_mulai_usaha")
    layout := "2006-01-02" // Format tanggal yang diharapkan, misalnya YYYY-MM-DD
    tanggalMulaiUsaha, err := time.Parse(layout, tanggalMulaiUsahaStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Invalid value for tanggal_mulai_usaha. Must be in the format YYYY-MM-DD.",
        })
    }
    
    umkm.TanggalMulaiUsaha = tanggalMulaiUsaha
    

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

	omsetData := c.FormValue("omset")
	if omsetData != "" {
		if err := json.Unmarshal([]byte(omsetData), &umkm.Omset); err != nil {
			return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid omset data", nil))
		}
	}
	
	var dokumenIDs []string
	var dokumenFiles []*multipart.FileHeader

	// Ambil dokumen IDs dan dokumen files dari request
	for i := 0; ; i++ {
		// Mengambil dokumen ID berdasarkan indeks
		dokIdKey := fmt.Sprintf("dok_id[%d]", i)
		if id := c.Request().FormValue(dokIdKey); id != "" {
			dokumenIDs = append(dokumenIDs, id)
		} else {
			break // Berhenti jika tidak ada lagi dokumen ID
		}

		// Mengambil dokumen file berdasarkan indeks
		dokUploadKey := fmt.Sprintf("dok_upload[%d]", i)
		if files := c.Request().MultipartForm.File[dokUploadKey]; len(files) > 0 {
			dokumenFiles = append(dokumenFiles, files...)
		} else {
			break // Berhenti jika tidak ada lagi dokumen file
		}
	}

	// Debugging: Print received IDs and files
	fmt.Println("Dokumen IDs:", dokumenIDs)
	fmt.Println("Dokumen Files:", dokumenFiles)

	// Validasi dokumenIDs dan dokumenFiles
	if len(dokumenIDs) == 0 || len(dokumenFiles) == 0 {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "No document files uploaded or IDs provided", nil))
	}

	umkm.Gambar = json.RawMessage([]byte("[]"))

	fmt.Printf("Form Data: %+v\n", umkm)

	// Get authenticated user ID
	userID, err := helper.GetAuthId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "Failed to get user ID", nil))
	}

	// Call service to create UMKM
	result, errSaveKategori := controller.umkmservice.CreateUmkm(*umkm, userID, fileHeaders, dokumenFiles, dokumenIDs)
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

// filter umkm name
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

func (controller *UmkmControllerImpl) GetUmkmListWeb(c echo.Context) error {
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

func (controller *UmkmControllerImpl) GetUmkmId(c echo.Context) error {
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
	deskripsi := c.FormValue("deskripsi")
	sektousaha := c.FormValue("sektor_usaha")
	tempatusaha := c.FormValue("status_tempat_usaha")
	kodeprov := c.FormValue("kode_prov")
	kodekab  := c.FormValue("kode_kab")
	kodekec  := c.FormValue("kode_kec")
	kodekel  := c.FormValue("kode_kelurahan")
	rt      := c.FormValue("rt")
	rw     := c.FormValue("rw")
	kodepos := c.FormValue("kode_pos")
	noNpwd   := c.FormValue("no_npwd")
	bahanbakar := c.FormValue("bahan_bakar")
	kapasitas         := c.FormValue("kapasitas")
	krtiteriausaha := c.FormValue("kriteria_usaha")
	bentukusaha    := c.FormValue("bentuk_usaha")
	nonib          := c.FormValue("nonib")
	

	log.Printf("Form values - Name: %s, NoNpwp: %s, KategoriUmkmId: %s, informasijambuka: %s, ", name, noNpwp, kategoriUmkmID, informasiJamBuka)

	// Ambil file dari form-data jika ada
	files := []*multipart.FileHeader{}
	if file, err := c.FormFile("gambar"); err == nil {
		files = append(files, file)
	} else if err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "failed to get uploaded file", nil))
	}

	//
	ekonomiKreatifStr := c.FormValue("ekonomi_kreatif")
	ekonomiKreatif, err := strconv.ParseBool(ekonomiKreatifStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid value for ekonomi_kreatif. Must be true or false.",
		})
	}

	tenagakerjaPriastr := c.FormValue("tenaga_kerja_pria")
    tenagakerjapria, err := strconv.Atoi(tenagakerjaPriastr)
    if err != nil {
        log.Printf("Error converting tenagakerja: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid tenaga kerja wanita", nil))
    }

	tenagaKerjaWanitastr := c.FormValue("tenaga_kerja_wanita")
    tenagakerjawanita, err := strconv.Atoi(tenagaKerjaWanitastr)
    if err != nil {
        log.Printf("Error converting tenagakerja: %v", err)
        return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid tenaga kerja wanita", nil))
    }

	// Ambil nilai dari form
tanggalMulaiUsahaStr := c.FormValue("tanggal_mulai_usaha")

// Tentukan format tanggal yang benar, misalnya "2006-01-02"
layout := "2006-01-02"

// Konversi string menjadi time.Time dengan urutan parameter yang benar
tanggalMulaiUsaha, err := time.Parse(layout, tanggalMulaiUsahaStr)
if err != nil {
    log.Printf("Error converting tanggal mulai usaha: %v", err)
    return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid tanggal mulai usaha format, must be in YYYY-MM-DD", nil))
}



	// Buat objek request manual
	request := web.Updateumkm{
		Name:                  name,
		NoNpwp:                noNpwp,
		Nama_Penanggung_Jawab: namaPenanggungJawab,
		No_Kontak:             noKontak,
		Lokasi:                lokasi,
		Deskripsi:             deskripsi,
		Kategori_Umkm_Id:      json.RawMessage(kategoriUmkmID),
		Informasi_JamBuka:     json.RawMessage(informasiJamBuka),
		Maps:                  json.RawMessage(maps),
		TenagaKerjaPria: tenagakerjapria,
		TenagaKerjaWanita: tenagakerjawanita,
		EkonomiKreatif: ekonomiKreatif,
		SektorUsaha: sektousaha,
		StatusTempatUsaha: tempatusaha,
		KodeProv: kodeprov,
		KodeKabupaten: kodekab,
		KodeKec: kodekec,
		KodeKel: kodekel,
		KodePos: kodepos,
		Rt: rt,
		Rw: rw,
		NoNpwd: noNpwd,
		BahanBakar: bahanbakar,
		TanggalMulaiUsaha: tanggalMulaiUsaha,
		Kapasitas: kapasitas,
		KriteriaUsaha: krtiteriausaha,
		BentukUsaha: bentukusaha,
		NoNib: nonib,
		// JenisUsaha: jenisusaha,
	}

	// Memanggil service untuk update UMKM
	result, err := controller.umkmservice.UpdateUmkmId(request, umkmid, files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, err.Error(), nil))
	}

	// Response sukses
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "berhasil di update", result))
}

func (controller *UmkmControllerImpl) GetUmmkmList(c echo.Context) error {
	filters, limit, page := helper.ExtractFilter(c.QueryParams())
	kategoriumkm := c.QueryParam("kategori")
	sortOrder := c.QueryParam("sort")

	getUmkm, totalCount, currentPage, totalPages, nextPage, prevPage, errGetUmkmDetail := controller.umkmservice.GetUmkmList(filters, limit, page, kategoriumkm, sortOrder)

	if errGetUmkmDetail != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", errGetUmkmDetail.Error(), model.Pagination{}, nil))
	}

	pagination := model.Pagination{
		CurrentPage:  currentPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalPages:   totalPages,
		TotalRecords: totalCount,
	}

	return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "berhasil", pagination, getUmkm))
}

func (controller *UmkmControllerImpl) GetUmkmListDetial(c echo.Context) error {
	IdUmkm := c.Param("id")
	id, _ := uuid.Parse(IdUmkm)

	// Ambil parameter limit dan page dari query string
	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	// Default nilai untuk limit dan page jika tidak disediakan
	limit := 10
	page := 1

	// Parsing parameter limit jika tersedia
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Parsing parameter page jika tersedia
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Panggil service untuk mendapatkan detail UMKM dengan pagination
	getUmkmDetailList, totalCount, currentPage, totalPages, nextPage, prevPage, errGetUmkmDetailList := controller.umkmservice.GetUmkmDetailList(id, limit, page)

	pagination := model.Pagination{
		CurrentPage:  currentPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalPages:   totalPages,
		TotalRecords: totalCount,
	}

	if errGetUmkmDetailList != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClientpagi(http.StatusInternalServerError, "false", errGetUmkmDetailList.Error(), model.Pagination{}, nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClientpagi(http.StatusOK, "true", "success", pagination, getUmkmDetailList))
}

func (controller *UmkmControllerImpl) DeleteUmkmId(c echo.Context) error {
	// Ambil ID dari URL dan konversi ke UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "Invalid ID format", nil))
	}

	if errDeleteProduk := controller.umkmservice.Delete(id); errDeleteProduk != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errDeleteProduk.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Delete Produk Success", nil))
}

// list acative umkm BackEnd
func (controller *UmkmControllerImpl) ListActveBack(c echo.Context) error {
	getSlider, errGetSlider := controller.umkmservice.GetUmkmActive()

	if errGetSlider != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetSlider.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getSlider))
}

//updateactive

func (conntroller *UmkmControllerImpl) UpdateSldierActive(c echo.Context) error {
	slider := new(web.UpdateActiveUmkm)
	id, _ := uuid.Parse(c.Param("id"))

	if err := c.Bind(slider); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	sliderUpdate, errsliderUpdate := conntroller.umkmservice.UpdateUmkmActive(*slider, id)

	if errsliderUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, errsliderUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "data berhasil diupdate", sliderUpdate))
}

func (controller *UmkmControllerImpl) GetUmkmActive(c echo.Context) error {
	getUmkm, errGetUmkm := controller.umkmservice.GetTestimonialActive()
	if errGetUmkm != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errGetUmkm.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "success", getUmkm))
}
