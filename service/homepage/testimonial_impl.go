package homepageservice

import (
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
    }

    saveTesttimonial, errSaveTesttimonial := service.testimonalrepository.CreateTestimonial(NewTestimonal)
    if errSaveTesttimonial != nil {
        return nil, errSaveTesttimonial
    }

    return map[string]interface{}{
        "quotes": saveTesttimonial.Quotes, // Ensure field names are correct
        "nama":   saveTesttimonial.Name,
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

