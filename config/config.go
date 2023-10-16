package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database        Database
	Server          Server
	ImageRepo       ImageRepo
	ProductsService ProductsService
}

type Database struct {
	Url    string
	Driver string
}

type Server struct {
	Port      int
	Framework string
}

type ImageRepo struct {
	Url string
}

type ProductsService struct {
	GetProductsLimit int
}

func New(envVarsPayload []byte) (Config, error) {
	c := Config{}
	envVars, err := godotenv.UnmarshalBytes(envVarsPayload)
	if err != nil {
		return c, err
	}
	serverPortClean, err := strconv.Atoi(envVars["SERVER_PORT"])
	if err != nil {
		return c, err
	}
	getProductsLimitClean, err := strconv.Atoi(envVars["PRODUCTS_SERVICE_GET_PRODUCTS_LIMIT"])
	if err != nil {
		return c, err
	}
	return Config{
		Database: Database{
			Url:    envVars["DB_URL"],
			Driver: envVars["DB_DRIVER"],
		},
		ImageRepo: ImageRepo{
			Url: envVars["IMAGE_REPO_URL"],
		},
		Server: Server{
			Framework: envVars["SERVER_FRAMEWORK"],
			Port:      serverPortClean,
		},
		ProductsService: ProductsService{
			GetProductsLimit: getProductsLimitClean,
		},
	}, nil
}
