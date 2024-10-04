package transaksirepo

import (
	"errors"
	"strconv"
	"umkm/model/domain"
	general_query_builder "umkm/query_builder/transaksi"

	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksirepositoryImpl struct {
	db                    *gorm.DB
	transaksiQueryBuilder general_query_builder.TransaksiQueryBuilder
}

func NewTransaksiRepositoryImpl(db *gorm.DB, transaksiQueryBuilder general_query_builder.TransaksiQueryBuilder) *TransaksirepositoryImpl {
	return &TransaksirepositoryImpl{
		db:                    db,
		transaksiQueryBuilder: transaksiQueryBuilder,
	}
}

// regoster
func (repo *TransaksirepositoryImpl) CreateRequetsTransaksi(transaksi domain.Transaksi) (domain.Transaksi, error) {
	if err := repo.db.Create(&transaksi).Error; err != nil {
		return domain.Transaksi{}, err
	}
	return transaksi, nil
}

// get kategori by id
func (repo *TransaksirepositoryImpl) GetRequestTransaksi(idTransaksi int) (domain.Transaksi, error) {
	var TrasnsaksiUmkm domain.Transaksi
	err := repo.db.Find(&TrasnsaksiUmkm, "id = ?", idTransaksi).Error
	if err != nil {
		return domain.Transaksi{}, errors.New("Transaksi tidak ditemukan")
	}
	return TrasnsaksiUmkm, nil
}

// func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID, filters string, filterTanggal string, limit int, page int, status int) ([]domain.Transaksi, error) {
// 	var transaksi []domain.Transaksi
// 	transaksiQueryBuilder, err := repo.transaksiQueryBuilder.GetBuilder(filters, limit, page, status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Terapkan filter tanggal jika ada
// 	if filterTanggal != "" {
// 		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
// 	}

// 	err = transaksiQueryBuilder.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
// 	if err != nil {
// 		return []domain.Transaksi{}, err
// 	}

// 	return transaksi, nil
// }

func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID, filters string, filterTanggal string, limit int, page int, status string) ([]domain.Transaksi, int, int, int, *int, *int, error) {
	var transaksi []domain.Transaksi
	var totalCount int64

	if limit <= 0 {
		limit = 15
	}

	// Query untuk menghitung total record
	transaksiQueryBuilder, err := repo.transaksiQueryBuilder.GetBuilder(filters, 0, 0, status) // Gunakan limit 0 dan page 0 untuk query count
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	if filterTanggal != "" {
		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
	}

	err = transaksiQueryBuilder.Model(&domain.Transaksi{}).Where("umkm_id = ?", umkmID).Count(&totalCount).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	totalPages := 1
	if limit > 0 {
		totalPages = int((totalCount + int64(limit) - 1) / int64(limit))
	}

	// Jika page > totalPages, return kosong
	if page > totalPages {
		return nil, int(totalCount), page, totalPages, nil, nil, nil
	}

	currentPage := page

	// Query untuk mengambil data transaksi dengan limit dan pagination
	transaksiQueryBuilder, err = repo.transaksiQueryBuilder.GetBuilder(filters, limit, page, status)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	if filterTanggal != "" {
		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
	}

	err = transaksiQueryBuilder.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Tentukan nextPage dan prevPage
	var nextPage *int
	if currentPage < totalPages {
		np := currentPage + 1
		nextPage = &np
	}

	var prevPage *int
	if currentPage > 1 {
		pp := currentPage - 1
		prevPage = &pp
	}

	return transaksi, int(totalCount), currentPage, totalPages, nextPage, prevPage, nil
}

// func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebTahun(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
// 	var results []map[string]interface{}
// 	var totalRecords int64

// 	// Hitung offset berdasarkan halaman
// 	offset := (page - 1) * limit

// 	// Query untuk menghitung total records
// 	totalQuery := repo.db.Model(&domain.Transaksi{}).
// 		Where("umkm_id = ?", umkmID)

// 	// Terapkan filter jika ada
// 	if filter != "" {
// 		totalQuery = totalQuery.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
// 	}

// 	// Query utama dengan pagination
// 	query := repo.db.Model(&domain.Transaksi{}).
// 		Select(`
// 			umkm_id, 
//             EXTRACT(YEAR FROM tanggal) AS year,
//             COUNT(*) AS jumlah_transaksi,
//             SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
//             SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
//             SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
//             SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
//         `).
// 		Where("umkm_id = ?", umkmID).
// 		Group("umkm_id,EXTRACT(YEAR FROM tanggal)").
// 		Order("EXTRACT(YEAR FROM tanggal) ASC").
// 		Limit(limit).
// 		Offset(offset)

// 	// Terapkan filter jika ada
// 	if filter != "" {
// 		query = query.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
// 	}

// 	// Eksekusi query
// 	if err := query.Scan(&results).Error; err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	// Hitung total records
// 	totalRecords = int64(len(results))

// 	// Hitung total halaman
// 	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

// 	// Hitung nilai nextPage dan prevPage
// 	var nextPage, prevPage *int
// 	if page < totalPages {
// 		next := page + 1
// 		nextPage = &next
// 	}
// 	if page > 1 {
// 		prev := page - 1
// 		prevPage = &prev
// 	}

// 	// Kembalikan hasil dengan pagination
// 	return results, int(totalRecords), page, totalPages, nextPage, prevPage, nil
// }

func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebTahunByUserID(userID int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    var results []map[string]interface{}
    var totalRecords int64

    // Hitung offset berdasarkan halaman
    offset := (page - 1) * limit

    // Query untuk menghitung total records
    totalQuery := repo.db.Model(&domain.Transaksi{}).
        Joins("JOIN hak_akses ON hak_akses.umkm_id = transaksi.umkm_id").
        Where("hak_akses.user_id = ?", userID)

    // Terapkan filter jika ada
    if filter != "" {
        totalQuery = totalQuery.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
    }

    // Query utama dengan pagination
	query := repo.db.Model(&domain.Transaksi{}).
    Select(`
        transaksi.umkm_id, 
        EXTRACT(YEAR FROM transaksi.tanggal) AS year,
        COUNT(*) AS jumlah_transaksi,
        SUM(CASE WHEN transaksi.status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
        SUM(CASE WHEN transaksi.status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
        SUM(CASE WHEN transaksi.status = 1 THEN transaksi.total_jml ELSE 0 END) AS total_berlaku,
        SUM(CASE WHEN transaksi.status = 0 THEN transaksi.total_jml ELSE 0 END) AS total_batal
    `).
    Joins("JOIN hak_akses ON hak_akses.umkm_id = transaksi.umkm_id").
    Where("hak_akses.user_id = ?", userID).
    Group("transaksi.umkm_id, EXTRACT(YEAR FROM transaksi.tanggal)").
    Order("EXTRACT(YEAR FROM transaksi.tanggal) ASC").
    Limit(limit).
    Offset(offset)


    // Terapkan filter jika ada
    if filter != "" {
        query = query.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
    }

    // Eksekusi query
    if err := query.Scan(&results).Error; err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total records
    totalRecords = int64(len(results))

    // Hitung total halaman
    totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

    // Hitung nilai nextPage dan prevPage
    var nextPage, prevPage *int
    if page < totalPages {
        next := page + 1
        nextPage = &next
    }
    if page > 1 {
        prev := page - 1
        prevPage = &prev
    }

    return results, int(totalRecords), page, totalPages, nextPage, prevPage, nil
}




// func (repo *TransaksirepositoryImpl) GetTransaksiByMonth(umkmID uuid.UUID, year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
// 	var results []map[string]interface{}
// 	var totalRecords int64

// 	// Hitung offset berdasarkan halaman
// 	offset := (page - 1) * limit

// 	// Query untuk mengambil transaksi dengan limit dan offset
// 	query := repo.db.Model(&domain.Transaksi{}).
// 		Select(`
// 			umkm_id,
// 			EXTRACT(YEAR FROM tanggal) AS year,
// 			EXTRACT(MONTH FROM tanggal) AS month,
// 			COUNT(*) AS jumlah_transaksi,
// 			SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
// 			SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
// 			SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
// 			SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
// 		`).
// 		Where("umkm_id = ? AND EXTRACT(YEAR FROM tanggal) = ?", umkmID, year). // Correct year filtering
// 		Group("umkm_id, year, month").   // Group by year and month
// 		Order("month ASC").
// 		Limit(limit).
// 		Offset(offset)

// 	// Menambahkan filter bulan jika diberikan
// 	if filter != "" {
// 		month, err := strconv.Atoi(filter)
// 		if err != nil {
// 			return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid month filter: %v", err)
// 		}
// 		query = query.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
// 	}

// 	// Jalankan query dan simpan hasilnya
// 	if err := query.Scan(&results).Error; err != nil {
// 		return nil, 0, 0, 0, nil, nil, err
// 	}

// 	// Hitung total records dari hasil data
// 	totalRecords = int64(len(results))

// 	// Hitung total halaman
// 	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

// 	// Hitung nilai nextPage dan prevPage
// 	var nextPage, prevPage *int
// 	if page < totalPages {
// 		next := page + 1
// 		nextPage = &next
// 	}
// 	if page > 1 {
// 		prev := page - 1
// 		prevPage = &prev
// 	}

// 	// Kembalikan hasil dengan pagination
// 	return results, int(totalRecords), page, totalPages, nextPage, prevPage, nil
// }
func (repo *TransaksirepositoryImpl) GetTransaksiByMonth(userId int, umkmID uuid.UUID, year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    var results []map[string]interface{}
    var totalRecords int64

    // Hitung offset berdasarkan halaman
    offset := (page - 1) * limit

    // Query dasar
    query := repo.db.Model(&domain.Transaksi{}).Select(` 
        umkm_id,
        EXTRACT(YEAR FROM tanggal) AS year,
        EXTRACT(MONTH FROM tanggal) AS month,
        COUNT(*) AS jumlah_transaksi,
        SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
        SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
        SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
        SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
    `)

    // Jika umkmID tidak kosong, filter berdasarkan umkmID
    if umkmID != (uuid.UUID{}) {
        query = query.Where("umkm_id = ?", umkmID)
    } else {
        // Jika umkmID kosong, ambil semua UMKM yang dimiliki user
        query = query.Where("umkm_id IN (SELECT umkm_id FROM hak_akses WHERE user_id = ?)", userId)
    }

    // Tambahkan filter tahun jika diberikan
    if year > 0 {
        query = query.Where("EXTRACT(YEAR FROM tanggal) = ?", year)
    }

    // Tambahkan limit dan offset
    query = query.Group("umkm_id, year, month").Order("month ASC").Limit(limit).Offset(offset)

    // Tambahkan filter bulan jika diberikan
    if filter != "" {
        month, err := strconv.Atoi(filter)
        if err != nil {
            return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid month filter: %v", err)
        }
        query = query.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
    }

    // Jalankan query dan simpan hasilnya
    if err := query.Scan(&results).Error; err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total records dari hasil data
    totalRecords = int64(len(results))

    // Hitung total halaman
    totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

    // Hitung nilai nextPage dan prevPage
    var nextPage, prevPage *int
    if page < totalPages {
        next := page + 1
        nextPage = &next
    }
    if page > 1 {
        prev := page - 1
        prevPage = &prev
    }

    // Kembalikan hasil dengan pagination
    return results, int(totalRecords), page, totalPages, nextPage, prevPage, nil
}





// func (repo *TransaksirepositoryImpl) GetTransaksiByDate(umkmID uuid.UUID, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
//     var results []map[string]interface{}
//     var totalRecords int64

//     if limit <= 0 {
//         limit = 15
//     }

//     // Hitung offset berdasarkan halaman
//     offset := (page - 1) * limit

//     // Query untuk menghitung total records
//     countQuery := repo.db.Model(&domain.Transaksi{}).
//         Select("COUNT(DISTINCT DATE(tanggal))").
//         Where("umkm_id = ?", umkmID).
//         Where("EXTRACT(YEAR FROM tanggal) = ?", year).
//         Where("EXTRACT(MONTH FROM tanggal) = ?", month)

//     if filter != "" {
//         date, err := strconv.Atoi(filter)
//         if err != nil {
//             return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid date filter: %v", err)
//         }
//         countQuery = countQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
//     }

//     // Hitung jumlah total records
//     if err := countQuery.Count(&totalRecords).Error; err != nil {
//         return nil, 0, 0, 0, nil, nil, fmt.Errorf("count query failed: %w", err)
//     }

//     totalPages := 1
//     if limit > 0 {
//         totalPages = int((totalRecords + int64(limit) - 1) / int64(limit))
//     }

//     // Jika page > totalPages, return kosong
//     if page > totalPages {
//         return nil, int(totalRecords), page, totalPages, nil, nil, nil
//     }

//     currentPage := page

//     // Query untuk mengambil data dengan limit dan offset
//     dataQuery := repo.db.Model(&domain.Transaksi{}).
//         Select(`
// 		umkm_id,
// 			EXTRACT(YEAR FROM tanggal) AS year,
// 			EXTRACT(MONTH FROM tanggal) AS month,
//             DATE(tanggal) as date,
//             COUNT(*) as jumlah_transaksi,
//             SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) as jml_transaksi_berlaku,
//             SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) as jml_transaksi_batal,
//             SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) as total_berlaku,
//             SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) as total_batal
//         `).
//         Where("umkm_id = ?", umkmID).
//         Where("EXTRACT(YEAR FROM tanggal) = ?", year).
//         Where("EXTRACT(MONTH FROM tanggal) = ?", month).
//         Group("umkm_id, year, month, DATE(tanggal)").
//         Order("DATE(tanggal)").
//         Limit(limit).
//         Offset(offset)

//     if filter != "" {
//         date, err := strconv.Atoi(filter)
//         if err != nil {
//             return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid date filter: %v", err)
//         }
//         dataQuery = dataQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
//     }

//     // Menjalankan query untuk mendapatkan data
//     if err := dataQuery.Scan(&results).Error; err != nil {
//         return nil, 0, 0, 0, nil, nil, fmt.Errorf("data query failed: %w", err)
//     }

//     // Tentukan nextPage dan prevPage
//     var nextPage *int
//     if currentPage < totalPages {
//         np := currentPage + 1
//         nextPage = &np
//     }

//     var prevPage *int
//     if currentPage > 1 {
//         pp := currentPage - 1
//         prevPage = &pp
//     }

//     return results, int(totalRecords), currentPage, totalPages, nextPage, prevPage, nil
// }

func (repo *TransaksirepositoryImpl) GetTransaksiByDate(userId int, umkmID uuid.UUID, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    var results []map[string]interface{}
    var totalRecords int64

    if limit <= 0 {
        limit = 15
    }

    // Hitung offset berdasarkan halaman
    offset := (page - 1) * limit

    // Query untuk menghitung total records
    countQuery := repo.db.Model(&domain.Transaksi{}).
        Select("COUNT(DISTINCT DATE(tanggal))")

    // Filter berdasarkan umkmID atau hak akses jika umkmID tidak disediakan
    if umkmID != (uuid.UUID{}) {
        countQuery = countQuery.Where("umkm_id = ?", umkmID)
    } else {
        countQuery = countQuery.Where("umkm_id IN (SELECT umkm_id FROM hak_akses WHERE user_id = ?)", userId)
    }

    // Jika tahun dan bulan disediakan, tambahkan filter
    if year > 0 {
        countQuery = countQuery.Where("EXTRACT(YEAR FROM tanggal) = ?", year)
    }
    if month > 0 {
        countQuery = countQuery.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
    }

    if filter != "" {
        date, err := strconv.Atoi(filter)
        if err != nil {
            return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid date filter: %v", err)
        }
        countQuery = countQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
    }

    // Hitung jumlah total records
    if err := countQuery.Count(&totalRecords).Error; err != nil {
        return nil, 0, 0, 0, nil, nil, fmt.Errorf("count query failed: %w", err)
    }

    totalPages := 1
    if limit > 0 {
        totalPages = int((totalRecords + int64(limit) - 1) / int64(limit))
    }

    // Jika page > totalPages, return kosong
    if page > totalPages {
        return nil, int(totalRecords), page, totalPages, nil, nil, nil
    }

    currentPage := page

    // Query untuk mengambil data dengan limit dan offset
    dataQuery := repo.db.Model(&domain.Transaksi{}).
        Select(`
            umkm_id,
            EXTRACT(YEAR FROM tanggal) AS year,
            EXTRACT(MONTH FROM tanggal) AS month,
            DATE(tanggal) as date,
            COUNT(*) as jumlah_transaksi,
            SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) as jml_transaksi_berlaku,
            SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) as jml_transaksi_batal,
            SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) as total_berlaku,
            SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) as total_batal
        `)

    // Filter berdasarkan umkmID atau hak akses
    if umkmID != (uuid.UUID{}) {
        dataQuery = dataQuery.Where("umkm_id = ?", umkmID)
    } else {
        dataQuery = dataQuery.Where("umkm_id IN (SELECT umkm_id FROM hak_akses WHERE user_id = ?)", userId)
    }

    // Tambahkan filter tahun dan bulan jika disediakan
    if year > 0 {
        dataQuery = dataQuery.Where("EXTRACT(YEAR FROM tanggal) = ?", year)
    }
    if month > 0 {
        dataQuery = dataQuery.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
    }

    dataQuery = dataQuery.Group("umkm_id, year, month, DATE(tanggal)").
        Order("DATE(tanggal)").
        Limit(limit).
        Offset(offset)

    if filter != "" {
        date, err := strconv.Atoi(filter)
        if err != nil {
            return nil, 0, 0, 0, nil, nil, fmt.Errorf("invalid date filter: %v", err)
        }
        dataQuery = dataQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
    }

    // Menjalankan query untuk mendapatkan data
    if err := dataQuery.Scan(&results).Error; err != nil {
        return nil, 0, 0, 0, nil, nil, fmt.Errorf("data query failed: %w", err)
    }

    // Hitung total records dari hasil data
    totalRecords = int64(len(results))

    // Hitung total halaman
    totalPages = int((totalRecords + int64(limit) - 1) / int64(limit))

    // Tentukan nextPage dan prevPage
    var nextPage *int
    if currentPage < totalPages {
        np := currentPage + 1
        nextPage = &np
    }

    var prevPage *int
    if currentPage > 1 {
        pp := currentPage - 1
        prevPage = &pp
    }

    return results, int(totalRecords), currentPage, totalPages, nextPage, prevPage, nil
}




func(repo *TransaksirepositoryImpl) DeleteTransaksiUmkmId(id uuid.UUID) error{
	return repo.db.Where("umkm_id = ?", id).Delete(&domain.Transaksi{}).Error
}