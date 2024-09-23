package dokumenumkmservice

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"umkm/helper" // Pastikan package helper di-import
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	dokumenumkmrepo "umkm/repository/dokumenumkm"

	"github.com/google/uuid"
)

type DokumenUmkmServiceImpl struct {
	dokumenumkmrepository dokumenumkmrepo.DokumenUmkmrRepo
}

func NewDokumenUmkmService(dokumenumkmrepository dokumenumkmrepo.DokumenUmkmrRepo) *DokumenUmkmServiceImpl {
	return &DokumenUmkmServiceImpl{
		dokumenumkmrepository: dokumenumkmrepository,
	}
}

type Dokumen struct {
    ID        string `json:"id"`
    NamaFile  string `json:"nama_file"`
    Path      string `json:"path"`
}

func (service *DokumenUmkmServiceImpl) CreateDokumenUmkm(produk web.CreateUmkmDokumenLegal, files []*multipart.FileHeader) (map[string]interface{}, error) {
    var dokumenList []Dokumen
    uploadDir := "uploads/files"

    // Buat direktori jika belum ada
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
            log.Printf("Failed to create directory: %v", err)
            return nil, fmt.Errorf("failed to create directory: %w", err)
        }
    }

    // Loop melalui file yang diunggah
    for _, file := range files {
        fileID := uuid.New().String()
        fileName := file.Filename
        fileExt := filepath.Ext(fileName)
        savePath := fmt.Sprintf("%s/%s%s", uploadDir, fileID, fileExt)

        // Simpan file di direktori uploads
        if err := helper.SaveFile(file, savePath); err != nil {
            log.Printf("Failed to save file %s: %v", fileName, err)
            return nil, fmt.Errorf("failed to save file %s: %w", fileName, err)
        }

        // Buat data dokumen
        dokumen := Dokumen{
            ID:       fileID,
            NamaFile: fileName,
            Path:     savePath,
        }

        // Tambahkan dokumen ke list
        dokumenList = append(dokumenList, dokumen)
    }

    // Membungkus array dokumen dalam objek JSON
    dokumenWrapper := map[string]interface{}{
        "dokumen_list": dokumenList,
    }

    // Serialisasi dokumen ke JSON
    dokumenJSON, err := json.Marshal(dokumenWrapper)
    if err != nil {
        log.Printf("Failed to marshal dokumenList: %v", err)
        return nil, fmt.Errorf("failed to marshal dokumenList: %w", err)
    }

    // Konversi []byte ke domain.JSONB
    dokumenJSONB := domain.JSONB{}
    if err := json.Unmarshal(dokumenJSON, &dokumenJSONB); err != nil {
        log.Printf("Failed to unmarshal JSON to domain.JSONB: %v", err)
        return nil, fmt.Errorf("failed to unmarshal JSON to domain.JSONB: %w", err)
    }

    // Update dokumen ke produk
    dokumenEntity := domain.UmkmDokumen{
        DokumenId:     produk.DokumenId,
        UmkmId:        produk.UmkmId,
        DokumenUpload: dokumenJSONB, // Simpan dalam format JSONB
    }

    // Simpan ke database
    if _, err := service.dokumenumkmrepository.CreateRequest(dokumenEntity); err != nil {
        log.Printf("Failed to create dokumen in repository: %v", err)
        return nil, fmt.Errorf("failed to create dokumen in repository: %w", err)
    }

    // Siapkan respons
    response := map[string]interface{}{
        "umkm_id":    produk.UmkmId,
        "dokumen_id": produk.DokumenId,
        "dok_upload": dokumenList,  // Dokumen yang berhasil diunggah
    }

    return response, nil
}

func(service *DokumenUmkmServiceImpl) GetDokumenUmkmId(id int, umkm uuid.UUID)(entity.DokumenLegalEntity, error){
    getDokumenUmkm, errDokumenUmkm := service.dokumenumkmrepository.GetId(id, umkm)

    if errDokumenUmkm != nil {
		return entity.DokumenLegalEntity{}, errDokumenUmkm
	}
    return entity.ToDokumenLegalEntity(getDokumenUmkm), nil
}


func (service *DokumenUmkmServiceImpl) UpdateDokumenUmkm(id int, umkmId uuid.UUID, files []*multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil dokumen yang ada berdasarkan ID dan umkmId
    existingDokumen, err := service.dokumenumkmrepository.GetId(id, umkmId)
    if err != nil {
        log.Printf("Failed to retrieve dokumen: %v", err)
        return nil, fmt.Errorf("dokumen not found: %w", err)
    }

    // Direktori penyimpanan
    uploadDir := "uploads/files"

    // Buat direktori jika belum ada
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
            log.Printf("Failed to create directory: %v", err)
            return nil, fmt.Errorf("failed to create directory: %w", err)
        }
    }

    var dokumenList []Dokumen

    // Hapus file lama jika ada
    if existingDokumen.DokumenUpload != nil {
        log.Printf("Dokumen upload list: %v", existingDokumen.DokumenUpload) // Log seluruh dokumen yang ada

        // Ambil dokumen list dari DokumenUpload
        if dokumenList, ok := existingDokumen.DokumenUpload["dokumen_list"].([]interface{}); ok {
            for _, dok := range dokumenList {
                // Cek apakah dokumen bertipe map
                if d, ok := dok.(map[string]interface{}); ok {
                    if path, exists := d["path"].(string); exists {
                        log.Printf("Attempting to remove old file: %s", path) // Log path yang ingin dihapus

                        // Hapus file lama jika file ada
                        if _, err := os.Stat(path); err == nil {
                            if err := os.Remove(path); err != nil {
                                log.Printf("Failed to remove old file %s: %v", path, err)
                            } else {
                                log.Printf("Successfully removed old file: %s", path)
                            }
                        } else {
                            log.Printf("File does not exist: %s", path)
                        }
                    } else {
                        log.Printf("Path not found in dokumen: %v", d)
                    }
                } else {
                    log.Printf("Invalid dokumen type: %v", dok)
                }
            }
        } else {
            log.Printf("Failed to parse dokumen list")
        }
    } else {
        log.Printf("No dokumen to process in DokumenUpload")
    }

    // Loop untuk menyimpan file yang diunggah
    for _, file := range files {
        fileID := uuid.New().String()
        fileName := file.Filename
        fileExt := filepath.Ext(fileName)
        savePath := fmt.Sprintf("%s/%s%s", uploadDir, fileID, fileExt)

        // Simpan file di direktori uploads
        if err := helper.SaveFile(file, savePath); err != nil {
            log.Printf("Failed to save file %s: %v", fileName, err)
            return nil, fmt.Errorf("failed to save file %s: %w", fileName, err)
        }

        // Buat data dokumen baru
        dokumen := Dokumen{
            ID:       fileID,
            NamaFile: fileName,
            Path:     savePath,
        }

        // Tambahkan dokumen ke list
        dokumenList = append(dokumenList, dokumen)
    }

    // Membungkus array dokumen dalam objek JSON
    dokumenWrapper := map[string]interface{}{
        "dokumen_list": dokumenList,
    }

    // Serialisasi dokumen ke JSON
    dokumenJSON, err := json.Marshal(dokumenWrapper)
    if err != nil {
        log.Printf("Failed to marshal dokumenList: %v", err)
        return nil, fmt.Errorf("failed to marshal dokumenList: %w", err)
    }

    // Konversi []byte ke domain.JSONB
    dokumenJSONB := domain.JSONB{}
    if err := json.Unmarshal(dokumenJSON, &dokumenJSONB); err != nil {
        log.Printf("Failed to unmarshal JSON to domain.JSONB: %v", err)
        return nil, fmt.Errorf("failed to unmarshal JSON to domain.JSONB: %w", err)
    }

    // Update dokumen ke produk
    dokumenEntity := domain.UmkmDokumen{
        UmkmId:        umkmId,
        DokumenUpload: dokumenJSONB, // Simpan dalam format JSONB
    }

    // Simpan ke database
    if _, err := service.dokumenumkmrepository.UpdateDokumen(id, umkmId, dokumenEntity); err != nil {
        log.Printf("Failed to update dokumen in repository: %v", err)
        return nil, fmt.Errorf("failed to update dokumen in repository: %w", err)
    }

    // Siapkan respons
    response := map[string]interface{}{
        "umkm_id":    umkmId,
        "dokumen_id": id,
        "dok_upload": dokumenList, // Dokumen yang berhasil diunggah
    }

    return response, nil
}
