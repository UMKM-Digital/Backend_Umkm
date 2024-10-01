package query_builder_masterlegal

import (
	querybuilder "umkm/query_builder"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type MasteLegalQueryBuilder interface {
	querybuilder.BaseQueryBuilderList
	GetBuilderMasterLegal(filters string, limit int, page int) (*gorm.DB, error)
	GetBuilderDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) (*gorm.DB, error)
	GetBuilderDokumenUmkmStatusAll(userId int, filters string, limit int, page int) (*gorm.DB, error)
}

type MasteLegalQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilderList
	db *gorm.DB
}

func NewMasteLegalQueryBuilder(db *gorm.DB) *MasteLegalQueryBuilderImpl {
	return &MasteLegalQueryBuilderImpl{
		BaseQueryBuilderList: querybuilder.NewBaseQueryBuilderList(db),
		db:                   db,
	}
}

func (MasteLegalQueryBuilder *MasteLegalQueryBuilderImpl) GetBuilderMasterLegal(filters string, limit int, page int) (*gorm.DB, error) {
	query := MasteLegalQueryBuilder.db

	// Implementasi filter di sini
	if filters != "" {
		searchPattern := "%" + filters + "%"
		query = query.Where("nama ILIKE ?", searchPattern)
	}

	offset := (page - 1) * limit
    query = query.Offset(offset).Limit(limit)

	query, err := MasteLegalQueryBuilder.GetQueryBuilderList(query,  limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}


func (builder *MasteLegalQueryBuilderImpl) GetBuilderDokumenUmkmStatus(umkmId uuid.UUID, filters string, limit int, page int) (*gorm.DB, error) {
	query := builder.db.Table("master_dokumen_legal").
		Select("master_dokumen_legal.id,master_dokumen_legal.nama, CASE WHEN umkm_dokumen_legal.id IS NOT NULL THEN 1 ELSE 0 END AS status, umkm_dokumen_legal.created_at AS tanggal_upload").
		Joins("LEFT JOIN umkm_dokumen_legal ON master_dokumen_legal.id = umkm_dokumen_legal.dokumen_id AND umkm_dokumen_legal.umkm_id = ?", umkmId)

	// Implementasi filter jika ada
	if filters != "" {
		searchPattern := "%" + filters + "%"
		query = query.Where("master_dokumen_legal.nama ILIKE ?", searchPattern)
	}
	
	// Tambahkan pagination (limit dan offset)
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)
	
	// Gunakan GetQueryBuilderList dengan menambahkan `builder` sebagai argumen pertama
	query, err := builder.GetQueryBuilderList( query, limit, page)
	if err != nil {
		return nil, err
	}
	
	return query, nil
}

// Func GetBuilderDokumenUmkmStatusAll untuk mengambil dokumen seluruh UMKM
func (builder *MasteLegalQueryBuilderImpl) GetBuilderDokumenUmkmStatusAll(userId int, filters string, limit int, page int) (*gorm.DB, error) {
	var umkmIds []uuid.UUID

	// Mengambil semua UMKM yang diakses oleh user
if err := builder.db.Table("hak_akses").
Select("umkm_id").
Where("user_id = ?", userId).
Find(&umkmIds).Error; err != nil {
return nil, err
}


	// Jika tidak ada UMKM, kembalikan query yang tidak mengembalikan hasil
	if len(umkmIds) == 0 {
		return builder.db.Table("master_dokumen_legal").Where("1 = 0"), nil
	}

	// Query untuk mengambil dokumen master legal dengan status upload
	// Query untuk mengambil dokumen master legal dengan status upload
// query := builder.db.Table("master_dokumen_legal").
// Select(`
// umkm.id AS umkm_id, 
// master_dokumen_legal.id, 
// master_dokumen_legal.nama, 
// CASE 
// 	WHEN umkm_id IS NOT NULL THEN 1 
// 	ELSE 0 
// END AS status`).
// Joins("JOIN umkm ON umkm.id IN (?)", umkmIds). // Mengambil UMKM yang terkait dengan hak akses
// Joins("LEFT JOIN umkm_dokumen_legal ud ON ud.umkm_id = umkm.id AND ud.dokumen_id = master_dokumen_legal.id").
// Group("umkm.id, master_dokumen_legal.id, master_dokumen_legal.nama").
// Order("umkm.id, master_dokumen_legal.id")

query := builder.db.Table("master_dokumen_legal").
    Select(`
        umkm.id AS umkm_id, 
        master_dokumen_legal.id, 
        master_dokumen_legal.nama, 
        CASE 
            WHEN ud.umkm_id IS NOT NULL THEN 1 
            ELSE 0 
        END AS status, 
		 ud.created_at AS tanggal_upload`).
    Joins("JOIN umkm ON umkm.id IN (?)", umkmIds). // Mengambil UMKM yang terkait dengan hak akses
    Joins("LEFT JOIN umkm_dokumen_legal ud ON ud.umkm_id = umkm.id AND ud.dokumen_id = master_dokumen_legal.id").
    Group("umkm.id, master_dokumen_legal.id, master_dokumen_legal.nama, ud.umkm_id, ud.created_at").
    Order("umkm.id, master_dokumen_legal.id")


log.Println("UMKM IDs: ", umkmIds)




	// Implementasi filter jika ada
	if filters != "" {
		searchPattern := "%" + filters + "%"
		query = query.Where("master_dokumen_legal.nama ILIKE ?", searchPattern)
	}

	// Menambahkan pagination
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	return query, nil
}
