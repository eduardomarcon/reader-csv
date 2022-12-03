package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *config

type config struct {
	API  APIConfig
	File FileConfig
}

type APIConfig struct {
	Port string
}

type FileConfig struct {
	Name string
}

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: os.Getenv("API_PORT"),
	}
	cfg.File = FileConfig{
		Name: os.Getenv("NAME_FILE"),
	}
}

func GetAPI() APIConfig {
	return cfg.API
}

func GetFile() FileConfig {
	return cfg.File
}
