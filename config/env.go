package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init 加载环境变量
func init() {
	if name := os.Getenv("APP_NAME"); name != "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file!")
		}
	}
}
