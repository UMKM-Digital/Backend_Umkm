package web

type CreatedBerita struct{
	Title string `validate:"required" json:"title"`
	Image string `validate:"required" json:"image"`
	Content string `validate:"required" json:"content"`
}
type UpdtaedBerita struct{
	Title string `validate:"required" json:"title"`
	Image string `validate:"required" json:"image"`
	Content string `validate:"required" json:"content"`
}