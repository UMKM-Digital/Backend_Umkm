package sliderservice

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
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
	sliderrepo "umkm/repository/homepage/slider"
)

type SliderServiceImpl struct {
	sliderrepository sliderrepo.Slider
}

func NewSliderService(sliderrepository sliderrepo.Slider) *SliderServiceImpl{
	return &SliderServiceImpl{
		sliderrepository: sliderrepository,
	}
}

func GenerateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *SliderServiceImpl) CreateSlider(slider web.CreatedSlider, file *multipart.FileHeader) (map[string]interface{}, error) {
	// Membuka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open the uploaded file")
	}
	defer src.Close()

	// Menghasilkan nama file acak untuk file yang diunggah
	ext := filepath.Ext(file.Filename)
	randomFileName := GenerateRandomFileName(ext)
	logoPath := filepath.Join("uploads/slide", randomFileName)

	// Menyimpan file ke server
	if err := helper.SaveFile(file, logoPath); err != nil {
		return nil, errors.New("failed to save image")
	}

	// Mengonversi path untuk menggunakan forward slashes
	logoPath = filepath.ToSlash(logoPath)

	// Membuat objek Slider baru
	NewSlider := domain.Slider{
		SlideDesc:  slider.SlideDesc,
		SlideTitle: slider.SlideTitle,
		Active:     1,
		Gambar:     logoPath,
	}

	// Menyimpan data slider ke database
	saveSlider, errSaveSlider := service.sliderrepository.Created(NewSlider)
	if errSaveSlider != nil {
		// Jika gagal menyimpan ke database, hapus file yang telah diunggah
		if removeErr := os.Remove(logoPath); removeErr != nil {
			// Jika penghapusan file gagal, log error-nya
			log.Printf("failed to remove image: %v", removeErr)
		}
		return nil, errSaveSlider
	}

	// Jika berhasil, kembalikan respons
	return map[string]interface{}{
		"slide_title": saveSlider.SlideTitle,
		"slide_desc":  saveSlider.SlideDesc,
		"image":       saveSlider.Gambar,
		"active":      saveSlider.Active,
	}, nil
}


func(service *SliderServiceImpl) GetSlider() ([]entity.SliderEntity, error){
    GetSliderList, err := service.sliderrepository.GetSlider()
    if err != nil {
        return nil, err
    }
    return entity.ToSliderEntities(GetSliderList), nil
}

func(service *SliderServiceImpl) GetSliderid(id int) (entity.SliderEntity, error){
	GetSlider, errGetSlider := service.sliderrepository.GetSliderId(id)

	if errGetSlider != nil {
		return entity.SliderEntity{}, errGetSlider
	}

	return entity.ToSliderEntity(GetSlider), nil
}

func(service *SliderServiceImpl) DeleteId(id int) error{
	gambartesti, err :=  service.sliderrepository.GetSliderId(id)
	if err != nil {
		return err
	}

	// Hapus file gambar (jika ada)
	filePath := filepath.Clean(gambartesti.Gambar)

	// Cek jika file ada sebelum menghapus
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", filePath)
	} else {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file %s: %v", filePath, err)
			return err
		}
	}

	return service.sliderrepository.DelSlider(id)
}

func(service *SliderServiceImpl) UpdateTestimonial(request web.UpdateSlider, Id int, file *multipart.FileHeader) (map[string]interface{}, error){
	getSliderById, err := service.sliderrepository.GetSliderId(Id)
    if err != nil {
        return nil, err
    }

    // Gunakan nilai yang ada jika tidak ada perubahan dalam request
    if request.SlideDesc == "" {
        request.SlideDesc = getSliderById.SlideDesc
    }
    if request.SlideTitle == "" {
        request.SlideTitle = getSliderById.SlideTitle
    }

    var Logo string
    if file != nil {
        // Hapus gambar lama jika ada
        if getSliderById.Gambar != "" {
            err := os.Remove(getSliderById.Gambar)
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
        randomFileName := GenerateRandomFileName(ext)
        Logo = filepath.Join("uploads/slide", randomFileName)

        // Menyimpan file ke server
        if err := helper.SaveFile(file, Logo); err != nil {
            return nil, errors.New("failed to save image")
        }

        // Mengonversi path untuk menggunakan forward slashes
        Logo = filepath.ToSlash(Logo)
    } else {
        // Gunakan gambar lama jika tidak ada gambar baru
        Logo = getSliderById.Gambar
    }

    // Buat objek Testimonal baru untuk pembaruan
    SliderRequest := domain.Slider{
        Id:          Id,
        SlideDesc: request.SlideDesc,
		SlideTitle: request.SlideTitle,
		Gambar: Logo,
    }

    // Update testimonial
    UpdateTestimonial, errUpdate := service.sliderrepository.UpdateSliderId(Id, SliderRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }
    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "slide_desc":   UpdateTestimonial.SlideDesc,
        "slide_title": UpdateTestimonial.SlideTitle,
        "gambar": UpdateTestimonial.Gambar,
    }
    return response, nil
}

func (service *SliderServiceImpl) UpdateSliderActive(request web.UpdateActiveSlider, Id int) (map[string]interface{}, error) {
  
    getTestimonialById, err := service.sliderrepository.GetSliderId(Id)
    if err != nil {
        return nil, err
    }

   
    if request.Active == getTestimonialById.Active {
       
        response := map[string]interface{}{
            "active": getTestimonialById.Active,
        }
        return response, nil
    }

   
    errUpdate := service.sliderrepository.UpdateActiveId(Id, request.Active)
    if errUpdate != nil {
        return nil, errUpdate
    }

  
    response := map[string]interface{}{
        "active": request.Active,
    }
    return response, nil
}
