package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Loadenv 加载环境变量
func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
