package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func LoadConfig() {
	once.Do(func() {
		godotenv.Load()
	})
}

func Get(key, defaultValue string) string {
	LoadConfig()
	value := os.Getenv(key)
	isValueExist := value != ""
	if isValueExist {
		return value
	}
	return defaultValue
}

func GetRequired(key string) string {
	LoadConfig()
	value := os.Getenv(key)
	isValueExist := value != ""
	if !isValueExist {
		log.Fatalf("required environment variable %q is not set", key)
	}
	return value
}
