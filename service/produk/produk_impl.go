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

// func (service *ProdukServiceImpl) CreateProduk(produk web.WebProduk, files map[string]*multipart.FileHeader) (map[string]interface{}, error) {
// 	var Gambar domain.JSONB
// 	var savedImageURLs []string

// 	// Handle gambar files
// 	if len(files) > 0 {
// 		for _, file := range files {
// 			ext := filepath.Ext(file.Filename)
// 			randomFileName := generateRandomFileName(ext)
// 			newImagePath := filepath.Join("uploads", randomFileName)

// 			if err := helper.SaveFile(file, newImagePath); err != nil {
// 				return nil, errors.New("failed to save image")
// 			}

// 			savedImageURLs = append(savedImageURLs, newImagePath)
// 		}

// 		Gambar = domain.JSONB{"urls": savedImageURLs}
// 	} else {
// 		Gambar = domain.JSONB{"urls": []string{}}
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
	var Gambar domain.JSONB
	var savedImageURLs []string

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
			savedImageURLs = append(savedImageURLs, filepath.ToSlash(newImagePath))
		}

		Gambar = domain.JSONB{"urls": savedImageURLs}
	} else {
		Gambar = domain.JSONB{"urls": []string{}}
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

func (service *ProdukServiceImpl) UpdateProduk(request web.UpdatedProduk, Id uuid.UUID, file *multipart.FileHeader, indexHapus []int) (map[string]interface{}, error) {
    // Ambil data produk berdasarkan ID
    getProdukById, err := service.produkrepository.FindById(Id)
    if err != nil {
        return nil, err
    }

    // Ambil gambar lama dan konversi dari domain.JSONB ke []string
    var oldGambarList []string
    if urls, ok := getProdukById.Gamabr["urls"].([]interface{}); ok {
        for _, img := range urls {
            if imgStr, ok := img.(string); ok {
                oldGambarList = append(oldGambarList, imgStr)
            }
        }
    }

    // Hapus gambar pada indeks yang diberikan
    toRemove := make(map[int]bool)
    for _, idx := range indexHapus {
        if idx >= 0 && idx < len(oldGambarList) {
            toRemove[idx] = true
        }
    }

    var newGambarList []string
    for i, img := range oldGambarList {
        if !toRemove[i] {
            newGambarList = append(newGambarList, img)
        }
    }

    // Jika ada file gambar baru
    var newGambar []string
    if file != nil {
        // Proses gambar baru
        src, err := file.Open()
        if err != nil {
            return nil, errors.New("failed to open the uploaded file")
        }
        defer src.Close()

        // Menghasilkan nama file acak untuk file yang diunggah
        ext := filepath.Ext(file.Filename)
        randomFileName := generateRandomFileName(ext)
        newImagePath := filepath.Join("uploads", randomFileName)

        // Simpan file ke server
        if err := helper.SaveFile(file, newImagePath); err != nil {
            return nil, errors.New("failed to save image")
        }

        // Tambahkan path baru ke daftar gambar
        newGambar = append(newGambar, filepath.ToSlash(newImagePath))
    }

    // Gabungkan gambar lama yang telah diperbarui dan gambar baru
    combinedGambar := append(newGambarList, newGambar...)

    // Konversi combinedGambar ke JSONB
    newGambarJSONB := domain.JSONB{"urls": combinedGambar}

    // Unmarshal KategoriProduk dari RawMessage ke JSONB
    var kategoriProdukJSONB domain.JSONB
    if len(request.KategoriProduk) > 0 {
        err = json.Unmarshal(request.KategoriProduk, &kategoriProdukJSONB)
        if err != nil {
            return nil, errors.New("failed to unmarshal KategoriProduk")
        }
    } else {
        kategoriProdukJSONB = getProdukById.KategoriProduk
    }

    // Buat objek Produk baru untuk pembaruan
    ProdukRequest := domain.Produk{
        IdUmkm:         Id,
        Nama:           request.Name,
        Gamabr:         newGambarJSONB,
        Harga:          request.Harga,
        Satuan:         request.Satuan,
        Min_pesanan:    request.MinPesanan,
        KategoriProduk: kategoriProdukJSONB,
        Deskripsi:      request.Deskripsi,
    }

    // Update produk
    updatedProduk, errUpdate := service.produkrepository.UpdatedProduk(Id, ProdukRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }

    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "id":            updatedProduk.IdUmkm,
        "name":          updatedProduk.Nama,
        "gambar": map[string]interface{}{
            "urls": combinedGambar,
        },
        "harga":         updatedProduk.Harga,
        "satuan":        updatedProduk.Satuan,
        "min_pesanan":   updatedProduk.Min_pesanan,
        "kategoriproduk": kategoriProdukJSONB,
        "deskripsi":     updatedProduk.Deskripsi,
    }

    return response, nil
}
