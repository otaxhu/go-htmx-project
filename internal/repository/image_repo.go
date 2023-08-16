package repository

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

//go:generate mockery --name ImageRepository
type ImageRepository interface {
	SaveImage(image *multipart.FileHeader) error
	DeleteImage(imageUrl string) error
}

type imageRepositoryImpl struct{}

func NewImageRepository() ImageRepository {
	return &imageRepositoryImpl{}
}

func (repo *imageRepositoryImpl) SaveImage(image *multipart.FileHeader) error {
	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	imagePath, err := filepath.Abs(fmt.Sprintf("./static/products/images/%s", image.Filename))
	if err != nil {
		return err
	}
	out, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return err
}

func (repo *imageRepositoryImpl) DeleteImage(imageUrl string) error {
	filename := filepath.Base(imageUrl)
	imagePath, err := filepath.Abs(fmt.Sprintf("./static/products/%s", filename))
	if err != nil {
		return err
	}
	return os.Remove(imagePath)
}