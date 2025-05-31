package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var (
	Port     string
	DataPath string
	DB       *gorm.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	Port = os.Getenv("SERVE_PORT")
	if Port == "" {
		log.Fatal("SERVE_PORT не задан")
	}

	DataPath = os.Getenv("DATA_PATH")
	if DataPath == "" {
		log.Fatal("DATA_PATH не задан")
	}

	DB = InitialiseDB()

	http.HandleFunc("/", FileHandler)           // это отдаёт сами файлы
	http.HandleFunc("/file-exists", FileExists) // а это API-проверка

	fmt.Println("Starting server on :" + Port)

	_ = http.ListenAndServe(":"+Port, nil)
}
