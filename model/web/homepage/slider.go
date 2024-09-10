package web

type CreatedSlider struct{
	SlideDesc string `validate:"required" json:"slide_desc"`
	SlideTitle string `validate:"required" json:"slide_title"`
	Gambar string `validate:"requuired" json:"gambar"`
}