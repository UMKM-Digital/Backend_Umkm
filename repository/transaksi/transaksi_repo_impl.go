package transaksirepo

// import (
// 	"errors"
// 	"strconv"
// 	"umkm/model/domain"
// 	general_query_builder "umkm/query_builder/transaksi"

// 	"fmt"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type TransaksirepositoryImpl struct {
// 	db                    *gorm.DB
// 	transaksiQueryBuilder general_query_builder.TransaksiQueryBuilder
// }

// func NewTransaksiRepositoryImpl(db *gorm.DB, transaksiQueryBuilder general_query_builder.TransaksiQueryBuilder) *TransaksirepositoryImpl {
// 	return &TransaksirepositoryImpl{
// 		db:                    db,
// 		transaksiQueryBuilder: transaksiQueryBuilder,
// 	}
// }

// // regoster
// func (repo *TransaksirepositoryImpl) CreateRequetsTransaksi(transaksi domain.Transaksi) (domain.Transaksi, error) {
// 	if err := repo.db.Create(&transaksi).Error; err != nil {
// 		return domain.Transaksi{}, err
// 	}
// 	return transaksi, nil
// }

// // get kategori by id
// func (repo *TransaksirepositoryImpl) GetRequestTransaksi(idTransaksi int) (domain.Transaksi, error) {
// 	var TrasnsaksiUmkm domain.Transaksi
// 	err := repo.db.Find(&TrasnsaksiUmkm, "id = ?", idTransaksi).Error
// 	if err != nil {
// 		return domain.Transaksi{}, errors.New("Transaksi tidak ditemukan")
// 	}
// 	return TrasnsaksiUmkm, nil
// }

// // func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID, filters string, filterTanggal string, limit int, page int, status int) ([]domain.Transaksi, error) {
// // 	var transaksi []domain.Transaksi
// // 	transaksiQueryBuilder, err := repo.transaksiQueryBuilder.GetBuilder(filters, limit, page, status)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	// Terapkan filter tanggal jika ada
// // 	if filterTanggal != "" {
// // 		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
// // 	}

// // 	err = transaksiQueryBuilder.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
// // 	if err != nil {
// // 		return []domain.Transaksi{}, err
// // 	}

// // 	return transaksi, nil
// // }
// func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID, filters string, filterTanggal string, limit int, page int, status string) ([]domain.Transaksi, int, error) {
// 	var transaksi []domain.Transaksi
// 	var totalCount int64

// 	// Query untuk menghitung total record
// 	transaksiQueryBuilder, err := repo.transaksiQueryBuilder.GetBuilder(filters, 0, 0, status) // Gunakan limit 0 dan page 0 untuk query count
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	if filterTanggal != "" {
// 		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
// 	}

// 	err = transaksiQueryBuilder.Model(&domain.Transaksi{}).Where("umkm_id = ?", umkmID).Count(&totalCount).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	// Query untuk mengambil data transaksi dengan limit dan pagination
// 	transaksiQueryBuilder, err = repo.transaksiQueryBuilder.GetBuilder(filters, limit, page, status)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	if filterTanggal != "" {
// 		transaksiQueryBuilder = transaksiQueryBuilder.Where("DATE(tanggal) = ?", filterTanggal)
// 	}

// 	err = transaksiQueryBuilder.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return transaksi, int(totalCount), nil
// }



// func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebTahun(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error) {
// 	var results []map[string]interface{}
// 	offset := (page - 1) * limit

// 	query := repo.db.Model(&domain.Transaksi{}).
// 		Select(`
//             EXTRACT(YEAR FROM tanggal) AS year,
//             COUNT(*) AS jumlah_transaksi,
//             SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
//             SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
//             SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
//             SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
//         `).
// 		Where("umkm_id = ?", umkmID).
// 		Group("EXTRACT(YEAR FROM tanggal)").
// 		Order("EXTRACT(YEAR FROM tanggal) ASC").
// 		Limit(limit).
// 		Offset(offset)

// 	// Terapkan filter jika ada
// 	if filter != "" {
// 		query = query.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
// 	}

// 	// Eksekusi query
// 	if err := query.Scan(&results).Error; err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func (repo *TransaksirepositoryImpl) GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, error) {
// 	var results []map[string]interface{}
// 	// Hitung offset berdasarkan halaman
// 	offset := (page - 1) * limit

// 	// Membuat query dengan model Transaksi
// 	query := repo.db.Model(&domain.Transaksi{}).
// 		Select(`
//             EXTRACT(MONTH FROM tanggal) AS month,
//             COUNT(*) AS jumlah_transaksi,
//             SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
//             SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
//             SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
//             SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
//         `).
// 		Where("umkm_id = ?", umkmID).
// 		Where("EXTRACT(YEAR FROM tanggal) = ?", year).
// 		Group("month").
// 		Order("month").
// 		Limit(limit).
// 		Offset(offset)

// 	// Menambahkan filter bulan jika diberikan
// 	if filter != "" {
// 		// Mengonversi filter ke integer
// 		month, err := strconv.Atoi(filter)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid month filter: %v", err)
// 		}
// 		query = query.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
// 	}

// 	// Jalankan query dan kembalikan hasilnya
// 	if err := query.Scan(&results).Error; err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func (repo *TransaksirepositoryImpl) GetTransaksiByDate(umkmID string, year int, month int, page int, limit int, filter string) (map[string]interface{}, error) {
//     var results []map[string]interface{}
//     var totalRecords int64

//     // Hitung offset berdasarkan halaman
//     offset := (page - 1) * limit

//     // Query untuk menghitung total records berdasarkan tanggal unik tanpa limit dan offset
//     countQuery := repo.db.Model(&domain.Transaksi{}).
//         Select("COUNT(DISTINCT DATE(tanggal))").
//         Where("umkm_id = ?", umkmID).
//         Where("EXTRACT(YEAR FROM tanggal) = ?", year).
//         Where("EXTRACT(MONTH FROM tanggal) = ?", month)

//     if filter != "" {
//         date, err := strconv.Atoi(filter)
//         if err != nil {
//             return nil, fmt.Errorf("invalid date filter: %v", err)
//         }
//         countQuery = countQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
//     }

//     // Hitung jumlah total records
//     if err := countQuery.Count(&totalRecords).Error; err != nil {
//         return nil, fmt.Errorf("count query failed: %w", err)
//     }

//     // Query untuk mengambil data dengan limit dan offset
//     dataQuery := repo.db.Model(&domain.Transaksi{}).
//         Select(`
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
//         Group("DATE(tanggal)").
//         Order("DATE(tanggal)").
//         Limit(limit).
//         Offset(offset)

//     if filter != "" {
//         date, err := strconv.Atoi(filter)
//         if err != nil {
//             return nil, fmt.Errorf("invalid date filter: %v", err)
//         }
//         dataQuery = dataQuery.Where("EXTRACT(DAY FROM tanggal) = ?", date)
//     }

//     // Menjalankan query untuk mendapatkan data
//     if err := dataQuery.Scan(&results).Error; err != nil {
//         return nil, fmt.Errorf("data query failed: %w", err)
//     }

//     // Membentuk respons akhir
//     response := map[string]interface{}{
//         "total_records": totalRecords,
//         "transactions":  results,
//     }

//     return response, nil
// }
