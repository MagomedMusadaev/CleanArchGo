package main

import (
	"CleanArchitectureGo/internal/handler"
	"CleanArchitectureGo/internal/repo"
	"CleanArchitectureGo/internal/service"
	"CleanArchitectureGo/pkg/logg"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	myhttp "CleanArchitectureGo/api/http"
)

func main() {

	// Загрузка переменных
	if err := godotenv.Load("F:\\CleanArchitectureGo\\.env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла", err)
	}

	// Инициализация логирования
	logg.Logging()

	// Подключение к PostgresSQL
	db := repo.ConnectPostgresDB()
	defer db.Close()

	// Инициализации репо слоя
	repository := repo.NewUserRepository(db)

	// Инициализация сервис слоя
	services := service.NewUserService(repository)

	// Инициализация handler слоя
	handlers := handler.NewUserService(services)

	// инициализации роута
	r := mux.NewRouter()

	// инитим роут
	myhttp.InitRoutes(r, handlers)

	// запускаем сервер
	port := os.Getenv("PORT")
	logg.Info("Сервер запушен на порту: " + port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
