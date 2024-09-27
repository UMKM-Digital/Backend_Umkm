package umkmservice

import (
	"context"
	"encoding/json"
	"log"
	"os"

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
	hakaksesrepo "umkm/repository/hakakses" // Tambahkan import untuk HakAkses repository
	produkrepo "umkm/repository/produk"
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
	db                 *gorm.DB
}

func NewUmkmService(umkmrepository umkmrepo.CreateUmkm, hakaksesrepository hakaksesrepo.CreateHakakses, db *gorm.DB, produkRepository produkrepo.CreateProduk) *UmkmServiceImpl {
	return &UmkmServiceImpl{
		umkmrepository:     umkmrepository,
		hakaksesrepository: hakaksesrepository,
		produkRepository: produkRepository,
		db:                 db,
	}
}
func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *UmkmServiceImpl) CreateUmkm(umkm web.UmkmRequest, userID int, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
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
		Deskripsi: umkm.Deskripsi,
	}

	saveUmkm, errSaveUmkm := service.umkmrepository.CreateRequest(newUmkm)
	if errSaveUmkm != nil {
		return nil, errSaveUmkm
	}

	hakAkses := domain.HakAkses{
		UserId: userID,
		UmkmId: saveUmkm.IdUmkm,
		Status: 0,
	}
	if err := service.hakaksesrepository.CreateHakAkses(&hakAkses); err != nil {
		return nil, err
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

	products, err := s.produkRepository.GetProductsByUmkmIds(ctx, umkmIDs)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }


    umkmResponses := entity.ToUmkmfilterEntities(umkmList, products) // Convert UMKM entities to responses

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

     // Hapus gambar lama jika ada
	// Hapus gambar lama jika ada
// Hapus gambar lama jika ada
// Hapus gambar lama jika ada
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




    // Simpan gambar baru di folder uploads
	var imageUrls []string
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


	func(services *UmkmServiceImpl) GetUmkmList(filters string, limit int, page int, kategori_umkm string)([]entity.UmkmEntityWebList,int, int, int, *int, *int, error){
		GetTestimonialList, totalCount, currentPage, totalPages, nextPage, prevPage, err := services.umkmrepository.GetUmkmList(filters, limit, page, kategori_umkm)
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


	
// //
// func (service *UmkmServiceImpl) Delete(id uuid.UUID) error {
// 	// Cari UMKM berdasarkan ID
// 	umkm, err := service.umkmrepository.GetUmkmID(id)
// 	if err != nil {
// 		return err
// 	}

// 	// Konversi JSONB ke map[string]interface{}
// 	var gambarURLs []string

// 	gambarMap := make(map[string]interface{})

// 	gambarBytes, err := json.Marshal(umkm.Images)
// 	if err != nil {
// 		return err
// 	}

// 	if err := json.Unmarshal(gambarBytes, &gambarMap); err != nil {
// 		return err
// 	}

// 	// Ambil gambar URLs dari map
// 	if urls, ok := gambarMap["urls"].([]interface{}); ok {
// 		for _, url := range urls {
// 			if urlStr, ok := url.(string); ok {
// 				gambarURLs = append(gambarURLs, urlStr)
// 			}
// 		}
// 	} else {
// 		return errors.New("invalid format for gambar URLs")
// 	}

// 	// Hapus file gambar
// 	for _, gambarURL := range gambarURLs {
// 		filePath := filepath.Join(gambarURL)
// 		filePath = filepath.Clean(filePath)
// 		log.Printf("Attempting to remove file: %s", filePath)

// 		if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 			log.Printf("File does not exist: %s", filePath)
// 			continue
// 		}

// 		if err := os.Remove(filePath); err != nil {
// 			log.Printf("Error removing file %s: %v", filePath, err)
// 			return err
// 		}
// 	}

// 	// Hapus produk yang terkait dengan UMKM
// 	if err := service.produkRepository.DeleteProdukUmkmId(id); err != nil {
// 		return err
// 	}

// 	// Hapus kategori produk yang terkait dengan UMKM
// 	if err := service.kategoriProdukRepository.DeleteKategoriUmkmId(id); err != nil {
// 		return err
// 	}

// 	// Hapus transaksi yang terkait dengan UMKM
// 	if err := service.transaksiRepository.DeleteTransaksiUmkmId(id); err != nil {
// 		return err
// 	}

// 	// Hapus dokumen yang terkait dengan UMKM
// 	if err := service.umkmrepository.DeleteUmkmId(id); err != nil {
// 		return err
// 	}

// 	// Hapus hak akses yang terkait dengan UMKM
// 	if err := service.hakaksesrepository.DeleteUmkmId(id); err != nil {
// 		return err
// 	}

// 	// Hapus umkm dari database
// 	return service.umkmrepository.DeleteUmkmId(id)
// }
