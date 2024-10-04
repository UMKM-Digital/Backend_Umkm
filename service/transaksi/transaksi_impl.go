package transaksiservice

import (
	"errors"
	"fmt"
	"time"

	// "time"
	"umkm/helper"
	"umkm/model/domain"
	"umkm/model/entity"
	"umkm/model/web"

	// querybuilder "umkm/query_builder"
	transaksirepo "umkm/repository/transaksi"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TranssaksiServiceImpl struct {
	transaksirepository transaksirepo.TransaksiRepo
	db                  *gorm.DB
}

func NewTransaksiservice(transaksirepository transaksirepo.TransaksiRepo, db *gorm.DB) *TranssaksiServiceImpl {
	return &TranssaksiServiceImpl{
		transaksirepository: transaksirepository,
		db:                  db,
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

	if transaksi.Status == 0 {
		fmt.Println("Status is 0, consider it valid")
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
		NoHp:             transaksi.NoHp,
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
		"total_jml":          saveTransaksi.TotalJml,
		"keterangan":         saveTransaksi.Keterangan,
		"no_hp":              saveTransaksi.NoHp,
	}, nil
}

// get by id
func (service *TranssaksiServiceImpl) GetKategoriUmkmId(id int) (entity.TransaksiEntity, error) {
	GetTransaksiUmkm, errGetTransaksiUmkm := service.transaksirepository.GetRequestTransaksi(id)

	if errGetTransaksiUmkm != nil {
		return entity.TransaksiEntity{}, errGetTransaksiUmkm
	}

	return entity.ToTransaksiEntity(GetTransaksiUmkm), nil
}

func (service *TranssaksiServiceImpl) GetTransaksiFilter(umkmID uuid.UUID, filtersTanggal map[string]string, allowedfiltersTanggal []string, filters string, limit int, page int, status string) ([]entity.TransasksiFilterEntity, int, int, int, *int, *int, error) {
	var filterTanggal string
	if tanggal, ok := filtersTanggal["tanggal"]; ok {
		filterTanggal = tanggal
	}

	transaksiDomainList, totalCount, currentPage, totalPages, nextPage, prevPage, err := service.transaksirepository.GetFilterTransaksi(umkmID, filters, filterTanggal, limit, page, status)
	if err != nil {
		return nil, 0, 0, 0, nil, nil, err
	}

	// Konversi hasil transaksi ke entitas
	transaksiEntities := entity.ToTransaksiFilterEntities(transaksiDomainList)

	return transaksiEntities, totalCount, currentPage, totalPages, nextPage, prevPage, nil
}




// transaksi web
func (service *TranssaksiServiceImpl) GetTransaksiByYear(id int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    // Panggil repository untuk mengambil data transaksi dengan pagination
    transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := service.transaksirepository.GetFilterTransaksiWebTahunByUserID(id, page, limit, filter)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, fmt.Errorf("failed to retrieve transactions: %w", err)
    }

    // Proses hasil query
    var results []map[string]interface{}
    for _, transaction := range transactions {
        // Buat struktur hasil
        result := map[string]interface{}{
		    "umkm_id":               transaction["umkm_id"], // tambahkan umkm_id
            "year":                  transaction["year"],
            "jumlah_transaksi":      transaction["jumlah_transaksi"],
            "jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
            "jml_transaksi_batal":   transaction["jml_transaksi_batal"],
            "total_berlaku":         transaction["total_berlaku"],
            "total_batal":           transaction["total_batal"],
        }
        results = append(results, result)
    }

    // Kembalikan hasil dengan detail pagination
    return results, totalRecords, currentPage, totalPages, nextPage, prevPage, nil
}


func (service *TranssaksiServiceImpl) GetTransaksiByMonth(umkmID uuid.UUID, user_id int,  year int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    // Panggil repository untuk mengambil data transaksi dengan pagination
    transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := service.transaksirepository.GetTransaksiByMonth(user_id, umkmID,  year, page, limit, filter)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, fmt.Errorf("failed to retrieve transactions: %w", err)
    }

    // Proses hasil query
    var results []map[string]interface{}
    for _, transaction := range transactions {
        // Buat struktur hasil
        result := map[string]interface{}{
			"umkm_id":      		 transaction["umkm_id"],
			"year":					  transaction["year"],
            "month":                 transaction["month"],
            "jumlah_transaksi":      transaction["jumlah_transaksi"],
            "jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
            "jml_transaksi_batal":   transaction["jml_transaksi_batal"],
            "total_berlaku":         transaction["total_berlaku"],
            "total_batal":           transaction["total_batal"],
        }
        results = append(results, result)
    }

    // Kembalikan hasil dengan detail pagination
    return results, totalRecords, currentPage, totalPages, nextPage, prevPage, nil
}




func (service *TranssaksiServiceImpl) GetTransaksiByDate(umkmID uuid.UUID, user_id int, year int, month int, page int, limit int, filter string) ([]map[string]interface{}, int, int, int, *int, *int, error) {
    // Call the repository to get the data along with pagination details
    transactions, totalRecords, currentPage, totalPages, nextPage, prevPage, err := service.transaksirepository.GetTransaksiByDate(user_id, umkmID, year, month, page, limit, filter)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, fmt.Errorf("failed to retrieve transactions: %w", err)
    }

    // Check if transactions data is in the expected format
    var results []map[string]interface{}
    for _, transaction := range transactions {
        // Check and convert the date field from `time.Time`
        parsedDate, ok := transaction["date"].(time.Time)
        if !ok {
            return nil, 0, 0, 0, nil, nil, fmt.Errorf("failed to convert date to time.Time")
        }

        // Format the date as "DD-MM-YYYY"
        formattedDate := parsedDate.Format("02-01-2006")

        // Create the result structure
        result := map[string]interface{}{
			"umkm_id":              transaction["umkm_id"],
			"year":              transaction["year"],
			"month":              transaction["month"],
            "date":                 formattedDate,
            "jumlah_transaksi":     transaction["jumlah_transaksi"],
            "jml_transaksi_berlaku": transaction["jml_transaksi_berlaku"],
            "jml_transaksi_batal":  transaction["jml_transaksi_batal"],
            "total_berlaku":        transaction["total_berlaku"],
            "total_batal":          transaction["total_batal"],
        }

        results = append(results, result)
    }

    // Return data and pagination details in a similar format as `GetProdukList`
    return results, totalRecords, currentPage, totalPages, nextPage, prevPage, nil
}

func (repo *TranssaksiServiceImpl) DeleteByUmkmID(umkmID uuid.UUID) error {
	return repo.db.Where("umkm_id = ?", umkmID).Delete(&domain.Produk{}).Error
}
