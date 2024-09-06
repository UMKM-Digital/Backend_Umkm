package aboutusservice

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"umkm/helper"
	domain "umkm/model/domain/homepage"
	entity "umkm/model/entity/homepage/brandlogo"
	web "umkm/model/web/homepage"
	aboutusrepo "umkm/repository/homepage/aboutus"
)

type AboutUsServiceImpl struct {
	aboutusrepo aboutusrepo.AboutUs
}

func NewAboutUsService(aboutusrepo aboutusrepo.AboutUs) *AboutUsServiceImpl {
	return &AboutUsServiceImpl{
		aboutusrepo: aboutusrepo,
	}
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *AboutUsServiceImpl) CreateAboutUs(aboutus web.CreateAboutUs, file *multipart.FileHeader) (map[string]interface{}, error) {
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
	newBrandlogo := domain.AboutUs{
		Description:  aboutus.Description,
		Image:  logoPath,
	}

	// Menyimpan brand logo baru ke repository
	saveBrandLogo, errSaveBrandLogo := service.aboutusrepo.CreatedAboutUs(newBrandlogo)
	if errSaveBrandLogo != nil {
		return nil, errSaveBrandLogo
	}

	return map[string]interface{}{
		"nama_barang": saveBrandLogo.Description,
		"gambar":      saveBrandLogo.Image,
	}, nil
}

// list brand logo
func (service *AboutUsServiceImpl) GetAboutUs() ([]entity.AboutUsEntity, error) {
	getAboutUs, err := service.aboutusrepo.GetAboutUs()

	if err != nil {
		return nil, err
	}

	return entity.ToAboutEntities(getAboutUs), nil
}

func (service *AboutUsServiceImpl) GetAboutUsid(id int) (entity.AboutUsEntity, error) {
	getAboutUs, errgetAboutUs := service.aboutusrepo.FindByAboutId(id)

	if errgetAboutUs != nil {
		return entity.AboutUsEntity{}, errgetAboutUs
	}

	return entity.ToAboutUsEntity(getAboutUs), nil
}

func (service *AboutUsServiceImpl) UpdateAboutUs(request web.UpdateAboutUs, Id int, file *multipart.FileHeader) (map[string]interface{}, error) {
	// Ambil data testimonial berdasarkan ID
	getAboutUsById, err := service.aboutusrepo.FindByAboutId(Id)
	if err != nil {
		return nil, err
	}

	// Gunakan nilai yang ada jika tidak ada perubahan dalam request

	if request.Description == "" {
		request.Description = getAboutUsById.Description
	}

	var Logo string
	if file != nil {
		// Hapus gambar lama jika aboutusrepo
		if getAboutUsById.Image != "" {
			err := os.Remove(getAboutUsById.Image)
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
		Logo = getAboutUsById.Image
	}

	// Buat objek Testimonal baru untuk pembaruan
	AboutUsRequest := domain.AboutUs{
		Id:        Id,
		Description: request.Description,
		Image: Logo,
	}

	// Update testimonial
	UpdateAboutUs, errUpdate := service.aboutusrepo.UpdateAboutUsId(Id, AboutUsRequest)
	if errUpdate != nil {
		return nil, errUpdate
	}
	// Bentuk respons yang akan dikembalikan
	response := map[string]interface{}{
		"description":       UpdateAboutUs.Description,
		"image": UpdateAboutUs.Image,
	}
	return response, nil
}
