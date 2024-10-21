package datarepo

import (
	"context"
	"fmt"
	"log"
	"time"
	"umkm/model/domain"

	"gorm.io/gorm"
)

type DatarepositoryImpl struct {
	db *gorm.DB
}

func NewDataRepositoryImpl(db *gorm.DB) *DatarepositoryImpl{
	return &DatarepositoryImpl{db: db}
}

func (repo *DatarepositoryImpl) CountProductByCategoryWithPercentage() (int64, error) {
    var totalProducts int64

    // Menghitung total produk
    err := repo.db.Model(&domain.Produk{}).Count(&totalProducts).Error
    if err != nil {
        return 0, err
    }

    return totalProducts, nil
}

//menunggu verifikasi
func (repo *DatarepositoryImpl) CountWaitingVerify() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", 0).Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}

//menghitung umkm yang terbina
func (repo *DatarepositoryImpl) CountUmkmBina() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", 1).Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}


//menghitung umkm yang ditolak
func (repo *DatarepositoryImpl) CountUmkmTertolak() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", 3).Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}


func (repo *DatarepositoryImpl) TotalUmkm() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}

//totalumkmyang mikro
func (repo *DatarepositoryImpl) TotalMikro() (int64, error) {
	var totalMikro int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.kriteria_usaha = ?", "Mikro").
		Count(&totalMikro).Error
	
	if err != nil {
		return 0, err
	}

	return totalMikro, nil
}

func (repo *DatarepositoryImpl) TotalMenengah() (int64, error) {
	var TotalMenengah int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.kriteria_usaha = ?", "Menengah").
		Count(&TotalMenengah).Error
	
	if err != nil {
		return 0, err
	}

	return TotalMenengah, nil
}

func (repo *DatarepositoryImpl) TotalKecil() (int64, error) {
	var TotalKecil int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.kriteria_usaha = ?", "Kecil").
		Count(&TotalKecil).Error
	
	if err != nil {
		return 0, err
	}

	return TotalKecil, nil
}
//sektor jaasa
func (repo *DatarepositoryImpl) TotalSektorJasa() (int64, error) {
	var TotalSektorJasa int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.sektor_usaha = ?", "Jasa").
		Count(&TotalSektorJasa).Error
	
	if err != nil {
		return 0, err
	}

	return TotalSektorJasa, nil
}

//sektorproduksi
func (repo *DatarepositoryImpl) TotalSektorProduksi() (int64, error) {
	var TotalSektorProduksi int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.sektor_usaha = ?", "Produksi").
		Count(&TotalSektorProduksi).Error
	
	if err != nil {
		return 0, err
	}

	return TotalSektorProduksi, nil
}

//perdagangan
func (repo *DatarepositoryImpl) TotalSektorPerdagangan() (int64, error) {
	var TotalSektorPerdagangan int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM bertipe Mikro
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.sektor_usaha = ?", "Perdagangan").
		Count(&TotalSektorPerdagangan).Error
	
	if err != nil {
		return 0, err
	}

	return TotalSektorPerdagangan, nil
}

//ekonomikreatif
func (repo *DatarepositoryImpl) TotalEkonomiKreatif() (int64, error) {
	var totalEkonomiKreatif int64

	// Menggabungkan tabel hak_akses dan umkm untuk menghitung UMKM yang ekonomi_kreatif = true
	err := repo.db.Table("umkm").
		Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
		Where("hak_akses.status = ?", 1).
		Where("umkm.ekonomi_kreatif = ?", true). // Mengambil hanya yang bernilai true
		Count(&totalEkonomiKreatif).Error
	
	if err != nil {
		return 0, err
	}

	return totalEkonomiKreatif, nil
}


type KategoriCount struct {
    IdSektorUsaha int    `json:"id_sektor_usaha"`
    Name          string `json:"name"`
    JumlahUmkm    int    `json:"jumlah_umkm"`
}

func (repo *DatarepositoryImpl) GrafikKategoriBySektorUsaha(ctx context.Context, sektorUsahaID int, kecamatan, kelurahan string, tahun int) ([]KategoriCount, error) {
    if sektorUsahaID <= 0 {
        return nil, gorm.ErrInvalidData
    }

    var results []KategoriCount
    // Query untuk mengambil semua kategori UMKM dengan filter tahun
    query := repo.db.WithContext(ctx).Table("kategori_umkm").
        Select("kategori_umkm.id_sektor_usaha, kategori_umkm.name, COALESCE(SUM(CASE WHEN umkm.kode_kec = ? AND umkm.kode_kelurahan = ? AND EXTRACT(YEAR FROM umkm.created_at) = ? THEN 1 ELSE 0 END), 0) AS jumlah_umkm", kecamatan, kelurahan, tahun).
        Where("kategori_umkm.id_sektor_usaha = ?", sektorUsahaID).
        Group("kategori_umkm.id_sektor_usaha, kategori_umkm.name").
        Joins("LEFT JOIN umkm ON kategori_umkm.name = (umkm.kategori_umkm_id->'nama'->>0)")

    err := query.Scan(&results).Error
    if err != nil {
        log.Println("Error executing query:", err)
        return nil, err
    }

    // Cek hasil
    log.Printf("Results from the database: %+v", results)
    return results, nil
}

func (repo *DatarepositoryImpl) TotalUmkmKriteriaUsahaPerBulan(tahun int) (map[string]map[string]int64, error) {
	// Membuat map untuk menyimpan total UMKM per kriteria usaha dan per bulan
	result := make(map[string]map[string]int64)
	kriteriaUsaha := []string{"Mikro", "Kecil", "Menengah"}

	// Inisialisasi map untuk setiap kriteria usaha dan bulan
	for _, kriteria := range kriteriaUsaha {
		result[kriteria] = make(map[string]int64)
		for bulan := 1; bulan <= 12; bulan++ {
			result[kriteria][fmt.Sprintf("%02d", bulan)] = 0
		}
	}

	// Iterasi setiap kriteria usaha untuk menghitung jumlah UMKM per bulan
	for _, kriteria := range kriteriaUsaha {
		var counts []struct {
			Bulan int64 // Pastikan tipe Bulan adalah int64 atau int
			Total int64
		}

		// Menggabungkan tabel hak_akses dan umkm untuk menghitung jumlah UMKM per kriteria usaha
		err := repo.db.Table("umkm").
			Select("EXTRACT(MONTH FROM umkm.created_at) AS bulan, COUNT(*) AS total").
			Joins("JOIN hak_akses ON hak_akses.umkm_id = umkm.id").
			Where("hak_akses.status = ?", 1).
			Where("umkm.kriteria_usaha = ?", kriteria).
			Where("EXTRACT(YEAR FROM umkm.created_at) = ?", tahun).
			Group("bulan").
			Order("bulan").
			Scan(&counts).Error
		
		if err != nil {
			return nil, err
		}

		// Memasukkan hasil perhitungan ke dalam map
		for _, count := range counts {
			// Format bulan menjadi string dengan 2 digit
			result[kriteria][fmt.Sprintf("%02d", count.Bulan)] = count.Total
		}
	}

	return result, nil
}


func (repo *DatarepositoryImpl) TotalUmkmBulan() (int64, error) {
    var TotalUmkmBulan int64

    // Ambil tanggal pertama dan terakhir dari bulan ini
    firstDayOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
    firstDayOfNextMonth := firstDayOfMonth.AddDate(0, 1, 0)

    // Hitung total UMKM bulan ini
    err := repo.db.Model(&domain.HakAkses{}).
        Where("status = ?", 1).
        Where("created_at BETWEEN ? AND ?", firstDayOfMonth, firstDayOfNextMonth).
        Count(&TotalUmkmBulan).Error

    if err != nil {
        return 0, err
    }

    return TotalUmkmBulan, nil
}


func (repo *DatarepositoryImpl) TotalUmkmBulanLalu() (int64, error) {
    var TotalUmkmBulanLalu int64

    // Ambil tanggal sekarang
    currentTime := time.Now()

    // Hitung total UMKM bulan lalu
    err := repo.db.Model(&domain.HakAkses{}).
        Where("status = ?", 1).
        Where("DATE_TRUNC('month', created_at) = DATE_TRUNC('month', ?::timestamp - INTERVAL '1 month')", currentTime).
        Count(&TotalUmkmBulanLalu).Error

    if err != nil {
        return 0, err
    }

    return TotalUmkmBulanLalu, nil
}


func (repo *DatarepositoryImpl) TotalUmkmTahun() (int64, error) {
    var totalUmkmTahun int64

    // Ambil tanggal pertama dan terakhir dari tahun ini
    firstDayOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
    lastDayOfYear := firstDayOfYear.AddDate(1, 0, 0) // Tambahkan 1 tahun untuk mendapatkan awal tahun berikutnya

    // Hitung total UMKM tahun ini
    err := repo.db.Model(&domain.HakAkses{}).
        Where("status = ?", 1).
        Where("created_at BETWEEN ? AND ?", firstDayOfYear, lastDayOfYear).
        Count(&totalUmkmTahun).Error

    if err != nil {
        return 0, err
    }

    return totalUmkmTahun, nil
}

func (repo *DatarepositoryImpl) TotalUmkmTahunLalu() (int64, error) {
    var TotalUmkmBulanLalu int64

    // Ambil tanggal sekarang
    currentTime := time.Now()

    // Hitung total UMKM bulan lalu
    err := repo.db.Model(&domain.HakAkses{}).
        Where("status = ?", 1).
        Where("DATE_TRUNC('year', created_at) = DATE_TRUNC('year', ?::timestamp - INTERVAL '1 year')", currentTime).
        Count(&TotalUmkmBulanLalu).Error

    if err != nil {
        return 0, err
    }

    return TotalUmkmBulanLalu, nil
}

func (repo *DatarepositoryImpl) PersentasiKenaikanUmkm() (float64, error) {
    // Dapatkan total UMKM bulan ini
    totalBulanIni, err := repo.TotalUmkmBulan()
    if err != nil {
        return 0, err
    }

    // Dapatkan total UMKM bulan lalu
    totalBulanLalu, err := repo.TotalUmkmBulanLalu()
    if err != nil {
        return 0, err
    }

    // Jika total bulan lalu adalah 0, maka kita tidak bisa membagi dengan 0
    if totalBulanLalu == 0 {
        if totalBulanIni > 0 {
            return 100, nil // Jika bulan lalu 0 dan bulan ini ada kenaikan, maka kenaikan 100%
        }
        return 0, nil // Jika bulan lalu 0 dan bulan ini juga 0, maka tidak ada kenaikan
    }

    // Hitung persentasi kenaikan
    persentasiKenaikan := (float64(totalBulanIni) - float64(totalBulanLalu)) / float64(totalBulanLalu) * 100

    return persentasiKenaikan, nil
}
func (repo *DatarepositoryImpl) PersentasiKenaikanUmkmTahun() (float64, error) {
    // Dapatkan total UMKM bulan ini
    totalTahunIni, err := repo.TotalUmkmTahun()
    if err != nil {
        return 0, err
    }

    // Dapatkan total UMKM bulan lalu
    totalTahunLalu, err := repo.TotalUmkmTahunLalu()
    if err != nil {
        return 0, err
    }

    // Jika total bulan lalu adalah 0, maka kita tidak bisa membagi dengan 0
    if totalTahunLalu == 0 {
        if totalTahunIni > 0 {
            return 100, nil // Jika bulan lalu 0 dan bulan ini ada kenaikan, maka kenaikan 100%
        }
        return 0, nil // Jika bulan lalu 0 dan bulan ini juga 0, maka tidak ada kenaikan
    }

    // Hitung persentasi kenaikan
    persentasiKenaikan := (float64(totalTahunIni) - float64(totalTahunLalu)) / float64(totalTahunLalu) * 100

    return persentasiKenaikan, nil
}
