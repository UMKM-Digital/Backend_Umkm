package datarepo

import (
	"context"
	"log"
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

