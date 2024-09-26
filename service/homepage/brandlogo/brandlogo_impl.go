package brandlogoservice

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
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
		CreatedAt: time.Now(),
	}

	// Menyimpan brand logo baru ke repository
	saveBrandLogo, errSaveBrandLogo := service.brandlogorepo.CreatedBrandLogo(newBrandlogo)
	if errSaveBrandLogo != nil {
		return nil, errSaveBrandLogo
	}

	return map[string]interface{}{
		"nama_barang": saveBrandLogo.BrandName,
		"gambar":      saveBrandLogo.BrandLogo,
		"created_at":      saveBrandLogo.CreatedAt,
	}, nil
}

// list brand logo
func (service *BrandLogoServiceImpl) GetBrandlogoList() ([]entity.BrandLogoEntity, error) {
	getBrandLogo, err := service.brandlogorepo.GetBrandLogo()

	if err != nil {
		return nil, err
	}

	return entity.ToBrandLogoEntities(getBrandLogo), nil
}

// delete logo
func (service *BrandLogoServiceImpl) DeleteBrandLogo(id int) error {
	// Cari brand logo berdasarkan ID
	brandlogo, err := service.brandlogorepo.FindById(id)
	if err != nil {
		return err
	}

	// Hapus file gambar (jika ada)
	filePath := filepath.Clean(brandlogo.BrandLogo)

	// Cek jika file ada sebelum menghapus
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", filePath)
	} else {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file %s: %v", filePath, err)
			return err
		}
	}

	// Hapus brand logo dari database
	return service.brandlogorepo.DeleteLogoId(id)
}

//get id
func (service *BrandLogoServiceImpl) GetBrandLogoid(id int) (entity.BrandLogoEntity, error) {
	getBrandlogo, errgetBrandlogo := service.brandlogorepo.FindById(id)

	if errgetBrandlogo != nil {
		return entity.BrandLogoEntity{}, errgetBrandlogo
	}

	return entity.ToBrandEntity(getBrandlogo), nil
}

func (service *BrandLogoServiceImpl) UpdateBrandLogo(request web.UpdateBrandLogo, Id int,file *multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil data testimonial berdasarkan ID
    getBrandLogolById, err := service.brandlogorepo.FindById(Id)
    if err != nil {
        return nil, err
    }

    // Gunakan nilai yang ada jika tidak ada perubahan dalam request
  
    if request.BrandName == "" {
        request.BrandName = getBrandLogolById.BrandName
    }

    var Logo string
    if file != nil {
        // Hapus gambar lama jika ada
        if getBrandLogolById.BrandLogo != "" {
            err := os.Remove(getBrandLogolById.BrandLogo)
            if err != nil {
                return nil, errors.New("failed to remove old image")
            }
        }

        // Proses gambar baru
        src, err := file.Open()
        if err != nil {
            return nil, errors.New("failed to open the uploaded file")
        }
        defer src.Close()

        // Menghasilkan nama file acak untuk file yang diunggah
        ext := filepath.Ext(file.Filename)
        randomFileName := generateRandomFileName(ext)
        Logo = filepath.Join("uploads/logo", randomFileName)

        // Menyimpan file ke server
        if err := helper.SaveFile(file, Logo); err != nil {
            return nil, errors.New("failed to save image")
        }

        // Mengonversi path untuk menggunakan forward slashes
        Logo = filepath.ToSlash(Logo)
    } else {
        // Gunakan gambar lama jika tidak ada gambar baru
        Logo = getBrandLogolById.BrandLogo
    }

    // Buat objek Testimonal baru untuk pembaruan
    BrandLogoRequest := domain.Brandlogo{
        Id: Id,
		BrandName: request.BrandName,
		BrandLogo: Logo,
    }

    // Update testimonial
    UpdateBrandLogo, errUpdate := service.brandlogorepo.UpdateBrandLogoId(Id, BrandLogoRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }
    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "name":   UpdateBrandLogo.BrandName,
		"brand_logo": UpdateBrandLogo.BrandLogo,
    }
    return response, nil
}
