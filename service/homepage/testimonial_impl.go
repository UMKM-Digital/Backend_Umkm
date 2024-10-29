package homepageservice

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
	"umkm/helper"
	domain "umkm/model/domain/homepage"
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
	testimonialrepo "umkm/repository/homepage"
    "os"
)

type TestimonalServiceImpl struct {
	testimonalrepository testimonialrepo.Testimonal
}

func NewTestimonialService(testimonalrepository testimonialrepo.Testimonal) *TestimonalServiceImpl {
    return &TestimonalServiceImpl{
        testimonalrepository: testimonalrepository,
    }
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}

func (service *TestimonalServiceImpl) CreateTestimonial(testimonal web.CreateTestimonial, file *multipart.FileHeader) (map[string]interface{}, error) {
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

    NewTestimonal := domain.Testimonal{
       Quotes: testimonal.Quotes,
	   Name: testimonal.Name,
       Active: 1,
       GambarTesti: logoPath,
       Created_at: time.Now(),
    }

    saveTesttimonial, errSaveTesttimonial := service.testimonalrepository.CreateTestimonial(NewTestimonal)
    if errSaveTesttimonial != nil {
        return nil, errSaveTesttimonial
    }

    return map[string]interface{}{
        "quotes": saveTesttimonial.Quotes, // Ensure field names are correct
        "nama":   saveTesttimonial.Name,
        "active": saveTesttimonial.Active,
        "gambar": saveTesttimonial.GambarTesti,
    }, nil
}

func (service *TestimonalServiceImpl) GetTestimonial() ([]entity.TesttimonialEntity, error) {
    GetTestimonialList, err := service.testimonalrepository.GetTestimonial()
    if err != nil {
        return nil, err
    }
    return entity.ToKategoriProdukEntities(GetTestimonialList), nil
}

//delete
func (service *TestimonalServiceImpl) DeleteTestimonial (id int) error {

    gambartesti, err := service.testimonalrepository.GetTransaksiByid(id)
	if err != nil {
		return err
	}

	// Hapus file gambar (jika ada)
	filePath := filepath.Clean(gambartesti.GambarTesti)

	// Cek jika file ada sebelum menghapus
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", filePath)
	} else {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file %s: %v", filePath, err)
			return err
		}
	}

	return service.testimonalrepository.DelTestimonial(id)
}

//get id
func (service *TestimonalServiceImpl) GetTestimonialid(id int) (entity.TesttimonialEntity, error) {
	GetTestimonial, errGetTestimonial := service.testimonalrepository.GetTransaksiByid(id)

	if errGetTestimonial != nil {
		return entity.TesttimonialEntity{}, errGetTestimonial
	}

	return entity.ToTestimonialEntity(GetTestimonial),nil
}

//update
func (service *TestimonalServiceImpl) UpdateTestimonial(request web.UpdateTestimonial, Id int,file *multipart.FileHeader) (map[string]interface{}, error) {
    // Ambil data testimonial berdasarkan ID
    getTestimonialById, err := service.testimonalrepository.GetTransaksiByid(Id)
    if err != nil {
        return nil, err
    }

    // Gunakan nilai yang ada jika tidak ada perubahan dalam request
    if request.Name == "" {
        request.Name = getTestimonialById.Name
    }
    if request.Quotes == "" {
        request.Quotes = getTestimonialById.Quotes
    }

    var Logo string
    if file != nil {
        // Hapus gambar lama jika ada
        if getTestimonialById.GambarTesti != "" {
            err := os.Remove(getTestimonialById.GambarTesti)
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
        Logo = getTestimonialById.GambarTesti
    }

    // Buat objek Testimonal baru untuk pembaruan
    TestimonalRequest := domain.Testimonal{
        Id:          Id,
        Name:        request.Name,
        Quotes:      request.Quotes,
        GambarTesti: Logo,
    }

    // Update testimonial
    UpdateTestimonial, errUpdate := service.testimonalrepository.UpdateTestimonialId(Id, TestimonalRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }
    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "name":   UpdateTestimonial.Name,
        "quotes": UpdateTestimonial.Quotes,
        "gambar": UpdateTestimonial.GambarTesti,
    }
    return response, nil
}

func (service *TestimonalServiceImpl) GetTestimonialActive() ([]entity.TesttimonialEntity, error) {
    GetTestimonialList, err := service.testimonalrepository.GetTestimonialActive(1)
    if err != nil {
        return nil, err
    }
    if len(GetTestimonialList) == 0 {
        log.Println("No testimonials found with active = 1")
        return nil, nil
    }
    return entity.ToKategoriProdukEntities(GetTestimonialList), nil
}

func (service *TestimonalServiceImpl) UpdateTestimonialActive(request web.UpdateActive, Id int) (map[string]interface{}, error) {
    getTestimonialById, err := service.testimonalrepository.GetTransaksiByid(Id)
    if err != nil {
        return nil, err
    }

    if request.Active == getTestimonialById.Active {
        response := map[string]interface{}{
            "active": getTestimonialById.Active,
        }
        return response, nil
    }

    errUpdate := service.testimonalrepository.UpdateActiveId(Id, request.Active)
    if errUpdate != nil {
        return nil, errUpdate
    }

    response := map[string]interface{}{
        "active": request.Active,
    }
    return response, nil
}

