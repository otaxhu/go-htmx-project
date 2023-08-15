package settings

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     uint
	Name     string
	Driver   string
}

type Server struct {
	Port      uint
	Framework string
}

func NewDatabase() (Database, error) {
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		return Database{}, ErrEnvVarNotFound
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return Database{}, ErrEnvVarNotFound
	}
	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return Database{}, ErrEnvVarNotFound
	}
	dbPortClean, err := strconv.Atoi(dbPort)
	if err != nil {
		return Database{}, ErrParsingPort
	}
	if dbPortClean <= 0 {
		return Database{}, ErrParsingPort
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return Database{}, ErrEnvVarNotFound
	}
	dbDriver, exists := os.LookupEnv("DB_DRIVER")
	if !exists {
		return Database{}, ErrEnvVarNotFound
	}
	return Database{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Port:     uint(dbPortClean),
		Name:     dbName,
		Driver:   dbDriver,
	}, nil
}

func NewServer() (Server, error) {
	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		return Server{}, ErrEnvVarNotFound
	}
	serverPortClean, err := strconv.Atoi(serverPort)
	if err != nil {
		return Server{}, ErrParsingPort
	}
	if serverPortClean <= 0 {
		return Server{}, ErrParsingPort
	}
	serverFramework, exists := os.LookupEnv("SERVER_FRAMEWORK")
	if !exists {
		return Server{}, ErrEnvVarNotFound
	}
	return Server{
		Port:      uint(serverPortClean),
		Framework: serverFramework,
	}, nil
}
