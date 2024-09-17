package beritaservice

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
	"umkm/helper"
	"umkm/model/domain"
	"os"
	"log"
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
	beritarepo "umkm/repository/homepage/berita"
)

type BeritaServiceImpl struct {
	Beritarepository beritarepo.BeritaRepo
}

func NewBeritaService(Beritarepository beritarepo.BeritaRepo) *BeritaServiceImpl {
	return &BeritaServiceImpl{
		Beritarepository: Beritarepository,
	}
}

func generateRandomFileName(ext string) string {
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(1000000))
	return randomString + ext
}


func (service *BeritaServiceImpl) CreatedBerita(Berita web.CreatedBerita, file *multipart.FileHeader, userID int) (map[string]interface{}, error) {
    // Membuka file gambar yang diunggah
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open the uploaded file")
	}
	defer src.Close()

	// Menggenerate nama file acak dan menyimpan gambar
	ext := filepath.Ext(file.Filename)
	randomFileName := generateRandomFileName(ext)
	logoPath := filepath.Join("uploads/berita", randomFileName)

	// Simpan gambar
	if err := helper.SaveFile(file, logoPath); err != nil {
		return nil, errors.New("failed to save image")
	}

	// Mengubah path menjadi bentuk yang sesuai
	logoPath = filepath.ToSlash(logoPath)

	// Membuat instance Berita baru dengan AuthorId yang diambil dari token
	NewBerita := domain.Berita{
		Title:    Berita.Title,
		Image:    logoPath,
		Content:  Berita.Content,
		AuthorId: userID,  // Masukkan AuthorId
	}

	// Menyimpan berita ke repository
	saveBerita, errSaveBerita := service.Beritarepository.CreateRequest(NewBerita)
	if errSaveBerita != nil {
		return nil, errSaveBerita
	}

	return map[string]interface{}{
		"title":   saveBerita.Title,
		"image":   saveBerita.Image,
		"content": saveBerita.Content,
	}, nil
}

func(service *BeritaServiceImpl) GetBeritaList(ctx context.Context, limit int, page int) ([]entity.BeritaFilterEntity, int, int, int, *int, *int, error){

    // Fetch UMKM entities based on IDs
    beritaList, totalCount, currentPage, totalPages, nextPage, prevPage, err := service.Beritarepository.GetBeritaList(ctx, limit, page)
    if err != nil {
        return nil, 0, 0, 0, nil, nil, err
    }

    beritaResponses := entity.ToberitafilterEntities(beritaList) // Convert UMKM entities to responses

    return beritaResponses, totalCount, currentPage, totalPages, nextPage, prevPage, nil
}

func (service *BeritaServiceImpl) DeleteBerita(id int) error {

    gambar, err := service.Beritarepository.GetBeritaByid(id)
	if err != nil {
		return err
	}

	// Hapus file gambarr (jika ada)
	filePath := filepath.Clean(gambar.Image)

	// Cek jika file ada sebelum menghapus
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", filePath)
	} else {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file %s: %v", filePath, err)
			return err
		}
	}

	return service.Beritarepository.DelBerita(id)
}

func (service *BeritaServiceImpl) GetBeritaByid(id int) (entity.BeritaFilterEntity, error){
	getBerita, errgetBerita := service.Beritarepository.GetBeritaByid(id)

	if errgetBerita != nil {
		return entity.BeritaFilterEntity{}, errgetBerita
	}

	return entity.ToBeritaFilterEntity(getBerita), nil
}

func(service *BeritaServiceImpl) UpdateBerita(request web.UpdtaedBerita, Id int,file *multipart.FileHeader) (map[string]interface{}, error){
	getBeritaById, err := service.Beritarepository.GetBeritaByid(Id)
    if err != nil {
        return nil, err
    }

    // Gunakan nilai yang ada jika tidak ada perubahan dalam request
  
    if request.Title == "" {
        request.Title = getBeritaById.Title
    }
    if request.Content == "" {
        request.Content = getBeritaById.Content
    }

    var Logo string
    if file != nil {
        // Hapus gambar lama jika ada
        if getBeritaById.Image != "" {
            err := os.Remove(getBeritaById.Image)
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
        Logo = filepath.Join("uploads/berita", randomFileName)

        // Menyimpan file ke server
        if err := helper.SaveFile(file, Logo); err != nil {
            return nil, errors.New("failed to save image")
        }

        // Mengonversi path untuk menggunakan forward slashes
        Logo = filepath.ToSlash(Logo)
    } else {
        // Gunakan gambar lama jika tidak ada gambar baru
        Logo = getBeritaById.Image
    }

    // Buat objek Testimonal baru untuk pembaruan
    BeritaRequest := domain.Berita{
        Id: Id,
		Title: request.Title,
		Content: request.Content,
		Image: Logo,
    }

    // Update testimonial
    UpdateBerita, errUpdate := service.Beritarepository.UpdateBeritaId(Id, BeritaRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }
    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "title":   UpdateBerita.Title,
		"content": UpdateBerita.Content,
		"imgae": UpdateBerita.Image,
    }
    return response, nil
}

