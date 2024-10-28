package datarepo

import (
	"context"
	"database/sql"
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
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", domain.Menunggu).Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}

//menghitung umkm yang terbina
func (repo *DatarepositoryImpl) CountUmkmBina() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", "disetujui").Count(&totalVerify).Error
	
	if err != nil {
		return 0, err
	}

	return totalVerify, nil
}


//menghitung umkm yang ditolak
func (repo *DatarepositoryImpl) CountUmkmTertolak() (int64, error) {
	var totalVerify int64

	// Menambahkan kondisi untuk menghitung hanya yang kolom hak akses bernilai 0
	err := repo.db.Model(&domain.HakAkses{}).Where("status = ?", "ditolak").Count(&totalVerify).Error
	
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
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
		Where("hak_akses.status = ?", "disetujui").
		Where("umkm.ekonomi_kreatif = ?", true). // Mengambil hanya yang bernilai true
		Count(&totalEkonomiKreatif).Error
	
	if err != nil {
		return 0, err
	}

	return totalEkonomiKreatif, nil
}


//belum
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
    // Query untuk mengambil semua kategori UMKM dengan filter tahun dan status "disetujui"
    query := repo.db.WithContext(ctx).Table("kategori_umkm").
        Select("kategori_umkm.id_sektor_usaha, kategori_umkm.name, COALESCE(SUM(CASE WHEN umkm.kode_kec = ? AND umkm.kode_kelurahan = ? AND EXTRACT(YEAR FROM umkm.created_at) = ? AND hak_akses.status = 'disetujui' THEN 1 ELSE 0 END), 0) AS jumlah_umkm", kecamatan, kelurahan, tahun).
        Where("kategori_umkm.id_sektor_usaha = ?", sektorUsahaID).
        Joins("LEFT JOIN umkm ON kategori_umkm.name = (umkm.kategori_umkm_id->'nama'->>0)").
        Joins("LEFT JOIN hak_akses ON umkm.id = hak_akses.umkm_id").  // JOIN dengan tabel hak_akses
        Group("kategori_umkm.id_sektor_usaha, kategori_umkm.name")

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
			Where("hak_akses.status = ?", "disetujui").
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
        Where("status = ?", "disetujui").
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
        Where("status = ?", "disetujui").
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
        Where("status = ?", "disetujui").
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
        Where("status = ?", "disetujui").
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


func (repo *DatarepositoryImpl) TotalOmzetBulanIni() (float64, error) {
    var TotalOmzetBulan float64

    // Ambil tanggal pertama dan terakhir dari bulan ini
   // Ambil tanggal awal dan akhir bulan ini
firstDayOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
firstDayOfNextMonth := firstDayOfMonth.AddDate(0, 1, 0)

// Format waktu ke string yang sesuai dengan PostgreSQL
firstDayOfMonthStr := firstDayOfMonth.Format("2006-01-02")
firstDayOfNextMonthStr := firstDayOfNextMonth.Format("2006-01-02")

    // Hitung total UMKM bulan ini
	err := repo.db.Table("omzets o").
    Select("SUM(o.nominal)").
    Joins("JOIN hak_akses h ON h.umkm_id = o.umkm_id").
    Where("h.status = ?", "disetujui").
    Where("o.bulan BETWEEN ? AND ?", firstDayOfMonthStr, firstDayOfNextMonthStr).
    Scan(&TotalOmzetBulan).Error
    if err != nil {
        return 0, err
    }

    return TotalOmzetBulan, nil
}

func (repo *DatarepositoryImpl) TotalOmzetBulanLalu() (float64, error) {
    var TotalUmkmBulanLalu float64

    // Ambil tanggal sekarang
	currentTime := time.Now()

    // Hitung tanggal awal dan akhir bulan lalu
    firstDayOfLastMonth := time.Date(currentTime.Year(), currentTime.Month()-1, 1, 0, 0, 0, 0, time.UTC)
    firstDayOfNextLastMonth := firstDayOfLastMonth.AddDate(0, 1, 0)

    // Format tanggal bulan lalu ke string
    firstDayOfLastMonthStr := firstDayOfLastMonth.Format("2006-01")
    firstDayOfNextLastMonthStr := firstDayOfNextLastMonth.Format("2006-01")

    // Hitung total UMKM bulan lalu
    err := repo.db.Table("omzets o").
		Select("SUM(o.nominal)").
        Joins("JOIN hak_akses h ON h.umkm_id = o.umkm_id").
    Where("h.status = ?", "disetujui").
	Where("TO_DATE(o.bulan, 'YYYY-MM-DD') BETWEEN TO_DATE(?, 'YYYY-MM') AND TO_DATE(?, 'YYYY-MM')",
	firstDayOfLastMonthStr, firstDayOfNextLastMonthStr).
	Scan(&TotalUmkmBulanLalu).Error

    if err != nil {
        return 0, err
    }

    return TotalUmkmBulanLalu, nil
}


//omzet tahun lalu
func (repo *DatarepositoryImpl) TotalomzestTahunIni() (float64, error) {
    var totalOmzetTahunIni sql.NullFloat64

    // Ambil tanggal pertama dan terakhir dari tahun ini
    firstDayOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
    lastDayOfYear := firstDayOfYear.AddDate(1, 0, 0)

    // Format tanggal ke string
    firstDayOfYearStr := firstDayOfYear.Format("2006-01-02")
    lastDayOfYearStr := lastDayOfYear.Format("2006-01-02")

    // Hitung total omzet tahun ini
    err := repo.db.Table("omzets o").
        Select("SUM(o.nominal)").
        Joins("JOIN hak_akses h ON h.umkm_id = o.umkm_id").
        Where("h.status = ?", "disetujui").
        Where("o.bulan BETWEEN ? AND ?", firstDayOfYearStr, lastDayOfYearStr).
        Scan(&totalOmzetTahunIni).Error

    if err != nil {
        return 0, err
    }

    // Jika total NULL, kembalikan 0
    if !totalOmzetTahunIni.Valid {
        return 0, nil
    }

    return totalOmzetTahunIni.Float64, nil
}


//omzet tahun lalu
func (repo *DatarepositoryImpl) TotalOmzetTahunLalu() (float64, error) {
    var totalOmzetTahunLalu sql.NullFloat64

    // Hitung tanggal awal dan akhir tahun lalu
    firstDayOfLastYear := time.Date(time.Now().Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC)
    lastDayOfLastYear := firstDayOfLastYear.AddDate(1, 0, 0)

    // Format tanggal ke string
    firstDayOfLastYearStr := firstDayOfLastYear.Format("2006-01-02")
    lastDayOfLastYearStr := lastDayOfLastYear.Format("2006-01-02")

    // Hitung total omzet tahun lalu
    err := repo.db.Table("omzets o").
        Select("SUM(o.nominal)").
        Joins("JOIN hak_akses h ON h.umkm_id = o.umkm_id").
        Where("h.status = ?", "disetujui").
        Where("o.bulan BETWEEN ? AND ?", firstDayOfLastYearStr, lastDayOfLastYearStr).
        Scan(&totalOmzetTahunLalu).Error

    if err != nil {
        return 0, err
    }

    // Jika total NULL, kembalikan 0
    if !totalOmzetTahunLalu.Valid {
        return 0, nil
    }

    return totalOmzetTahunLalu.Float64, nil
}

//persentasiomsettahunini
func (repo *DatarepositoryImpl) Persentasiomzetbulan() (float64, error) {
    // Dapatkan total UMKM bulan ini
    totalBulanIni, err := repo.TotalOmzetBulanIni()
    if err != nil {
        return 0, err
    }

    // Dapatkan total UMKM bulan lalu
    totalBulanLalu, err := repo.TotalOmzetBulanLalu()
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

//tahun
func (repo *DatarepositoryImpl) Persentasiomzettahun() (float64, error) {
    // Dapatkan total UMKM bulan ini
    totalTahunIni, err := repo.TotalomzestTahunIni()
    if err != nil {
        return 0, err
    }

    // Dapatkan total UMKM bulan lalu
    totalTahunLalu, err := repo.TotalOmzetTahunLalu()
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

//data pengguna
 func(repo *DatarepositoryImpl) TotalUmkmPengguna(id int)(int64, error){
    var totalVerify int64
    err := repo.db.Model(&domain.HakAkses{}).Where("user_id = ? AND status = ?", id, "disetujui").Count(&totalVerify).Error

    if err != nil{
        return 0, err
    }

    return totalVerify, nil
 }

 //data pengguna produk
 func (repo *DatarepositoryImpl) TotalProdukPengguna(id int) (int64, error) {
    var totalProduk int64
    var umkmIDs []string  // Mengubah tipe menjadi []string untuk menampung UUID

    // Langkah 1: Dapatkan semua umkm_id dari HakAkses yang disetujui untuk user_id tertentu
    err := repo.db.Model(&domain.HakAkses{}).
        Where("user_id = ? AND status = ?", id, "disetujui").
        Pluck("umkm_id", &umkmIDs).Error

    if err != nil {
        return 0, err
    }

    if len(umkmIDs) == 0 {
        return 0, nil // Tidak ada UMKM yang disetujui untuk user ini
    }

    // Langkah 2: Hitung total produk yang dimiliki semua UMKM tersebut
    err = repo.db.Model(&domain.Produk{}).
        Where("umkm_id IN (?)", umkmIDs).
        Count(&totalProduk).Error

    if err != nil {
        return 0, err
    }

    return totalProduk, nil
}

func(repo *DatarepositoryImpl) TotalTransaksi(id int)(int64, error){
    var totalTransaksi int64
    var umkmIDs[] string

    err := repo.db.Model(&domain.HakAkses{}).
    Where("user_id = ? AND status = ?", id, "disetujui").
    Pluck("umkm_id", &umkmIDs).Error

if err != nil {
    return 0, err
}

if len(umkmIDs) == 0 {
    return 0, nil // Tidak ada UMKM yang disetujui untuk user ini
}

// Langkah 2: Hitung total produk yang dimiliki semua UMKM tersebut
err = repo.db.Model(&domain.Transaksi{}).
    Where("umkm_id IN (?)", umkmIDs).
    Count(&totalTransaksi).Error

if err != nil {
    return 0, err
}

return totalTransaksi, nil
}


func (repo *DatarepositoryImpl) TotalOmzetPenggunaPerBulan(id int, tahun int) (map[string]int64, error) {
    var umkmIDs []string
    var hasilPerBulan = make(map[string]int64)

    // Langkah 1: Dapatkan semua umkm_id dari HakAkses yang disetujui untuk user_id tertentu
    err := repo.db.Model(&domain.HakAkses{}).
        Where("user_id = ? AND status = ?", id, "disetujui").
        Pluck("umkm_id", &umkmIDs).Error

    if err != nil {
        return nil, err
    }

    if len(umkmIDs) == 0 {
        return hasilPerBulan, nil // Tidak ada UMKM yang disetujui untuk user ini
    }

    // Langkah 2: Hitung total omzet per bulan pada tahun tertentu
    type OmzetPerBulan struct {
        Bulan int   // Untuk menyimpan bulan
        Total int64 // Total omzet bulan tersebut
    }

    var omzetPerBulan []OmzetPerBulan

    // Query untuk mengambil total omzet per bulan pada tahun yang diberikan
    err = repo.db.Model(&domain.Omset{}).
        Select("EXTRACT(MONTH FROM bulan) AS bulan, SUM(amount) AS total").
        Where("umkm_id IN (?) AND EXTRACT(YEAR FROM bulan) = ?", umkmIDs, tahun).
        Group("bulan").
        Order("bulan").
        Scan(&omzetPerBulan).Error

    if err != nil {
        return nil, err
    }

    // Menyimpan hasil per bulan ke dalam map dengan format nama bulan
    bulanNama := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
    for _, omzet := range omzetPerBulan {
        hasilPerBulan[bulanNama[omzet.Bulan-1]] = omzet.Total
    }

    return hasilPerBulan, nil
}


