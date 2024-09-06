package web

type CreateAboutUs struct{
	Image string `validate:"required" json:"image"`
	Description string `validate:"required" json:"description"`
}

type UpdateAboutUs struct{
	Image string `validate:"required" json:"image"`
	Description string `validate:"required" json:"description"`
} 
