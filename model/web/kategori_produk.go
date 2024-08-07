package web

import "github.com/google/uuid"

type CreateCategoriProduk struct {
    UmkmId uuid.UUID `validate:"required" json:"umkm_id"`
    Name   string    `validate:"required" json:"name"`
}
