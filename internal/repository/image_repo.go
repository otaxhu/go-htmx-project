package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/otaxhu/go-htmx-project/config"
)

//go:generate mockery --name ImageRepository
type ImageRepository interface {
	SaveImage(image *multipart.FileHeader) error
	DeleteImage(imageUrl string) error
}

type imageRepositoryImpl struct{}

func NewImageRepository(imageRepoCfg config.ImageRepo) ImageRepository {
	return &imageRepositoryImpl{}
}

func (repo *imageRepositoryImpl) SaveImage(image *multipart.FileHeader) error {
	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	imagePath, err := filepath.Abs("./public/images/products/" + image.Filename)
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
	imagePath, err := filepath.Abs("./public/images/products/" + filename)
	if err != nil {
		return err
	}
	return os.Remove(imagePath)
}
