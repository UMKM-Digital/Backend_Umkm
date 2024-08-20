package brandlogoservice

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
	"umkm/helper"
	domain "umkm/model/domain/homepage"
	entity "umkm/model/entity/homepage/brandlogo"
	web "umkm/model/web/homepage"
	brandrepo "umkm/repository/homepage/brandlogo"
)

type BrandLogoServiceImpl struct {
	brandlogorepo brandrepo.Brandlogo
}

func NewBrandLogoService(brandlogorepo brandrepo.Brandlogo) *BrandLogoServiceImpl {
	return &BrandLogoServiceImpl{
		brandlogorepo: brandlogorepo,
	}
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *BrandLogoServiceImpl) CreateBrandlogo(brandlogo web.CreatedBrandLogo, file *multipart.FileHeader) (map[string]interface{}, error) {
	// Membuka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open the uploaded file")
	}
	defer src.Close()

	// Menghasilkan nama file acak untuk file yang diunggah
	ext := filepath.Ext(file.Filename)
	randomFileName := generateRandomFileName(ext)
	logoPath := filepath.Join("uploads/logo", randomFileName)

	// Menyimpan file ke server
	if err := helper.SaveFile(file, logoPath); err != nil {
		return nil, errors.New("failed to save image")
	}

	// Mengonversi path untuk menggunakan forward slashes
	logoPath = filepath.ToSlash(logoPath)

	// Membuat instance baru untuk Brandlogo
	newBrandlogo := domain.Brandlogo{
		BrandName:  brandlogo.BrandName,
		BrandLogo:  logoPath,
		Created_at: time.Now(),
	}

	// Menyimpan brand logo baru ke repository
	saveBrandLogo, errSaveBrandLogo := service.brandlogorepo.CreatedBrandLogo(newBrandlogo)
	if errSaveBrandLogo != nil {
		return nil, errSaveBrandLogo
	}

	return map[string]interface{}{
		"Nama Brand": saveBrandLogo.BrandName,
		"gambar":     saveBrandLogo.BrandLogo,
		"active":     saveBrandLogo.Created_at,
	}, nil
}


//list brand logo
func (service *BrandLogoServiceImpl) GetBrandlogoList() ([]entity.BrandLogoEntity, error) {
	getBrandLogo, err := service.brandlogorepo.GetBrandLogo()

    if err != nil {
        return nil, err
    }

	return entity.ToBrandLogoEntities(getBrandLogo), nil
}