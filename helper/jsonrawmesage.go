 package helper

// import (
// 	"encoding/json"
// 	"strings"
// )


// func ConvertIDKategoriProduk(jsonData string) (string, error) {
// 	var resp Response
// 	err := json.Unmarshal([]byte(jsonData), &resp)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Process and convert id_kategori_produk to kategori
// 	ids := strings.Split(resp.Data.IDKategoriProduk.ID[0], ",")
// 	names := strings.Split(resp.Data.IDKategoriProduk.Nama[0], ",")

// 	var kategori []struct {
// 		ID   string `json:"id"`
// 		Nama string `json:"nama"`
// 	}

// 	for i := range ids {
// 		if i < len(names) {
// 			kategori = append(kategori, struct {
// 				ID   string `json:"id"`
// 				Nama string `json:"nama"`
// 			}{
// 				ID:   strings.TrimSpace(ids[i]),
// 				Nama: strings.TrimSpace(names[i]),
// 			})
// 		}
// 	}

// 	// Set the processed kategori to the response
// 	resp.Data.Kategori = kategori

// 	// Marshal JSON back to string
// 	updatedJSON, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(updatedJSON), nil
// }
