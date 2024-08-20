package homepageservice

import (
	"time"
	domain "umkm/model/domain/homepage"
	entity "umkm/model/entity/homepage"
	web "umkm/model/web/homepage"
	testimonialrepo "umkm/repository/homepage"
)

type TestimonalServiceImpl struct {
	testimonalrepository testimonialrepo.Testimonal
}

func NewTestimonialService(testimonalrepository testimonialrepo.Testimonal) *TestimonalServiceImpl {
    return &TestimonalServiceImpl{
        testimonalrepository: testimonalrepository,
    }
}

func (service *TestimonalServiceImpl) CreateTestimonial(testimonal web.CreateTestimonial) (map[string]interface{}, error) {
    NewTestimonal := domain.Testimonal{
       Quotes: testimonal.Quotes,
	   Name: testimonal.Name,
       Active: 0,
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
	return service.testimonalrepository.DelTransaksi(id)
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
func (service *TestimonalServiceImpl) UpdateTestimonial(request web.UpdateTestimonial, Id int) (map[string]interface{}, error) {
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
    
    // Buat objek Testimonal baru untuk pembaruan
    TestimonalRequest := domain.Testimonal{
        Id: Id,
        Name:       request.Name,
        Quotes: request.Quotes,
    }

    // Update testimonial
    UpdateTestimonial, errUpdate := service.testimonalrepository.UpdateTransaksiId(Id, TestimonalRequest)
    if errUpdate != nil {
        return nil, errUpdate
    }

    // Bentuk respons yang akan dikembalikan
    response := map[string]interface{}{
        "name":   UpdateTestimonial.Name,
        "quotes": UpdateTestimonial.Quotes,
    }
    return response, nil
}
