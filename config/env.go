package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// init 加载环境变量
func init() {
	_, filename, _, _ := runtime.Caller(0)
	// fmt.Println("Current test filename: " + filename)

	dir, err := filepath.Abs(filepath.Dir(filename) + "/../")

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)

	err = godotenv.Load(dir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

}
