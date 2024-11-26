package storagegambarcontroller

import (
	"encoding/json"
	"net/http"
	storagegambarservice "umkm/service/storagegambarproduk"
)

type StorageGambarController struct {
	Service storagegambarservice.StorageGambarProduk
}

func NewStorageGambarController(service storagegambarservice.StorageGambarProduk) *StorageGambarController {
	return &StorageGambarController{Service: service}
}

func (c *StorageGambarController) UploadFiles(w http.ResponseWriter, r *http.Request) {
	// Panggil service untuk menyimpan banyak file
	filePaths, err := c.Service.SaveImages(r)
	if err != nil {
		http.Error(w, "Failed to upload files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Buat respons JSON dengan daftar file yang berhasil diunggah
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message": "Files uploaded successfully",
		"paths":   filePaths,
	}
	json.NewEncoder(w).Encode(response)
}
