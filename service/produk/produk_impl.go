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

// func(service *ProdukServiceImpl)GetProdukId(id uuid.UUID)(entity.ProdukEntity, error){
// 	GetProduk, errGetProduk := service.produkrepository.ProdukById(id)

// 	if errGetProduk != nil {
// 		return entity.ProdukEntity{}, errGetProduk
// 	}

// 	return entity.ToProdukEntity(GetProduk), nil
// }

func (service *ProdukServiceImpl) GetProdukList(Produkid uuid.UUID, filters string, limit int, page int, kategori_produk_id string) ([]entity.ProdukList, error) {
	// getProdukList, err := service.produkrepository.GetProduk(Produkid)

    // if err != nil {
    //     return nil, err
    // }

	getProdukList, errGetProdukList := service.produkrepository.GetProduk(Produkid,filters,limit,page, kategori_produk_id)

	if errGetProdukList != nil {
		return []entity.ProdukList{}, errGetProdukList
	}

	return entity.ToProdukEntities(getProdukList),nil
}