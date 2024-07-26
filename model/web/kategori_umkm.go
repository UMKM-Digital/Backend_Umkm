package web

type CreateCategoriUmkm struct{
	Name string `validate:"required" json:"name"`
}

type UpdateCategoriUmkm struct{
	Name string `validate:"required" json:"name"`
} 