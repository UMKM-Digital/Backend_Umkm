package entity

import (
	"time"
	"umkm/model/domain"

	"github.com/google/uuid"
)

type UmkmFilterEntity struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Gambar domain.JSONB `json:"gambar"`
	Lokasi string `json:"lokasi"`
	PenanggunagJawab string `json:"nama_penanggung_jawab"`
	KategoriUmkm domain.JSONB `json:"kategori_umkm_id"`
	 TotalProduk        int           `json:"total_produk"`
	 Progres int 	`json:"progres"`
	 CreatedAt time.Time   `json:"created_at"`
	 Status string `json:"status"`
}

func ToUmkmFilterEntity(umkm domain.UMKM, products []domain.Produk, dokumenList []domain.UmkmDokumen, masterDokumenList []domain.MasterDokumenLegal, hakAkses domain.HakAkses) UmkmFilterEntity {
	totalProduk := CalculateTotalProdukByUmkm(umkm.IdUmkm, products)
	progres := CalculateProgresDokumen(umkm.IdUmkm, dokumenList, masterDokumenList)
	status := "tidak ada status"
    if hakAkses.Status != "" {
        status = string(hakAkses.Status)
    }

	return UmkmFilterEntity{
	Id: umkm.IdUmkm,
	Name: umkm.Name,
	Gambar: umkm.Images,
	Lokasi: umkm.Lokasi,
	PenanggunagJawab: umkm.NamaPenanggungJawab,
	KategoriUmkm: umkm.KategoriUmkmId,
	TotalProduk:        totalProduk, // Menambahkan total produk
	Progres:            progres, 
	CreatedAt: umkm.CreatedAt,
	Status: status,
	}
}

func ToUmkmfilterEntities(umkmList []domain.UMKM, products []domain.Produk, dokumenList []domain.UmkmDokumen, masterDokumenList []domain.MasterDokumenLegal, hakAksesList []domain.HakAkses) []UmkmFilterEntity {
    var umkmEntities []UmkmFilterEntity
    for _, umkm := range umkmList {
        // Cari status hak akses berdasarkan ID UMKM
        var hakAkses domain.HakAkses
        for _, ha := range hakAksesList {
            if ha.UmkmId == umkm.IdUmkm {
                hakAkses = ha
                break
            }
        }
        
        // Konversi UMKM ke UmkmFilterEntity
        umkmEntity := ToUmkmFilterEntity(umkm, products, dokumenList, masterDokumenList, hakAkses)
        umkmEntities = append(umkmEntities, umkmEntity)
    }
    return umkmEntities
}



//menghitung produk
func CalculateTotalProdukByUmkm(umkmId uuid.UUID, products []domain.Produk) int {
    total := 0
    for _, produk := range products {
        if produk.UmkmId == umkmId {
            total++
        }
    }
    return total
}


func CalculateProgresDokumen(umkmId uuid.UUID, dokumenList []domain.UmkmDokumen, masterDokumenList []domain.MasterDokumenLegal) int {

	totalWajib := 0
	totalDiisi := 0

	// Loop untuk hitung dokumen yang wajib
	for _, masterDokumen := range masterDokumenList {
		if masterDokumen.Iswajib == 1 { 
			totalWajib++

		
			for _, dokumen := range dokumenList {
				if dokumen.UmkmId == umkmId && dokumen.DokumenId == masterDokumen.IdMasterDokumenLegal {
					totalDiisi++
					break
				}
			}
		}
	}

	if totalWajib == 0 {
		return 100
	}

	// Hitung progres dalam persen
	return (totalDiisi * 100) / totalWajib
}
