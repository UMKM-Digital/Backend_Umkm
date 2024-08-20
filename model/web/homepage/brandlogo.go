package web

type CreatedBrandLogo struct{
	BrandName string `validate:"required" json:"brand_name"`
	BrandLogo string `validate:"required" json:"brand_logo"`
}