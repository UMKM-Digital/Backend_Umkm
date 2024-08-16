package web

type CreateTestimonial struct{
	Quotes string `validate:"required" json:"quote"`
	Name string `validate:"required" json:"name"`
}

type UpdateTestimonial struct{
	Name string `validate:"required" json:"name"`
	Quotes string `validate:"required" json:"quote"`
} 