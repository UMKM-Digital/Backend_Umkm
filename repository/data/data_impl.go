package datarepo

import (
	"context"
	"database/sql"
	"fmt"
	"math"
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

func (repo *DatarepositoryImpl) GrafikKategoriBySektorUsaha(ctx context.Context, sektorUsahaID int, kodeKecamatan, kodeKelurahan string, tahun int) ([]KategoriCount, error) {
    var results []KategoriCount

    query := repo.db.WithContext(ctx).Table("kategori_umkm").
        Select("kategori_umkm.id_sektor_usaha, kategori_umkm.name, COALESCE(COUNT(DISTINCT CASE WHEN (master_kec.kode_kec = ? OR ? = 'all') AND (master_kel.kode_kel = ? OR ? = '') AND EXTRACT(YEAR FROM umkm.created_at) = ? AND hak_akses.status = 'disetujui' THEN umkm.id END), 0) AS jumlah_umkm", 
            kodeKecamatan, kodeKecamatan, kodeKelurahan, kodeKelurahan, tahun).
        Joins("LEFT JOIN umkm ON kategori_umkm.name = (umkm.kategori_umkm_id->'nama'->>0)").
        Joins("LEFT JOIN hak_akses ON umkm.id = hak_akses.umkm_id").
        Joins("LEFT JOIN master.kecamatan AS master_kec ON umkm.kode_kec = master_kec.nama").
        Joins("LEFT JOIN master.kelurahan AS master_kel ON umkm.kode_kelurahan = master_kel.nama")

    // If sektorUsahaID is 0, fetch all categories from "kategori_umkm"
    if sektorUsahaID != 0 && sektorUsahaID != 4 {
        query = query.Where("kategori_umkm.id_sektor_usaha = ?", sektorUsahaID)
    }

    // Continue the rest of the query
    query = query.Group("kategori_umkm.id_sektor_usaha, kategori_umkm.name").
        Order("kategori_umkm.name ASC").
        Find(&results)

    if query.Error != nil {
        return nil, query.Error
    }

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

firstDayOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
firstDayOfNextMonth := firstDayOfMonth.AddDate(0, 1, 0)

// Format waktu ke string yang sesuai dengan PostgreSQL
firstDayOfMonthStr := firstDayOfMonth.Format("2006-01")
firstDayOfNextMonthStr := firstDayOfNextMonth.Format("2006-01")

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

    // Hitung tahun dan bulan lalu
    year := currentTime.Year()
    month := currentTime.Month()

    // Jika bulan adalah Januari, maka bulan lalu adalah Desember tahun lalu
    if month == 1 {
        year -= 1
        month = 12
    } else {
        month -= 1
    }

    // Format tahun dan bulan ke string
    lastMonth := fmt.Sprintf("%d-%02d", year, month) // Format menjadi 'YYYY-MM'


    // Hitung total UMKM bulan lalu
    err := repo.db.Table("omzets o").
        Select("SUM(o.nominal)").
        Joins("JOIN hak_akses h ON h.umkm_id = o.umkm_id").
        Where("h.status = ?", "disetujui").
        Where("o.bulan = ?", lastMonth).
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
    firstDayOfLastYearStr := firstDayOfLastYear.Format("2006-01")
    lastDayOfLastYearStr := lastDayOfLastYear.Format("2006-01")

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

    persentasiKenaikan = math.Round(persentasiKenaikan*100) / 100

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

    persentasiKenaikan = math.Round(persentasiKenaikan*100) / 100


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


func (repo *DatarepositoryImpl) TotalOmzetPenggunaPerBulan(id int, tahun int) (map[string]map[string]int64, error) {
    var umkmIDs []string
    hasilPerUMKM := make(map[string]map[string]int64)

    // Step 1: Get all umkm_id from approved HakAkses for a specific user_id
    err := repo.db.Model(&domain.HakAkses{}).
        Where("user_id = ? AND status = ?", id, "disetujui").
        Pluck("umkm_id", &umkmIDs).Error

    if err != nil {
        return nil, err
    }

    // Initialize hasilPerUMKM with 0 for each UMKM and month in YYYY-MM format
    for _, umkmID := range umkmIDs {
        hasilPerUMKM[umkmID] = make(map[string]int64)
        for month := 1; month <= 12; month++ {
            bulan := fmt.Sprintf("%d-%02d", tahun, month)
            hasilPerUMKM[umkmID][bulan] = 0
        }
    }

    if len(umkmIDs) == 0 {
        return hasilPerUMKM, nil
    }

    // Step 2: Calculate total omzet per UMKM and per month for a specific year
    type OmzetPerBulan struct {
        UmkmID string
        Bulan  string
        Total  int64
    }

    var omzetPerBulan []OmzetPerBulan

    // Query for total omzet per UMKM and month for a given year
    err = repo.db.Model(&domain.Omset{}).
        Select("umkm_id, TO_CHAR(TO_DATE(bulan, 'YYYY-MM-DD'), 'YYYY-MM') AS bulan, SUM(nominal)::BIGINT AS total").
        Where("umkm_id IN (?) AND EXTRACT(YEAR FROM TO_DATE(bulan, 'YYYY-MM-DD')) = ?", umkmIDs, tahun).
        Group("umkm_id, bulan").
        Order("umkm_id, bulan").
        Scan(&omzetPerBulan).Error

    if err != nil {
        return nil, err
    }

    for _, omzet := range omzetPerBulan {
        hasilPerUMKM[omzet.UmkmID][omzet.Bulan] = omzet.Total
    }

    // Step 4: Map each UMKM ID to its name
    var umkmNames []struct {
        ID   string
        Nama string
    }
    err = repo.db.Model(&domain.UMKM{}).
        Where("id IN (?)", umkmIDs).
        Select("id, name AS nama").
        Scan(&umkmNames).Error
    if err != nil {
        return nil, err
    }

    // Create hasilPerNama with unique keys for UMKM name and ID
    hasilPerNama := make(map[string]map[string]int64)
    for _, umkm := range umkmNames {
        key := fmt.Sprintf("%s", umkm.Nama) // Make key unique by including ID
        if data, exists := hasilPerUMKM[umkm.ID]; exists {
            hasilPerNama[key] = data
        }
    }

    return hasilPerNama, nil
}
