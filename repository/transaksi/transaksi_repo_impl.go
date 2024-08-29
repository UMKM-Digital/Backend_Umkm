package transaksirepo

import (
	"errors"
	"strconv"
	"umkm/model/domain"
	general_query_builder "umkm/query_builder/transaksi"

	"github.com/google/uuid"
	"gorm.io/gorm"
    "fmt"
)

type TransaksirepositoryImpl struct {
	db *gorm.DB
	eventQueryBuilder  general_query_builder.EventQueryBuilder
}

func NewTransaksiRepositoryImpl(db *gorm.DB, eventQueryBuilder  general_query_builder.EventQueryBuilder) *TransaksirepositoryImpl {
	return &TransaksirepositoryImpl{
		db: db,
		eventQueryBuilder: eventQueryBuilder,
	}
}

// regoster
func (repo *TransaksirepositoryImpl) CreateRequetsTransaksi(transaksi domain.Transaksi)(domain.Transaksi, error) {
	if err := repo.db.Create(&transaksi).Error; err != nil {
		return domain.Transaksi{}, err
	}
	return transaksi, nil
}

//get kategori by id
func(repo *TransaksirepositoryImpl) GetRequestTransaksi(idTransaksi int)(domain.Transaksi, error){
	var TrasnsaksiUmkm domain.Transaksi
	err := repo.db.Find(&TrasnsaksiUmkm, "id = ?", idTransaksi).Error
	if err != nil{
		return domain.Transaksi{}, errors.New("Transaksi tidak ditemukan")
	}
	return TrasnsaksiUmkm, nil
}

func (repo *TransaksirepositoryImpl) GetFilterTransaksi(umkmID uuid.UUID) ([]domain.Transaksi, error) {
	var transaksi []domain.Transaksi
	err := repo.db.Where("umkm_id = ?", umkmID).Find(&transaksi).Error
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}


func (repo *TransaksirepositoryImpl) GetFilterTransaksiWebTahun(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error) {
    var results []map[string]interface{}
    offset := (page - 1) * limit

    // Buat query dengan filter
    query := repo.db.Model(&domain.Transaksi{}).
        Select(`
            EXTRACT(YEAR FROM tanggal) AS year,
            COUNT(*) AS jumlah_transaksi,
            SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
            SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
            SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
            SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
        `).
        Where("umkm_id = ?", umkmID).
        Group("EXTRACT(YEAR FROM tanggal)").
        Order("EXTRACT(YEAR FROM tanggal) ASC").
        Limit(limit).
        Offset(offset)

    // Terapkan filter jika ada
    if filter != "" {
        query = query.Where("EXTRACT(YEAR FROM tanggal) = ?", filter)
    }

    // Eksekusi query
    if err := query.Scan(&results).Error; err != nil {
        return nil, err
    }

    return results, nil
}


func (repo *TransaksirepositoryImpl) GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, error) {
    var results []map[string]interface{}
    // Hitung offset berdasarkan halaman
    offset := (page - 1) * limit

    // Membuat query dengan model Transaksi
    query := repo.db.Model(&domain.Transaksi{}).
        Select(`
            EXTRACT(MONTH FROM tanggal) AS month,
            COUNT(*) AS jumlah_transaksi,
            SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) AS jml_transaksi_berlaku,
            SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) AS jml_transaksi_batal,
            SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) AS total_berlaku,
            SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) AS total_batal
        `).
        Where("umkm_id = ?", umkmID).
        Where("EXTRACT(YEAR FROM tanggal) = ?", year).
        Group("month").
        Order("month").
        Limit(limit).
        Offset(offset)

    // Menambahkan filter bulan jika diberikan
    if filter != "" {
        // Mengonversi filter ke integer
        month, err := strconv.Atoi(filter)
        if err != nil {
            return nil, fmt.Errorf("invalid month filter: %v", err)
        }
        query = query.Where("EXTRACT(MONTH FROM tanggal) = ?", month)
    }

    // Jalankan query dan kembalikan hasilnya
    if err := query.Scan(&results).Error; err != nil {
        return nil, err
    }

    return results, nil
}

 
// GetTransaksiByDate retrieves transactions by date, using EventQueryBuilder for pagination and filtering
func (repo *TransaksirepositoryImpl) GetTransaksiByDate(umkmID string, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, error) {
    var results []map[string]interface{}
    // Hitung offset berdasarkan halaman
    offset := (page - 1) * limit


    // Membangun query SQL dengan filter tahun dan bulan
    query := repo.db.Model(&domain.Transaksi{}).
        Select(`
            DATE(tanggal) as date,
            COUNT(*) as jumlah_transaksi,
            SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) as jml_transaksi_berlaku,
            SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) as jml_transaksi_batal,
            SUM(CASE WHEN status = 1 THEN total_jml ELSE 0 END) as total_berlaku,
            SUM(CASE WHEN status = 0 THEN total_jml ELSE 0 END) as total_batal
        `).
        Where("umkm_id = ?", umkmID).
        Where("EXTRACT(YEAR FROM tanggal) = ?", year).
        Where("EXTRACT(MONTH FROM tanggal) = ?", month).
        Group("date").
        Order("date").
        Limit(limit).
        Offset(offset)

        if filter != "" {
            // Mengonversi filter ke integer
            date, err := strconv.Atoi(filter)
            if err != nil {
                return nil, fmt.Errorf("invalid date filter: %v", err)
            }
            query = query.Where("EXTRACT(DAY FROM tanggal) = ?", date)
        }

    // Menjalankan query
    if err := query.Scan(&results).Error; err != nil {
        return nil, err
    }

    return results, nil
}
