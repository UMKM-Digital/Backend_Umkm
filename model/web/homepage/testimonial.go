package web

type CreateTestimonial struct{
	Quotes string `validate:"required" json:"quote"`
	Name string `validate:"required" json:"name"`
	Gambar string `validate:"required" json:"gambar_testi"`
}

type UpdateTestimonial struct{
	Name string `validate:"required" json:"name"`
	Quotes string `validate:"required" json:"quote"`
	Gambar string `validate:"required" json:"gambar_testi"`
} 

type UpdateActive struct{
	Active int `validate:"required" json:"active"`
}