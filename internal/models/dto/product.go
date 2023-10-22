package dto

type SaveProduct struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	// Image       *multipart.FileHeader `validate:"required"`
}

type GetProduct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// ImageUrl    string
}

type UpdateProduct struct {
	Id          string `json:"-" validate:"required,uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Image       *multipart.FileHeader
}
