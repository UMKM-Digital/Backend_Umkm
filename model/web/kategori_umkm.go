package web

type CreateCategoriUmkm struct{
	Name string `validate:"required" json:"name"`
}