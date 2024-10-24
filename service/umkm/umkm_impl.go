package umkmservice

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	// "database/sql"
	// "encoding/json"
	"errors"
	"math/rand"
	"mime/multipart"
	"path/filepath"

	// "fmt"
	// "os/user"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	querybuilder "umkm/query_builder"

	// querybuilder "umkm/query_builder"
	dokumenumkmrepo "umkm/repository/dokumenumkm"
	hakaksesrepo "umkm/repository/hakakses" // Tambahkan import untuk HakAkses repository
	masterdokumenlegalrepo "umkm/repository/masterdokumenlegal"
	omsetrepo "umkm/repository/omset"
	produkrepo "umkm/repository/produk"
	transaksirepo "umkm/repository/transaksi"
	umkmrepo "umkm/repository/umkm"

	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UmkmServiceImpl struct {
	umkmrepository     umkmrepo.CreateUmkm
	hakaksesrepository hakaksesrepo.CreateHakakses // Tambahkan field untuk HakAkses repository
	produkRepository produkrepo.CreateProduk
	transaksiRepository transaksirepo.TransaksiRepo
	dokumenrepository dokumenumkmrepo.DokumenUmkmrRepo
	masterdokumen masterdokumenlegalrepo.MasterDokumenLegal
	omsetrepo     omsetrepo.OmsetRepo
	db                 *gorm.DB
}

func NewUmkmService(umkmrepository umkmrepo.CreateUmkm, hakaksesrepository hakaksesrepo.CreateHakakses, db *gorm.DB, produkRepository produkrepo.CreateProduk,transaksiRepository transaksirepo.TransaksiRepo,dokumenrepository dokumenumkmrepo.DokumenUmkmrRepo, masterdokumen masterdokumenlegalrepo.MasterDokumenLegal, omsetrepo  omsetrepo.OmsetRepo) *UmkmServiceImpl {
	return &UmkmServiceImpl{
		umkmrepository:     umkmrepository,
		hakaksesrepository: hakaksesrepository,
		produkRepository: produkRepository,
		transaksiRepository: transaksiRepository,
		dokumenrepository: dokumenrepository,
		masterdokumen: masterdokumen,
		omsetrepo: omsetrepo,
		db:                 db,
	}
}

type Dokumen struct {
    ID        string `json:"id"`
    NamaFile  string `json:"nama_file"`
    Path      string `json:"path"`
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateRandomFileNameDokumen generates a random file name with the format dokumen_YYYYMMDD_randomString_HHMMSS.ext
func generateRandomFileNameDokumen(ext string) string {
    // Dapatkan tanggal saat ini
    now := time.Now()
    // Format tahun, bulan, dan tanggal
    datePrefix := now.Format("20060102") // Format: YYYYMMDD

    // Buat string acak 14 karakter
    randomString := generateRandomString(16)

    // Format jam, menit, dan detik
    timeSuffix := now.Format("150405") // Format: HHMMSS

    // Gabungkan prefix "dokumen", tanggal, angka acak, dan suffix waktu
    return fmt.Sprintf("dokumen_umkm_%s_%s_%s%s", datePrefix, randomString, timeSuffix, ext)
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
    bytes := make([]byte, length)
    for i := range bytes {
        bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(bytes)
}



func (service *UmkmServiceImpl) CreateUmkm(umkm web.UmkmRequest, userID int, files map[string]*multipart.FileHeader, dokumenFiles []*multipart.FileHeader, dokumenIDs []string) (map[string]interface{}, error) {
	kategoriUmkmId, err := helper.RawMessageToJSONB(umkm.Kategori_Umkm_Id)
	if err != nil {
		return nil, errors.New("invalid type for Kategori_Umkm_Id")
	}

	informasiJamBuka, err := helper.RawMessageToJSONB(umkm.Informasi_JamBuka)
	if err != nil {
		return nil, errors.New("invalid type for Informasi_JamBuka")
	}

	Maps, err := helper.RawMessageToJSONB(umkm.Maps)
	if err != nil {
		return nil, errors.New("invalid type for Maps")
	}

	var Images domain.JSONB
	var savedImageURLs []string
	if len(files) > 0 {
		for _, file := range files {
			ext := filepath.Ext(file.Filename)
			randomFileName := generateRandomFileName(ext)
			newImagePath := filepath.Join("uploads", randomFileName)

			if err := helper.SaveFile(file, newImagePath); err != nil {
				return nil, errors.New("failed to save image")
			}

			// Save the image path in a format with forward slashes
			savedImageURLs = append(savedImageURLs, filepath.ToSlash(newImagePath))
		}

		Images = domain.JSONB{"urls": savedImageURLs}
	} else {
		Images = domain.JSONB{"urls": []string{}}
	}

	newUmkm := domain.UMKM{
		Name:                umkm.Name,
		NoNpwp:              umkm.NoNpwp,
		Images:              Images,
		KategoriUmkmId:      kategoriUmkmId,
		NamaPenanggungJawab: umkm.Nama_Penanggung_Jawab,
		InformasiJambuka:    informasiJamBuka,
		NoKontak:            umkm.No_Kontak,
		Lokasi:              umkm.Lokasi,
		Maps:                Maps,
		Active: 0,
		Deskripsi: umkm.Deskripsi,
		Bentukusaha: umkm.BentukUsaha,
		KodeProv: umkm.KodeProv,
		KodeKabupaten: umkm.KodeKabupaten,
		KodeKecamatan: umkm.KodeKec,
		KodeKelurahan: umkm.KodeKel,
		KodePos: umkm.KodePos,
		RT: umkm.Rt,
		Rw: umkm.Rw,
		BahanBakar: umkm.BahanBakar,
		TanggalMulaiUsaha: umkm.TanggalMulaiUsaha,
		Kapasitas: umkm.Kapasitas,
		TenagaKerjaPria: umkm.TenagaKerjaPria,
		TenagaKerjaWanita: umkm.TenagaKerjaWanita,
		NominalAset: umkm.NominalAset,
		NominalSendiri: umkm.NominalSendiri,
		EkonomiKreatif: umkm.EkonomiKreatif,
		KriteriaUsaha: umkm.KriteriaUsaha,
		NoNpwd: umkm.NoNpwd,
		NoNib: umkm.NoNib,
		SektorUsaha: umkm.SektorUsaha,
		JenisUsaha: umkm.JenisUsaha,
		BentukUsaha: umkm.BentukUsaha,
		StatusTempatUsaha: umkm.StatusTempatUsaha,
	}

	saveUmkm, errSaveUmkm := service.umkmrepository.CreateRequest(newUmkm)
	if errSaveUmkm != nil {
		return nil, errSaveUmkm
	}

	hakAkses := domain.HakAkses{
		UserId: userID,
		UmkmId: saveUmkm.IdUmkm,
		Status: domain.StatusEnum("menunggu"),
	}
	if err := service.hakaksesrepository.CreateHakAkses(&hakAkses); err != nil {
		return nil, err
	}


    // **Simpan Omset**
	for _, omsetReq := range umkm.Omset {
		omset := domain.Omset{
			Bulan:       omsetReq.Bulan,
			Nominal: omsetReq.JumlahOmset,
			UmkmId:      saveUmkm.IdUmkm, // Gunakan UmkmId dari UMKM yang baru dibuat
		}
	
		// Simpan setiap omset dan tangkap kedua nilai kembalian
		_, err := service.omsetrepo.CreateRequest(omset)
		if err != nil {
			return nil, err
		}
	
		// Optional: kamu bisa menggunakan `savedOmset` untuk kebutuhan lain
	}

	
// **Mengelola dan Simpan Dokumen**
  // Mengelola dan Simpan Dokumen
  var dokumenList []Dokumen
  uploadDir := "uploads/files"

  // Create directory if it doesn't exist
  if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
	  if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
		  return nil, errors.New("failed to create directory for documents")
	  }
  }

  // Loop through uploaded documents
  for i, file := range dokumenFiles {
	  fileID := uuid.New().String()
	  fileName := file.Filename
	  fileExt := filepath.Ext(fileName)
	  savePath := fmt.Sprintf("%s/%s%s", uploadDir, fileID, fileExt)

	  // Save document file
	  if err := helper.SaveFile(file, savePath); err != nil {
		  return nil, fmt.Errorf("failed to save document file %s: %w", fileName, err)
	  }

	  // Create document entry
	  dokumen := Dokumen{
		  ID:       fileID,
		  NamaFile: fileName,
		  Path:     savePath,
	  }

	  // Add document to the list
	  dokumenList = append(dokumenList, dokumen)

	  // Convert dokumen ID to int
	  dokumenID, err := strconv.Atoi(dokumenIDs[i])
	  if err != nil {
		  return nil, fmt.Errorf("failed to convert document ID: %w", err)
	  }

	  // Save each document to the database
	  dokumenEntity := domain.UmkmDokumen{
		  UmkmId:        saveUmkm.IdUmkm,
		  DokumenId:     dokumenID, // Pastikan ID ini sesuai dengan yang dibutuhkan
		  DokumenUpload: domain.JSONB{"dokumen_list": dokumenList},
	  }

	  if _, err := service.dokumenrepository.CreateRequest(dokumenEntity); err != nil {
		  return nil, err
	  }
  }



	return map[string]interface{}{
		"name":                  saveUmkm.Name,
		"kategori_umkm_id":      saveUmkm.KategoriUmkmId,
		"nama_penanggung_jawab": saveUmkm.NamaPenanggungJawab,
		"informasi_jam":         saveUmkm.InformasiJambuka,
		"no_kontak":             saveUmkm.NoKontak,
		"lokasi":                saveUmkm.Lokasi,
		"images":                saveUmkm.Images,
		"user_id":               userID,
		"status":                hakAkses.Status,
		"deskripsi": saveUmkm.Deskripsi,
	}, nil
}

// func (s *UmkmServiceImpl) GetUmkmListByUserId(ctx context.Context, userId int, filters string, limit int, page int) (map[string]interface{}, error) {
//     // Mendapatkan Hak Akses berdasarkan user ID
//     hakAksesList, err := s.hakaksesrepository.GetHakAksesByUserId(ctx, userId)
//     if err != nil {
//         return nil, err
//     }

//     // Membuat slice untuk menampung UMKM IDs dari Hak Akses
//     var umkmIDs []uuid.UUID
//     for _, hakAkses := range hakAksesList {
//         umkmIDs = append(umkmIDs, hakAkses.UmkmId)
//     }

//     // Mengambil daftar UMKM berdasarkan UMKM IDs dengan pagination
//     umkmList, totalCount, currentPage, totalPages, nextPage, prevPage, err := s.umkmrepository.GetUmkmListByIds(ctx, umkmIDs, filters, limit, page)
//     if err != nil {
//         return nil, err
//     }

//     // Mengonversi UMKM list ke entitas yang sesuai untuk response
//     umkmEntitiesList, err := entity.ToUmkmEntities(umkmList, s.db)
//     if err != nil {
//         return nil, err
//     }

//     // Mengembalikan hasil dalam format map, termasuk pagination detail
//     result := map[string]interface{}{
//         "total_records": totalCount,
//         "current_page":  currentPage,
//         "total_pages":   totalPages,
//         "next_page":     nextPage,
//         "prev_page":     prevPage,
//         "umkm_list":     umkmEntitiesList,
//     }

//     return result, nil
// }

func (s *UmkmServiceImpl) GetUmkmListByUserId(ctx context.Context, userId int, filters string, limit int, page int) ([]entity.UmkmFilterEntity, int, int, int, *int, *int, error) {
    // Fetch HakAkses and UMKM IDs
    hakAksesList, err := s.hakaksesrepository.GetHakAksesByUserId(ctx, userId)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    var umkmIDs []uuid.UUID
    for _, hakAkses := range hakAksesList {
        umkmIDs = append(umkmIDs, hakAkses.UmkmId)
    }

    // Fetch UMKM entities based on IDs
    umkmList, totalCount, currentPage, totalPages, nextPage, prevPage, err := s.umkmrepository.GetUmkmListByIds(ctx, umkmIDs, filters, limit, page)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }
	
	// Fetch UmkmDokumen by UMKM IDs
    umkmDokumenList, err := s.dokumenrepository.GetUmkmDokumenByUmkmIds(ctx, umkmIDs)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Fetch MasterDokumenLegal
    masterDokumenLegalList, err := s.masterdokumen.GetAllMasterDokumenLegal(ctx)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

	products, err := s.produkRepository.GetProductsByUmkmIds(ctx, umkmIDs)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }


    umkmResponses := entity.ToUmkmfilterEntities(umkmList, products, umkmDokumenList, masterDokumenLegalList) // Convert UMKM entities to responses

    return umkmResponses, totalCount, currentPage, totalPages, nextPage, prevPage, nil
}


func (service *UmkmServiceImpl) GetUmkmFilter(ctx context.Context, userID int, filters map[string]string, allowedFilters []string) ([]entity.UmkmFilterEntity, error) {
	queryBuilder := querybuilder.NewBaseQueryBuilderName(service.db)

	query, err := queryBuilder.GetQueryBuilderName(filters, allowedFilters)
	if err != nil {
		return nil, err
	}

	hakakseslist, err := service.hakaksesrepository.GetHakAksesByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	var umkmIds []uuid.UUID
	for _, hakakses := range hakakseslist {
		umkmIds = append(umkmIds, hakakses.UmkmId)
	}

	var umkmList []entity.UmkmFilterEntity
	result := query.Find(&umkmList)
	if result.Error != nil {
		return nil, result.Error
	}

	return umkmList, nil
}

func (service *UmkmServiceImpl) GetUmkmListWeb(ctx context.Context, userId int) ([]entity.UmkmEntityList, error) {
	hakAksesList, err := service.hakaksesrepository.GetHakAksesByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Collect UMKM IDs from HakAkses
	var umkmIDs []uuid.UUID
	for _, hakAkses := range hakAksesList {
		umkmIDs = append(umkmIDs, hakAkses.UmkmId)
	}

	// Fetch UMKM entities based on the collected IDs
	umkmList, err := service.umkmrepository.GetUmkmListWeb(ctx, umkmIDs)
	if err != nil {
		return nil, err
	}

	// Convert UMKM entities to UmkmEntity format
	return entity.ToUmkmListEntities(umkmList), nil
}


func(service *UmkmServiceImpl) GetUmkmId(id uuid.UUID)(entity.UmkmEntity, error){
	GetUmkm, errGetUmkm := service.umkmrepository.GetUmkmID(id)

	if errGetUmkm != nil {
		return entity.UmkmEntity{}, errGetUmkm
	}

	return entity.ToUmkmEntity(GetUmkm), nil
}

func (service *UmkmServiceImpl) UpdateUmkmId(request web.Updateumkm, umkmid uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil data UMKM berdasarkan ID
    getUmkmById, err := service.umkmrepository.GetUmkmID(umkmid)
    if err != nil {
        return nil, err
    }

// Hapus gambar lama jika ada gambar baru yang diunggah
if len(files) > 0 {
    if images, ok := getUmkmById.Images["urls"].([]interface{}); ok {
        for _, img := range images {
            // Cek apakah gambar bertipe string
            if imgPath, ok := img.(string); ok {
                oldImagePath := imgPath // Pastikan path gambar tidak mengandung "uploads/" lagi
                log.Printf("Attempting to remove old image: %s", oldImagePath) // Log path yang ingin dihapus

                // Cek apakah file ada sebelum dihapus
                if _, err := os.Stat(oldImagePath); err == nil {
                    if err := os.Remove(oldImagePath); err != nil {
                        log.Printf("Failed to remove old image %s: %v", oldImagePath, err)
                    } else {
                        log.Printf("Successfully removed old image: %s", oldImagePath)
                    }
                } else {
                    log.Printf("Image does not exist: %s", oldImagePath)
                }
            } else {
                log.Printf("Invalid image type: %v", img)
            }
        }
    } else {
        log.Printf("Failed to parse images from UMKM")
    }
}
// // Hapus gambar lama jika ada
// if images, ok := getUmkmById.Images["urls"].([]interface{}); ok {
//     for _, img := range images {
//         // Cek apakah gambar bertipe string
//         if imgPath, ok := img.(string); ok {
//             oldImagePath := imgPath // Pastikan path gambar tidak mengandung "uploads/" lagi
//             log.Printf("Attempting to remove old image: %s", oldImagePath) // Log path yang ingin dihapus

//             // Cek apakah file ada sebelum dihapus
//             if _, err := os.Stat(oldImagePath); err == nil {
//                 if err := os.Remove(oldImagePath); err != nil {
//                     log.Printf("Failed to remove old image %s: %v", oldImagePath, err)
//                 } else {
//                     log.Printf("Successfully removed old image: %s", oldImagePath)
//                 }
//             } else {
//                 log.Printf("Image does not exist: %s", oldImagePath)
//             }
//         } else {
//             log.Printf("Invalid image type: %v", img)
//         }
//     }
// } else {
//     log.Printf("Failed to parse images from UMKM")
// }

    // Simpan gambar baru di folder uploads
	var imageUrls []string
	// for _, file := range files {
	// 	// Mendapatkan ekstensi file
	// 	ext := filepath.Ext(file.Filename)
	// 	// Menghasilkan nama file acak dengan ekstensi yang sesuai
	// 	filename := generateRandomFileName(ext)
	// 	filePath := fmt.Sprintf("uploads/%s", filename)
	
	// 	// Gunakan helper untuk menyimpan file
	// 	if err := helper.SaveFile(file, filePath); err != nil {
	// 		return nil, err
	// 	}
	
	// 	// Tambahkan path gambar baru ke array imageUrls
	// 	imageUrls = append(imageUrls, filePath)  // Format: uploads/filename.jpg
	// }
	


if len(files) > 0 {
    // Ada file baru yang diunggah
    for _, file := range files {
        // Mendapatkan ekstensi file
        ext := filepath.Ext(file.Filename)
        // Menghasilkan nama file acak dengan ekstensi yang sesuai
        filename := generateRandomFileName(ext)
        filePath := fmt.Sprintf("uploads/%s", filename)

        // Gunakan helper untuk menyimpan file
        if err := helper.SaveFile(file, filePath); err != nil {
            return nil, err
        }

        // Tambahkan path gambar baru ke array imageUrls
        imageUrls = append(imageUrls, filePath)  // Format: uploads/filename.jpg
    }
} else {
    // Tidak ada file baru yang diunggah, gunakan gambar lama
    if images, ok := getUmkmById.Images["urls"].([]interface{}); ok {
        for _, img := range images {
            if imgPath, ok := img.(string); ok {
                imageUrls = append(imageUrls, imgPath)
            }
        }
    }
}

 // Mengupdate KategoriUmkmId
	var kategoriUmkm domain.JSONB
	if len(request.Kategori_Umkm_Id) == 0 {
		kategoriUmkm = getUmkmById.KategoriUmkmId // Pakai data lama jika tidak ada perubahan
	} else {
		if err := json.Unmarshal(request.Kategori_Umkm_Id, &kategoriUmkm); err != nil {
			return nil, fmt.Errorf("format kategori_umkm_id tidak valid: %v", err)
		}
	}

    // Tentukan informasi_jambuka (gunakan data lama jika input kosong)
    var informasiJamBuka domain.JSONB
    if len(request.Informasi_JamBuka) == 0 {
        informasiJamBuka = getUmkmById.InformasiJambuka // Pakai data lama jika tidak ada perubahan
    } else {
        if err := json.Unmarshal(request.Informasi_JamBuka, &informasiJamBuka); err != nil {
            return nil, fmt.Errorf("format informasi_jambuka tidak valid: %v", err)
        }
    }

    // Tentukan maps (gunakan data lama jika input kosong)
    var maps domain.JSONB
    if len(request.Maps) == 0 {
        maps = getUmkmById.Maps // Pakai data lama jika tidak ada perubahan
    } else {
        if err := json.Unmarshal(request.Maps, &maps); err != nil {
            return nil, fmt.Errorf("format maps tidak valid: %v", err)
        }
    }

    // Buat data produk yang akan diperbarui
    produkRequest := domain.UMKM{
        Name:                request.Name,
        NoNpwp:              request.NoNpwp,
        NamaPenanggungJawab: request.Nama_Penanggung_Jawab,
        NoKontak:            request.No_Kontak,
        Lokasi:              request.Lokasi,
		KategoriUmkmId:      kategoriUmkm, // Update KategoriUmkmId di sini
		InformasiJambuka: informasiJamBuka,
		Deskripsi: request.Deskripsi,
        Images: map[string]interface{}{
            "urls": imageUrls,  // Format gambar baru yang disimpan
        },
    }

    // Update data UMKM di repository
    updatedUmkm, err := service.umkmrepository.UpdateUmkmId(umkmid, produkRequest) // Tangkap dua nilai
    if err != nil {
        return nil, err
    }

    // Return response sukses
    response := map[string]interface{}{
        "code":    200,
        "message": "UMKM berhasil diperbarui",
        "data":    updatedUmkm, // Kembalikan hasil update
    }

    return response, nil
}


	func(services *UmkmServiceImpl) GetUmkmList(filters string, limit int, page int, kategori_umkm string, sortOrder string)([]entity.UmkmEntityWebList,int, int, int, *int, *int, error){
		GetTestimonialList, totalCount, currentPage, totalPages, nextPage, prevPage, err := services.umkmrepository.GetUmkmList(filters, limit, page, kategori_umkm, sortOrder)
		if err != nil {
			return nil, 0, 0, 0, nil, nil, err
		}
		umkmEntites := entity.ToUmkmEntitiesWebList(GetTestimonialList)

		return umkmEntites,totalCount, currentPage, totalPages, nextPage,prevPage, nil
	}


	func(service *UmkmServiceImpl) GetUmkmDetailList(id uuid.UUID, limit int, page int)([]entity.UmkmDetailEntity,int, int, int, *int, *int, error){
		GetTestimonialList,totalCount, currentPage, totalPages, nextPage, prevPage, err := service.umkmrepository.GetUmkmListDetailPaginated(id, limit, page)
		if err != nil {
			return nil, 0, 0, 0, nil, nil, err
		}
		umkmDetial := entity.ToUmkmEntitiesDetailList(GetTestimonialList)

		return umkmDetial,totalCount, currentPage, totalPages, nextPage,prevPage, nil
	}


	
	func (service *UmkmServiceImpl) Delete(id uuid.UUID) error {
		// Cari UMKM berdasarkan ID
		umkm, err := service.umkmrepository.GetUmkmID(id)
		if err != nil {
			return err
		}
	
		// Konversi JSONB ke map[string]interface{} untuk mengakses URL gambar
		var gambarURLs []string
		gambarMap := make(map[string]interface{})
	
		gambarBytes, err := json.Marshal(umkm.Images)
		if err != nil {
			return err
		}
	
		if err := json.Unmarshal(gambarBytes, &gambarMap); err != nil {
			return err
		}
	
		// Ambil gambar URLs dari map
		if urls, ok := gambarMap["urls"].([]interface{}); ok {
			for _, url := range urls {
				if urlStr, ok := url.(string); ok {
					gambarURLs = append(gambarURLs, urlStr)
				}
			}
		} else {
			return errors.New("invalid format for gambar URLs")
		}
	
		// Hapus file gambar
		for _, gambarURL := range gambarURLs {
			filePath := filepath.Clean(gambarURL)
			log.Printf("Attempting to remove file: %s", filePath)
	
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				log.Printf("File does not exist: %s", filePath)
				continue
			}
	
			if err := os.Remove(filePath); err != nil {
				log.Printf("Error removing file %s: %v", filePath, err)
				return err
			}
		}


		produkList, err := service.produkRepository.GetProdukByUmkmId(id)
		if err != nil {
			return err
		}
	 
		 // Hapus gambar dari setiap produk
		 for _, produk := range produkList {
			 var gambarProdukURLs []string
			 gambarProdukMap := make(map[string]interface{})
	 
			 gambarBytes, err := json.Marshal(produk.Gamabr)
			 if err != nil {
				 return err
			 }
	 
			 if err := json.Unmarshal(gambarBytes, &gambarProdukMap); err != nil {
				 return err
			 }
	 
			 // Ambil gambar URLs dari map untuk produk
			 if urls, ok := gambarProdukMap["urls"].([]interface{}); ok {
				 for _, url := range urls {
					 if urlMap, ok := url.(map[string]interface{}); ok {
						 if gambarURL, ok := urlMap["gambar"].(string); ok {
							 gambarProdukURLs = append(gambarProdukURLs, gambarURL)
						 }
					 }
				 }
			 } else {
				 return errors.New("invalid format for produk gambar URLs")
			 }

			    // Hapus file gambar produk
				for _, gambarURL := range gambarProdukURLs {
					filePath := filepath.Clean(gambarURL)
					log.Printf("Attempting to remove product file: %s", filePath)
		
					if _, err := os.Stat(filePath); os.IsNotExist(err) {
						log.Printf("File does not exist: %s", filePath)
						continue
					}
		
					if err := os.Remove(filePath); err != nil {
						log.Printf("Error removing product file %s: %v", filePath, err)
						return err
					}
				}
		
				// Hapus produk dari database
			}
		

			///
			// Hapus dokumen dari folder dan database
// Hapus dokumen dari folder dan database
// Hapus dokumen dari folder dan database
// Hapus dokumen dari folder
dokumenList, err := service.dokumenrepository.GetDokumnByUmkmId(id)
if err != nil {
    return err
}

// Hapus file dokumen dari setiap dokumen
for _, dokumen := range dokumenList {
    // Mendapatkan path file dokumen dari JSONB
    var dokumenUpload map[string]interface{}
    gambarBytes, err := json.Marshal(dokumen.DokumenUpload)
    if err != nil {
        return err
    }

    // Unmarshal untuk mendapatkan struktur yang benar
    if err := json.Unmarshal(gambarBytes, &dokumenUpload); err != nil {
        return err
    }

    // Debugging: Log isi dokumenUpload
    log.Printf("dokumenUpload content: %+v", dokumenUpload)

    // Ambil array dokumen_list
    dokumenListArray, ok := dokumenUpload["dokumen_list"].([]interface{})
    if !ok {
        return errors.New("invalid format for dokumen_list")
    }

    // Iterasi melalui setiap dokumen dalam dokumen_list
    for _, doc := range dokumenListArray {
        docMap, ok := doc.(map[string]interface{})
        if !ok {
            return errors.New("invalid format for dokumen in dokumen_list")
        }

        // Ambil path dari dokumen
        filePath, ok := docMap["path"].(string)
        if !ok {
            return errors.New("invalid format for dokumen path")
        }

        log.Printf("Attempting to remove document file: %s", filePath)

        // Cek apakah file ada
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            log.Printf("File does not exist: %s", filePath)
            continue
        }

        // Hapus file dokumen
        if err := os.Remove(filePath); err != nil {
            log.Printf("Error removing document file %s: %v", filePath, err)
            return err
        }
    }

    // Di sini tidak ada penghapusan dari database, hanya hapus file
}





	
		// Hapus produk yang terkait dengan UMKM
		if err := service.produkRepository.DeleteProdukUmkmId(id); err != nil {
			return err
		}

		
	
		// Hapus transaksi yang terkait dengan UMKM
		if err := service.transaksiRepository.DeleteTransaksiUmkmId(id); err != nil {
			return err
		}
	
		// Hapus hak akses yang terkait dengan UMKM
		if err := service.hakaksesrepository.DeleteUmkmId(id); err != nil {
			return err
		}
	
		// Hapus dokumen UMKM dari database
		if err := service.dokumenrepository.DeleteDokumenUmkmId(id); err != nil {
			return err
		}
		if err := service.dokumenrepository.DeleteDokumenUmkmId(id); err != nil {
			return err
		}
		if err := service.umkmrepository.DeleteUmkmId(id); err != nil {
			return err
		}
	
		return nil
	}
	
	//list Top Umkm Back
	func (service *UmkmServiceImpl) GetUmkmActive() ([]entity.UmkmActive, error) {
		GetUmkmActiveList, err := service.umkmrepository.ListUmkmActiceBack()
		if err != nil {
			return nil, err
		}
		return entity.ToUmkmListEntitiesActive(GetUmkmActiveList), nil
	}

	//update active
	func (service *UmkmServiceImpl) UpdateUmkmActive(request web.UpdateActiveUmkm, Id uuid.UUID) (map[string]interface{}, error) {
  
		getTestimoniUmkmActive, err := service.umkmrepository.GetUmkmID(Id)
		if err != nil {
			return nil, err
		}
	
	   
		if request.Active == getTestimoniUmkmActive.Active {
		   
			response := map[string]interface{}{
				"active": getTestimoniUmkmActive.Active,
			}
			return response, nil
		}
	
	   
		errUpdate := service.umkmrepository.UpdateActiveId(Id, request.Active)
		if errUpdate != nil {
			return nil, errUpdate
		}
	
	  
		response := map[string]interface{}{
			"active": request.Active,
		}
		return response, nil
	}


	func (service *UmkmServiceImpl) GetTestimonialActive() ([]entity.UmkmActive, error) {
		GetTestimonialList, err := service.umkmrepository.GetUmkmActive(1)
		if err != nil {
			return nil, err
		}
		if len(GetTestimonialList) == 0 {
			log.Println("No umkm found with active = 1")
			return nil, nil
		}
		return entity.ToUmkmListEntitiesActive(GetTestimonialList), nil
	}