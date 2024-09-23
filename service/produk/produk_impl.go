package produkservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"

	// "umkm/model/entity"

	// "umkm/model/entity"
	"umkm/model/web"
	produkrepo "umkm/repository/produk"

	"os"

	"github.com/google/uuid"
)

type ProdukServiceImpl struct {
	produkrepository produkrepo.CreateProduk
}

func NewProdukService(produkrepository produkrepo.CreateProduk) *ProdukServiceImpl {
	return &ProdukServiceImpl{
		produkrepository: produkrepository,
	}
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *ProdukServiceImpl) CreateProduk(produk web.WebProduk, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
	var Gambar domain.JSONB
	var savedImageURLs []map[string]interface{}

	// Handle gambar files
	if len(files) > 0 {
		for _, file := range files {
			ext := filepath.Ext(file.Filename)
			randomFileName := generateRandomFileName(ext)
			newImagePath := filepath.Join("uploads", randomFileName)

			if err := helper.SaveFile(file, newImagePath); err != nil {
				return nil, errors.New("failed to save image")
			}

			// Save the image path in a format with forward slashes
			// savedImageURLs = append(savedImageURLs, filepath.ToSlash(newImagePath))
			savedImageURLs = append(savedImageURLs, map[string]interface{}{
				"id":     len(savedImageURLs) + 1, // Assuming ID is generated sequentially
				"gambar": filepath.ToSlash(newImagePath),
			})
		}

		Gambar = domain.JSONB{"urls": savedImageURLs}
	} else {
		Gambar = domain.JSONB{"urls": []interface{}	{}}
	}

	KategoriProduk, err := helper.RawMessageToJSONB(produk.KategoriProduk)
	if err != nil {
		return nil, errors.New("invalid type for KategoriProduk")
	}

	newProduk := domain.Produk{
		UmkmId:           produk.UmkmId,
		Nama:             produk.Name,
		Gamabr:           Gambar,
		Harga:            produk.Harga,
		Satuan:           produk.Satuan,
		Min_pesanan:      produk.MinPesanan,
		KategoriProduk:   KategoriProduk,
		Deskripsi:        produk.Deskripsi,
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


func (service *ProdukServiceImpl) GetProdukList(Produkid uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]entity.ProdukList, int, int, int, *int, *int, error) {
	getProdukList, totalCount, currentPage, totalPages, nextPage, prevPage, errGetProdukList := service.produkrepository.GetProduk(Produkid, filters, limit, page, kategori_produk_id)
	if errGetProdukList != nil {
		return nil, 0, 0, 0, nil, nil, errGetProdukList
	}

	// Konversi hasil produk ke entitas
	produkEntities := entity.ToProdukEntities(getProdukList)

	return produkEntities,totalCount, currentPage, totalPages, nextPage,prevPage, nil
}

// func (service *ProdukServiceImpl) UpdateProduk(request web.UpdatedProduk, Id uuid.UUID, file *multipart.FileHeader, indexHapus []string) (map[string]interface{}, error) {
//     // Ambil data produk berdasarkan ID
//     getProdukById, err := service.produkrepository.FindById(Id)
//     if err != nil {
//         return nil, err
//     }

//     // Ambil gambar lama dan konversi dari domain.JSONB ke []map[string]interface{}
//     var oldGambarList []map[string]interface{}
//     var deletedImageID string

//     if urls, ok := getProdukById.Gamabr["urls"].([]interface{}); ok {
//         for _, img := range urls {
//             if imgMap, ok := img.(map[string]interface{}); ok {
//                 imgID, _ := imgMap["id"].(string)
//                 imgPath, _ := imgMap["gambar"].(string)
//                 oldGambarList = append(oldGambarList, map[string]interface{}{"id": imgID, "gambar": imgPath})
//             }
//         }
//     }

//     // Hapus gambar berdasarkan ID yang diberikan
//  // Hapus gambar berdasarkan ID yang diberikan
// var newGambarList []map[string]interface{}
// for _, img := range oldGambarList {
//     imgID := strings.TrimSpace(strings.ToLower(img["id"].(string)))
//     if !containsID(indexHapus, imgID) {
//         newGambarList = append(newGambarList, img)
//     } else {
//         log.Println("Menghapus Gambar dengan ID:", imgID)
//         deletedImageID = imgID // Simpan ID gambar yang dihapus

//         // Hapus file fisik dari sistem
//         filePath := img["gambar"].(string)
//         if err := os.Remove(filePath); err != nil {
//             log.Printf("Gagal menghapus file gambar: %v", err)
//         } else {
//             log.Printf("Berhasil menghapus file gambar: %s", filePath)
//         }
//     }
// }


//     // Jika ada file gambar baru
//     if file != nil {
//         // Proses gambar baru
//         src, err := file.Open()
//         if err != nil {
//             return nil, errors.New("failed to open the uploaded file")
//         }
//         defer src.Close()

//         // Menghasilkan nama file acak untuk file yang diunggah
//         ext := filepath.Ext(file.Filename)
//         randomFileName := generateRandomFileName(ext)
//         newImagePath := filepath.Join("uploads", randomFileName)

//         // Simpan gambar di folder uploads
//         if err := helper.SaveFile(file, newImagePath); err != nil {
//             return nil, errors.New("failed to save image")
//         }

//         // Gunakan kembali ID yang dihapus atau buat ID baru jika tidak ada gambar yang dihapus
//         newImageID := deletedImageID
//         if newImageID == "" {
//             newImageID = generateUniqueImageID() // Buat ID baru hanya jika tidak ada gambar yang dihapus
//         }

//         // Tambahkan gambar baru dengan ID yang sesuai
//         newGambarList = append(newGambarList, map[string]interface{}{
//             "id":    newImageID,
//             "gambar": filepath.ToSlash(newImagePath),
//         })
//     }

//     // Gabungkan gambar lama yang telah diperbarui dan gambar baru
//     combinedGambar := newGambarList

//     // Konversi combinedGambar ke JSONB
//     newGambarJSONB := domain.JSONB{"urls": combinedGambar}

//     // Unmarshal KategoriProduk dari RawMessage ke JSONB
//     var kategoriProdukJSONB domain.JSONB
//     if len(request.KategoriProduk) > 0 {
//         err = json.Unmarshal(request.KategoriProduk, &kategoriProdukJSONB)
//         if err != nil {
//             return nil, errors.New("failed to unmarshal KategoriProduk")
//         }
//     } else {
//         kategoriProdukJSONB = getProdukById.KategoriProduk
//     }

//     // Buat objek Produk baru untuk pembaruan
//     ProdukRequest := domain.Produk{
//         IdUmkm:         Id,
//         Nama:           request.Name,
//         Gamabr:         newGambarJSONB,
//         Harga:          request.Harga,
//         Satuan:         request.Satuan,
//         Min_pesanan:    request.MinPesanan,
//         KategoriProduk: kategoriProdukJSONB,
//         Deskripsi:      request.Deskripsi,
//     }

//     // Update produk
//     updatedProduk, errUpdate := service.produkrepository.UpdatedProduk(Id, ProdukRequest)
//     if errUpdate != nil {
//         return nil, errUpdate
//     }

//     // Bentuk respons yang akan dikembalikan
//     response := map[string]interface{}{
//         "id":            updatedProduk.IdUmkm,
//         "name":          updatedProduk.Nama,
//         "gambar": map[string]interface{}{
//             "urls": combinedGambar,
//         },
//         "harga":         updatedProduk.Harga,
//         "satuan":        updatedProduk.Satuan,
//         "min_pesanan":   updatedProduk.Min_pesanan,
//         "kategoriproduk": kategoriProdukJSONB,
//         "deskripsi":     updatedProduk.Deskripsi,
//     }

//     return response, nil
// }

// // Helper function to check if an ID is in a list
// func containsID(slice []string, item string) bool {
//     for _, v := range slice {
//         if strings.TrimSpace(strings.ToLower(v)) == item {
//             return true
//         }
//     }
//     return false
// }

// // GenerateUniqueImageID generates a new UUID
// func generateUniqueImageID() string {
//     id, err := uuid.NewV6()
//     if err != nil {
//         return "" // Handle error appropriately
//     }
//     return id.String()
// }

// func (service *ProdukServiceImpl) UpdateProduk(request web.UpdatedProduk, id uuid.UUID, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
//     // Ambil data produk berdasarkan ID
//     getProdukById, err := service.produkrepository.FindById(id)
//     if err != nil {
//         return nil, err
//     }

//     var newGambarList []map[string]interface{}
//     var existingGambar []map[string]interface{}

//     // Convert existing gambar from []interface{} to []map[string]interface{}
//     gambarInterfaces, ok := getProdukById.Gamabr["urls"].([]interface{})
//     if !ok {
//         return nil, errors.New("invalid format for existing gambar data")
//     }
//     for _, img := range gambarInterfaces {
//         imgMap, ok := img.(map[string]interface{})
//         if !ok {
//             return nil, errors.New("invalid format for image data")
//         }
//         existingGambar = append(existingGambar, imgMap)
//     }

//     // Process images
//     for imgID, file := range files {
//         // Generate random file name and save the image
//         ext := filepath.Ext(file.Filename)
//         randomFileName := generateRandomFileName(ext)
//         newImagePath := filepath.Join("uploads", randomFileName)

//         src, err := file.Open()
//         if err != nil {
//             return nil, errors.New("failed to open the uploaded file")
//         }
//         defer src.Close()

//         if err := helper.SaveFile(file, newImagePath); err != nil {
//             return nil, errors.New("failed to save image")
//         }

//         // Add or update image record
//         newGambarList = append(newGambarList, map[string]interface{}{
//             "id":     imgID,
//             "gambar": filepath.ToSlash(newImagePath),
//         })
//     }

//     // Update gambar yang telah diubah
//     newGambarJSONB := domain.JSONB{"urls": newGambarList}

//     // Create new product object for update
//     ProdukRequest := domain.Produk{
//         Nama:       request.Name,
//         Gamabr:     newGambarJSONB,
//         Harga:      request.Harga,
//         Satuan:     request.Satuan,
//         Min_pesanan: request.MinPesanan,
//         Deskripsi:  request.Deskripsi,
//     }

//     // Update the product in the repository
//     updatedProduk, err := service.produkrepository.UpdatedProduk(id, ProdukRequest)
//     if err != nil {
//         return nil, err
//     }

//     // Return updated data
//     response := map[string]interface{}{
//         "name":        updatedProduk.Nama,
//         "gambar": map[string]interface{}{
//             "urls": newGambarList,
//         },
//         "harga":       updatedProduk.Harga,
//         "satuan":      updatedProduk.Satuan,
//         "min_pesanan": updatedProduk.Min_pesanan,
//         "deskripsi":   updatedProduk.Deskripsi,
//     }

//     return response, nil
// }




func (service *ProdukServiceImpl) UpdateProduk(request web.UpdatedProduk, id uuid.UUID, files map[int]*multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil produk berdasarkan ID
    getProdukById, err := service.produkrepository.FindById(id)
    if err != nil {
        return nil, err
    }

    // Parse data gambar JSONB yang ada
    var existingGambar []map[string]interface{}
    gambarJSONB, ok := getProdukById.Gamabr["urls"].([]interface{})
    if !ok {
        return nil, errors.New("format gambar yang ada tidak valid")
    }
    for _, img := range gambarJSONB {
        imgMap, ok := img.(map[string]interface{})
        if !ok {
            return nil, errors.New("format data gambar tidak valid")
        }
        existingGambar = append(existingGambar, imgMap)
    }

    // Buat map untuk pencarian gambar berdasarkan ID
    gambarMap := make(map[int]int)
    for idx, imgMap := range existingGambar {
        id, ok := imgMap["id"].(float64) // ID dalam JSONB adalah float64
        if !ok {
            return nil, errors.New("format ID gambar tidak valid")
        }
        gambarMap[int(id)] = idx
    }

    // Log existing gambar
    log.Printf("Existing gambar data: %v", existingGambar)

    // Proses file gambar baru
    for imgID, file := range files {
		
        // Buat nama file baru dan simpan gambar
        ext := filepath.Ext(file.Filename)
        randomFileName := generateRandomFileName(ext)
        newImagePath := filepath.Join("uploads", randomFileName)

        src, err := file.Open()
        if err != nil {
            return nil, errors.New("gagal membuka file yang diunggah")
        }
        defer src.Close()

        if err := helper.SaveFile(file, newImagePath); err != nil {
            return nil, errors.New("gagal menyimpan gambar")
        }

        // Perbarui gambar yang ada atau tambahkan yang baru
        if idx, found := gambarMap[imgID]; found {

			 // Hapus gambar lama jika ada
			 oldImagePath := existingGambar[idx]["gambar"].(string)
			 if oldImagePath != "" {
				 err := os.Remove(oldImagePath)
				 if err != nil {
					 return nil, errors.New("gagal menghapus gambar lama")
				 }
			 }

            // Perbarui gambar yang sudah ada
            log.Printf("Updating existing gambar ID: %d at index: %d", imgID, idx)
            existingGambar[idx]["gambar"] = filepath.ToSlash(newImagePath)
        } else {
            // Tambahkan sebagai gambar baru jika ID tidak ditemukan
            log.Printf("Adding new gambar ID: %d", imgID)
            existingGambar = append(existingGambar, map[string]interface{}{
                "id":     imgID,
                "gambar": filepath.ToSlash(newImagePath),
            })
        }
    }

    // Persiapkan data gambar JSONB yang diperbarui
    updatedGambarJSONB := domain.JSONB{"urls": existingGambar}

		//kategori
	
		    // Kategori
			var existingKategori []string
			kategoriJSONB, ok := getProdukById.KategoriProduk["nama"].([]interface{})
			if ok {
				for _, kat := range kategoriJSONB {
					if katStr, ok := kat.(string); ok {
						existingKategori = append(existingKategori, katStr)
					}
				}
			}
		
			log.Printf("Kategori sebelum update: %v", existingKategori)

			// Ganti kategori pada indeks yang ditentukan
			for idx, newKat := range request.KategoriProduk {
				if idx < len(existingKategori) {
					existingKategori[idx] = newKat // Ganti kategori pada indeks yang sesuai
				} else {
					log.Printf("Kategori baru di luar batas, tidak ada kategori untuk indeks %d", idx)
				}
			}
	  

    // Buat objek update produk
    ProdukRequest := domain.Produk{
        Nama:       request.Name,
        Gamabr:     updatedGambarJSONB,
        Harga:      request.Harga,
        Satuan:     request.Satuan,
        Min_pesanan: request.MinPesanan,
        Deskripsi:  request.Deskripsi,
		KategoriProduk: domain.JSONB{"nama":existingKategori},
    }

    // Perbarui produk di repository
    updatedProduk, err := service.produkrepository.UpdatedProduk(id, ProdukRequest)
    if err != nil {
        return nil, err
    }

    // Kembalikan data yang diperbarui
    response := map[string]interface{}{
        "name":        updatedProduk.Nama,
        "gambar": map[string]interface{}{
            "urls": existingGambar,
        },
        "harga":       updatedProduk.Harga,
        "satuan":      updatedProduk.Satuan,
        "min_pesanan": updatedProduk.Min_pesanan,
        "deskripsi":   updatedProduk.Deskripsi,
		"kategori": existingKategori,
    }

    return response, nil
}




// Helper function to check if an image already exists
func imageExists(existingGambar []map[string]interface{}, imgID string) bool {
    for _, img := range existingGambar {
        if img["id"].(string) == imgID {
            return true
        }
    }
    return false
}
