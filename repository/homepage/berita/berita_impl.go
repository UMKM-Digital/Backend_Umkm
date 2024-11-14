package beritarepo

import (
	"context"
	"umkm/model/domain"
	query_builder_berita "umkm/query_builder/berita"
	"errors"

	"gorm.io/gorm"
)

type BeritaRepoImpl struct {
	db *gorm.DB
	beritarepo query_builder_berita.BeritaQueryBuilder
}

func NewBerita(db *gorm.DB, beritarepo query_builder_berita.BeritaQueryBuilder) *BeritaRepoImpl {
	return &BeritaRepoImpl{
		db: db,
		beritarepo: beritarepo,
	}
}

func (repo *BeritaRepoImpl) CreateRequest(Berita domain.Berita) (domain.Berita, error) {
	err := repo.db.Create(&Berita).Error
	if err != nil {
		return domain.Berita{}, err
	}
	return Berita, nil
}

func (repo *BeritaRepoImpl)  GetBeritaList(ctx context.Context, limit int, page int) ([]domain.Berita, int, int, int, *int, *int, error){
	var berita []domain.Berita
    var totalcount int64

    // Dapatkan query dengan filter dan pagination
    query, err := repo.beritarepo.GetBuilder( limit, page)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Filter berdasarkan beritaIDs
    err =  query.Preload("User").Order("id DESC").Find(&berita).Error
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total records dari hasil pencarian, tanpa pagination
    totalQuery, err := repo.beritarepo.GetBuilder( 0, 0) // Tanpa pagination
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung jumlah total records
    err = totalQuery.Model(&domain.Berita{}).Count(&totalcount).Error
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    // Hitung total pages
    totalPages := 1
    if limit > 0 {
        totalPages = int((totalcount + int64(limit) - 1) / int64(limit))
    }

    // Jika page > totalPages, return kosong
    if page > totalPages {
        return nil, int(totalcount), page, totalPages, nil, nil, nil
    }

    currentPage := page

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

    return berita, int(totalcount), currentPage, totalPages, nextPage, prevPage, nil
}

func (repo *BeritaRepoImpl) DelBerita(id int) error {
	if err := repo.db.Delete(&domain.Berita{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BeritaRepoImpl) GetBeritaByid(id int) (domain.Berita, error) {
    var beritadata domain.Berita

    // Gunakan `First` untuk mendapatkan satu entri berdasarkan ID
    err := repo.db.Preload("User").First(&beritadata, "id = ?", id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Berita{}, errors.New("testimonial not found")
        }
        return domain.Berita{}, err
    }

    return beritadata, nil
}


func (repo *BeritaRepoImpl) UpdateBeritaId(id int, berita domain.Berita) (domain.Berita, error) {
    // Periksa apakah berita dengan ID ini ada
    var existingberita domain.Berita
    if err := repo.db.First(&existingberita, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return domain.Berita{}, errors.New("berita not found")
        }
        return domain.Berita{}, err
    }

    // Lakukan pembaruan
    if err := repo.db.Model(&existingberita).Updates(berita).Error; err != nil {
        return domain.Berita{}, errors.New("failed to update berita")
    }

    return berita, nil
}