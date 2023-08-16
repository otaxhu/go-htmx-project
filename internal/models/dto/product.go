package dto

import "mime/multipart"

type SaveProduct struct {
	Name        string                `validate:"required"`
	Description string                `validate:"required"`
	Image       *multipart.FileHeader `validate:"required"`
}

type GetProduct struct {
	Id          string
	Name        string
	Description string
	ImageUrl    string
}

type UpdateProduct struct {
	Id          string `validate:"uuid"`
	Name        string
	Description string
	Image       *multipart.FileHeader
}
