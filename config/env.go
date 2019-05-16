package config

import (
	"log"

	"github.com/joho/godotenv"
)

// init 加载环境变量
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
