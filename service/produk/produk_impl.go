package produkservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"

	// "umkm/model/entity"

	// "umkm/model/entity"
	"umkm/model/web"
	hakaksesrepo "umkm/repository/hakakses"
	produkrepo "umkm/repository/produk"
	umkmrepo "umkm/repository/umkm"

	"os"

	"github.com/google/uuid"
)

type ProdukServiceImpl struct {
	produkrepository produkrepo.CreateProduk
    HakAkses hakaksesrepo.CreateHakakses
    Umkm umkmrepo.CreateUmkm
}

func NewProdukService(produkrepository produkrepo.CreateProduk, HakAkses hakaksesrepo.CreateHakakses,   Umkm umkmrepo.CreateUmkm) *ProdukServiceImpl {
	return &ProdukServiceImpl{
		produkrepository: produkrepository,
       HakAkses: HakAkses,
       Umkm: Umkm,
	}
}

func generateRandomFileName(ext string) string {
	now := time.Now()
	
	// Format the date as YYMMDD
	datePrefix := now.Format("060102") // Format to YYMMDD

	// Generate a new UUID
	uniqueID := uuid.New().String()

	// Include seconds in the file name
	seconds := now.Format("150405") // Format to HHMMSS

	// Combine everything into the final file name
	return fmt.Sprintf("%s-%s-%s%s", datePrefix, uniqueID, seconds, ext)
}

// func (service *ProdukServiceImpl) CreateProduk(produk web.WebProduk, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
// 	var Gambar domain.JSONB
// 	var savedImageURLs []map[string]interface{}

// 	// Handle gambar files
// 	if len(files) > 0 {
// 		for _, file := range files {
// 			ext := filepath.Ext(file.Filename)
// 			randomFileName := generateRandomFileName(ext)
// 			newImagePath := filepath.Join("uploads/Produk", randomFileName)

// 			if err := helper.SaveFile(file, newImagePath); err != nil {
// 				return nil, errors.New("failed to save image")
// 			}

// 			// Save the image path in a format with forward slashes
// 			// savedImageURLs = append(savedImageURLs, filepath.ToSlash(newImagePath))
// 			savedImageURLs = append(savedImageURLs, map[string]interface{}{
// 				"id":     len(savedImageURLs) + 1, // Assuming ID is generated sequentially
// 				"gambar": filepath.ToSlash(newImagePath),
// 			})
// 		}

// 		Gambar = domain.JSONB{"urls": savedImageURLs}
// 	} else {
// 		Gambar = domain.JSONB{"urls": []interface{}	{}}
// 	}

// 	KategoriProduk, err := helper.RawMessageToJSONB(produk.KategoriProduk)
// 	if err != nil {
// 		return nil, errors.New("invalid type for KategoriProduk")
// 	}

// 	newProduk := domain.Produk{
// 		UmkmId:           produk.UmkmId,
// 		Nama:             produk.Name,
// 		Gamabr:           Gambar,
// 		Harga:            produk.Harga,
// 		Satuan:           produk.Satuan,
// 		Min_pesanan:      produk.MinPesanan,
// 		KategoriProduk:   KategoriProduk,
// 		Deskripsi:        produk.Deskripsi,
// 	}

// 	saveProduk, errSaveProduk := service.produkrepository.CreateRequest(newProduk)
// 	if errSaveProduk != nil {
// 		return nil, errSaveProduk
// 	}

// 	return map[string]interface{}{
// 		"umkm_id":            saveProduk.UmkmId,
// 		"nama":               saveProduk.Nama,
// 		"gambar":             saveProduk.Gamabr,
// 		"harga":              saveProduk.Harga,
// 		"satuan":             saveProduk.Satuan,
// 		"minimal_pesanan":    saveProduk.Min_pesanan,
// 		"kategori_produk_id": saveProduk.KategoriProduk,
// 		"deskripsi":          saveProduk.Deskripsi,
// 	}, nil
// }


func (service *ProdukServiceImpl) CreateProduk(produk web.WebProduk, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
    // var Gambar domain.JSONB
    // var savedImageURLs []map[string]interface{}

    // // Handle gambar files
    // if len(files) > 0 {
    //     // Create a slice to preserve order
    //     orderedFiles := make([]*multipart.FileHeader, 0, len(files))
        
    //     // Collect files in the order they are received
    //     for _, file := range files {
    //         orderedFiles = append(orderedFiles, file)
    //     }

    //     for _, file := range orderedFiles {
    //         ext := filepath.Ext(file.Filename)
    //         randomFileName := generateRandomFileName(ext)
    //         newImagePath := filepath.Join("uploads/Produk", randomFileName)

    //         if err := helper.SaveFile(file, newImagePath); err != nil {
    //             return nil, errors.New("failed to save image")
    //         }

    //         imageID := uuid.New().String() // Menghasilkan ID gambar yang unik
    //         newImageEntry := map[string]interface{}{
    //             "id":     imageID, // Menggunakan UUID sebagai ID
    //             "gambar": filepath.ToSlash(newImagePath),
    //         }

    //         // Cek jika gambar sudah ada sebelum menambahkannya
    //         exists := false
    //         for _, img := range savedImageURLs {
    //             if img["gambar"] == newImageEntry["gambar"] {
    //                 exists = true
    //                 break
    //             }
    //         }

    //         if !exists {
    //             savedImageURLs = append(savedImageURLs, newImageEntry)
    //         }
    //     }

    //     Gambar = domain.JSONB{"urls": savedImageURLs}
    // } else {
    //     Gambar = domain.JSONB{"urls": []interface{}{}}
    // }

     // Proses untuk menyimpan gambar baru
     var newImages []map[string]interface{}
     for _, file := range files {
         if file != nil {
             ext := filepath.Ext(file.Filename)
             randomFileName := generateRandomFileName(ext)
             newImagePath := filepath.Join("uploads/Produk", randomFileName)
 
             src, err := file.Open()
             if err != nil {
                 return nil, errors.New("gagal membuka file yang diunggah")
             }
             defer src.Close()
 
             if err := helper.SaveFile(file, newImagePath); err != nil {
                 return nil, errors.New("gagal menyimpan gambar")
             }
 
             // Tambahkan gambar baru dengan ID yang otomatis
             newImage := map[string]interface{}{
                 "id":     len(newImages) + 1, // Menghitung ID baru
                 "gambar": filepath.ToSlash(newImagePath),
             }
             newImages = append(newImages, newImage)
         }
     }
 
     // Persiapkan data gambar JSONB yang diperbarui
     updatedGambarJSONB := domain.JSONB{"urls": newImages}

    KategoriProduk, err := helper.RawMessageToJSONB(produk.KategoriProduk)
    if err != nil {
        return nil, errors.New("invalid type for KategoriProduk")
    }

    newProduk := domain.Produk{
        UmkmId:         produk.UmkmId,
        Nama:           produk.Name,
        Gamabr:         updatedGambarJSONB,
        Harga:          produk.Harga,
        Satuan:         produk.Satuan,
        Min_pesanan:    produk.MinPesanan,
        KategoriProduk: KategoriProduk,
        Deskripsi:      produk.Deskripsi,
    }

    saveProduk, errSaveProduk := service.produkrepository.CreateRequest(newProduk)
    if errSaveProduk != nil {
        return nil, errSaveProduk
    }

    return map[string]interface{}{
        "umkm_id":            saveProduk.UmkmId,
        "nama":               saveProduk.Nama,
        "gambar":             saveProduk.Gamabr,
        "harga":              saveProduk.Harga,
        "satuan":             saveProduk.Satuan,
        "minimal_pesanan":    saveProduk.Min_pesanan,
        "kategori_produk_id": saveProduk.KategoriProduk,
        "deskripsi":          saveProduk.Deskripsi,
    }, nil
}




//
func (service *ProdukServiceImpl) DeleteProduk(id uuid.UUID) error {
	// Cari produk berdasarkan ID
	produk, err := service.produkrepository.FindById(id)
	if err != nil {
		return err
	}

	// Konversi JSONB ke map[string]interface{}
	var gambarURLs []string
	gambarMap := make(map[string]interface{})

	gambarBytes, err := json.Marshal(produk.Gamabr)
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
		// Normalisasi path
		// Jika gambarURL sudah berisi prefix 'uploads/', kita tidak perlu menambahkannya lagi
		filePath := filepath.Join(gambarURL)

		// Normalisasi path untuk memastikan tidak ada folder berulang
		filePath = filepath.Clean(filePath)

		// Log path file yang akan dihapus
		log.Printf("Attempting to remove file: %s", filePath)

		// Cek jika file ada sebelum menghapus
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("File does not exist: %s", filePath)
			continue
		}

		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file %s: %v", filePath, err)
			return err
		}
	}

	// Hapus produk dari database
	return service.produkrepository.DeleteProdukId(id)
}

func(service *ProdukServiceImpl) GetProdukId(id uuid.UUID)(entity.ProdukEntity, error){
	GetProduk, errGetProduk := service.produkrepository.FindById(id)

	if errGetProduk != nil {
		return entity.ProdukEntity{}, errGetProduk
	}

	return entity.ToProdukEntity(GetProduk), nil
}


func (service *ProdukServiceImpl) GetProdukList(Produkid uuid.UUID, filters string, limit int, page int, kategori_produk_id string, sort string) ([]entity.ProdukList, int, int, int, *int, *int, error) {
	getProdukList, totalCount, currentPage, totalPages, nextPage, prevPage, errGetProdukList := service.produkrepository.GetProduk(Produkid, filters, limit, page, kategori_produk_id, sort)
	if errGetProdukList != nil {
		return nil, 0, 0, 0, nil, nil, errGetProdukList
	}

	// Konversi hasil produk ke entitas
	produkEntities := entity.ToProdukEntities(getProdukList)

	return produkEntities,totalCount, currentPage, totalPages, nextPage,prevPage, nil
}

func (service *ProdukServiceImpl) UpdateProduk(request web.UpdatedProduk, id uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil produk berdasarkan ID
    getProdukById, err := service.produkrepository.FindById(id)
    if err != nil {
        return nil, err
    }

    // Ambil gambar yang sudah ada dari produk
    var existingGambar []map[string]interface{}
    if getProdukById.Gamabr["urls"] != nil {
        gambarJSONB, ok := getProdukById.Gamabr["urls"].([]interface{})
        if ok {
            for _, img := range gambarJSONB {
                imgMap, ok := img.(map[string]interface{})
                if ok {
                    existingGambar = append(existingGambar, imgMap) // Simpan gambar yang sudah ada
                }
            }
        }
    }

    // Hapus gambar lama jika ada
    for _, img := range existingGambar {
        oldImagePath := img["gambar"].(string)
        if oldImagePath != "" {
            err := os.Remove(oldImagePath)
            if err != nil {
                return nil, errors.New("gagal menghapus gambar lama")
            }
        }
    }

    // Proses untuk menyimpan gambar baru
    var newImages []map[string]interface{}
    for _, file := range files {
        if file != nil {
            ext := filepath.Ext(file.Filename)
            randomFileName := generateRandomFileName(ext)
            newImagePath := filepath.Join("uploads/Produk", randomFileName)

            src, err := file.Open()
            if err != nil {
                return nil, errors.New("gagal membuka file yang diunggah")
            }
            defer src.Close()

            if err := helper.SaveFile(file, newImagePath); err != nil {
                return nil, errors.New("gagal menyimpan gambar")
            }

            // Tambahkan gambar baru dengan ID yang otomatis
            newImage := map[string]interface{}{
                "id":     len(newImages) + 1, // Menghitung ID baru
                "gambar": filepath.ToSlash(newImagePath),
            }
            newImages = append(newImages, newImage)
        }
    }

    // Persiapkan data gambar JSONB yang diperbarui
    updatedGambarJSONB := domain.JSONB{"urls": newImages}

    // Update produk
    getProdukById.Gamabr = updatedGambarJSONB

    // Kategori
    var kategoriProduk domain.JSONB
    if len(request.KategoriProduk) == 0 {
        kategoriProduk = getProdukById.KategoriProduk // Pakai data lama jika tidak ada perubahan
    } else {
        if err := json.Unmarshal(request.KategoriProduk, &kategoriProduk); err != nil {
            return nil, fmt.Errorf("format kategori tidak valid: %v", err)
        }
    }

    // Buat objek update produk
    ProdukRequest := domain.Produk{
        Nama:           request.Name,
        Gamabr:         getProdukById.Gamabr, // Gambar baru sudah disimpan
        Harga:          request.Harga,
        Satuan:         request.Satuan,
        Min_pesanan:    request.MinPesanan,
        Deskripsi:      request.Deskripsi,
        KategoriProduk: kategoriProduk,
    }

    // Perbarui produk di repository
    updatedProduk, err := service.produkrepository.UpdatedProduk(id, ProdukRequest)
    if err != nil {
        return nil, err
    }

    // Kembalikan data yang diperbarui
    response := map[string]interface{}{
        "name":        updatedProduk.Nama,
        "gambar":      updatedProduk.Gamabr,
        "harga":       updatedProduk.Harga,
        "satuan":      updatedProduk.Satuan,
        "min_pesanan": updatedProduk.Min_pesanan,
        "deskripsi":   updatedProduk.Deskripsi,
        "kategori":    updatedProduk.KategoriProduk,
    }

    return response, nil
}




func (service *ProdukServiceImpl) GetProduk(limit int, page int, filters string, kategoriproduk string, sort string) ([]entity.ProdukWebEntity, int, int, int, *int, *int, error) {
	// Panggil GetProdukList dari repository dengan parameter tambahan
	GetProdukList, totalCount, currentPage, totalPages, nextPage, prevPage, err := service.produkrepository.GetProdukList(limit, page, filters, kategoriproduk, sort)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Konversi hasil query ke dalam response entity
	Produkresponse := entity.ToProdukWebEntities(GetProdukList)

	return Produkresponse, totalCount, currentPage, totalPages, nextPage, prevPage, nil
}


func(service *ProdukServiceImpl) GetProdukWebId(id uuid.UUID)(entity.ProdukWebIdEntity, error){
	GetProduk, errGetProduk := service.produkrepository.FindWebId(id)

	if errGetProduk != nil {
		return entity.ProdukWebIdEntity{}, errGetProduk
	}

	return entity.ToProdukWebIdEntity(GetProduk), nil
}


//get all produk berdasarkan login
func (service *ProdukServiceImpl) GetProdukByUser(userId int) ([]entity.ProdukEntityDetailMobile, error) {
    // Ambil daftar UMKM yang user memiliki hak akses
    umkmIds, err := service.HakAkses.GetUmkmIdsByUserId(userId)
    if err != nil {
        return nil, err
    }
    
    produkList, err := service.produkrepository.GetProdukByUmkmLogin(umkmIds)
    if err != nil {
        return nil, err
    }
    
    var produkDetailList []entity.ProdukEntityDetailMobile
    for _, produk := range produkList {
        // Ambil informasi UMKM berdasarkan umkm_id produk
        umkm, err := service.Umkm.FindById(produk.UmkmId)
        if err != nil {
            return nil, err
        }
        
        // Tambahkan nama UMKM ke entity produk
        produkEntity := entity.ProdukEntityDetailMobile{
            Id:           produk.IdUmkm,
            Gambar:       produk.Gamabr,
            NamaUmkm:     umkm.Name,
            Nama:         produk.Nama,
            Harga:        produk.Harga,
            KategdoriProduk: produk.KategoriProduk,
        }
        
        produkDetailList = append(produkDetailList, produkEntity)
    }

    return produkDetailList, nil
}
