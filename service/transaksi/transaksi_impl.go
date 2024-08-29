package transaksiservice

import (
	"errors"
	"fmt"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"
	querybuilder "umkm/query_builder"
	transaksirepo "umkm/repository/transaksi"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TranssaksiServiceImpl struct {
	transaksirepository transaksirepo.TransaksiRepo
	db *gorm.DB
}

func NewTransaksiservice(transaksirepository transaksirepo.TransaksiRepo, db *gorm.DB) *TranssaksiServiceImpl {
	return &TranssaksiServiceImpl{
		transaksirepository: transaksirepository,
		db: db,
	}
}

// register
// service/transaksiservice/auth_service_impl.go
func (service *TranssaksiServiceImpl) CreateTransaksi(transaksi web.CreateTransaksi) (map[string]interface{}, error) {
	date, err := helper.ParseDate(transaksi.Tanggal)
	IDKategoriProduk, err := helper.RawMessageToJSONB(transaksi.IDKategoriProduk)
	if err != nil {
		return nil, errors.New("invalid type for idKategoriProduk")
	}

	invoiceNumber, err := helper.GenerateInvoiceNumber(service.db, transaksi.UmkmId)
    if err != nil {
        return nil, errors.New("failed to generate invoice number")
    }

	qrcode, err := helper.GenerateValidationTicket()
	if err != nil {
		return nil, errors.New("failed to generate QR code")
	}

	if qrcode == "" {
		return nil, errors.New("QR code generation returned empty value")
	}

	// Logging QR code untuk memastikan
	fmt.Printf("Generated QR Code: %s\n", qrcode)
	fmt.Printf("Received Total Jml: %f\n", transaksi.TotalJml) // Log total_jml

	totalJmlDecimal := decimal.NewFromFloat(transaksi.TotalJml)

	newTransaksi := domain.Transaksi{
		UmkmId:           transaksi.UmkmId,
		NoInvoice:        invoiceNumber,
		Tanggal:          date,
		Nameclient:       transaksi.NamaClient,
		IdKategoriProduk: IDKategoriProduk,
		TotalJml:         totalJmlDecimal,
		Keterangan:       transaksi.Keteranagan,
		Status:           transaksi.Status,
		TiketValidasi:    qrcode, // Pastikan ini terisi
		NoHp: transaksi.NoHp,
	}

	saveTransaksi, errSaveTransaksi := service.transaksirepository.CreateRequetsTransaksi(newTransaksi)
	if errSaveTransaksi != nil {
		return nil, errSaveTransaksi
	}

	return map[string]interface{}{
		"id_umkm":            saveTransaksi.UmkmId,
		"no_invoice":         saveTransaksi.NoInvoice,
		"tanggal":            saveTransaksi.Tanggal,
		"name_client":        saveTransaksi.Nameclient,
		"id_kategori_produk": saveTransaksi.IdKategoriProduk,
		"total_jml":              saveTransaksi.TotalJml,
		"keterangan":         saveTransaksi.Keterangan,
		"no_hp":         saveTransaksi.NoHp,
	}, nil
}


//get by id
func (service *TranssaksiServiceImpl) GetKategoriUmkmId(id int)(entity.TransaksiEntity, error) {
	GetTransaksiUmkm, errGetTransaksiUmkm := service.transaksirepository.GetRequestTransaksi(id)

	if errGetTransaksiUmkm != nil {
		return entity.TransaksiEntity{}, errGetTransaksiUmkm
	}

	return entity.ToTransaksiEntity(GetTransaksiUmkm), nil
}

//filter
// func (service *TranssaksiServiceImpl) GetTransaksiFilter(umkmID uuid.UUID) ([]entity.TransasksiFilterEntity, error) {
//     // Misalkan Anda memiliki metode repository untuk mendapatkan transaksi berdasarkan umkmID
//     domainTransaksiList, err := service.transaksirepository.GetFilterTransaksi(umkmID)
//     if err != nil {
//         return nil, err
//     }
    
//     // Konversi dari domain.Transaksi ke entity.TransasksiFilterEntity
//     return entity.ToTransaksiFilterEntities(domainTransaksiList), nil
// }


func (service *TranssaksiServiceImpl) GetTransaksiFilter(umkmID uuid.UUID, filters map[string]string, allowedFilters []string) ([]entity.TransasksiFilterEntity, error) {
    queryBuilder := querybuilder.NewBaseQueryBuilder(service.db)

    // Menggunakan queryBuilder untuk menerapkan filter
    query, err := queryBuilder.GetQueryBuilder(filters, allowedFilters)
    if err != nil {
        return nil, err
    }

    // Menambahkan filter untuk umkmID
    query = query.Where("umkm_id = ?", umkmID)

    // Debug: Tampilkan SQL yang dihasilkan
    sql := service.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
        return tx.Model(&entity.TransasksiFilterEntity{}).Where("umkm_id = ?", umkmID).Where("tanggal = ?", filters["tanggal"])
    })
    fmt.Println("Generated SQL Query:", sql)

    var transaksiList []entity.TransasksiFilterEntity
    result := query.Find(&transaksiList)
    if result.Error != nil {
        return nil, result.Error
    }

    return transaksiList, nil
}

//transaksi web
func (service *TranssaksiServiceImpl) GetTransaksiByYear(umkmID string, page int, limit int, filter string) ([]map[string]interface{}, error) {
    var results []map[string]interface{}

    transactions, err := service.transaksirepository.GetFilterTransaksiWebTahun(umkmID, page, limit, filter)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve transactions: %w", err)
    }

    // Process the results
    for _, transaction := range transactions {
        result := map[string]interface{}{
            "year":               transaction["year"],  // Access value using the key
            "jumlah_transaksi":   transaction["jumlah_transaksi"],
            "jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
            "jml_transaksi_batal": transaction["jml_transaksi_batal"],
            "total_berlaku":      transaction["total_berlaku"], // Assuming these are strings already
            "total_batal":        transaction["total_batal"],
        }
        results = append(results, result)
    }

    return results, nil
}


func (service *TranssaksiServiceImpl) GetTransaksiByMonth(umkmID string, year int, page int, limit int, filter string) ([]map[string]interface{}, error) {
    var results []map[string]interface{}

    transactions, err := service.transaksirepository.GetTransaksiByMonth(umkmID,year,page,limit,filter)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve transactions: %w", err)
    }

    // Parse hasil query
     // Process the results
     for _, transaction := range transactions {
        result := map[string]interface{}{
            "month":               transaction["month"],  // Access value using the key
            "jumlah_transaksi":   transaction["jumlah_transaksi"],
            "jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
            "jml_transaksi_batal": transaction["jml_transaksi_batal"],
            "total_berlaku":      transaction["total_berlaku"], // Assuming these are strings already
            "total_batal":        transaction["total_batal"],
        }
        results = append(results, result)
    }

    return results, nil
}


func (service *TranssaksiServiceImpl) GetTransaksiByDate(umkmID string, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	transactions, err := service.transaksirepository.GetTransaksiByDate(umkmID, year, month, page, limit, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transactions: %w", err)
	}

	for _, transaction := range transactions {
		// Asumsikan bahwa "date" adalah tipe time.Time
		parsedDate, ok := transaction["date"].(time.Time)
		if !ok {
			return nil, fmt.Errorf("failed to convert date to time.Time")
		}

		// Format tanggal menjadi Tanggal Bulan Tahun
		formattedDate := parsedDate.Format("02-01-2006")

		result := map[string]interface{}{
			"date":                 formattedDate,
			"jumlah_transaksi":     transaction["jumlah_transaksi"],
			"jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
			"jml_transaksi_batal":  transaction["jml_transaksi_batal"],
			"total_berlaku":        transaction["total_berlaku"],
			"total_batal":          transaction["total_batal"],
		}
		results = append(results, result)
	}
	return results, nil
}
