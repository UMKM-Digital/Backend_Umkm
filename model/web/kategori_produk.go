package web


type CreateCategoriProduk struct {
    Name   string    `validate:"required" json:"name"`
}


type UpdateCategoriProduk struct {
    Name   string    `validate:"required" json:"nama"`
}
