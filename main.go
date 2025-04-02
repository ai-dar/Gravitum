package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"test/handlers"

	"test/models"

	"test/repository"

	"test/routes"

	"test/service"
)

func main() {
	// Чтение строки подключения из переменных окружения
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=1234 dbname=users_db port=5432 sslmode=disable"
	}

	// Подключение к PostgreSQL с использованием GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Автоматическая миграция моделей. При изменении структуры моделей GORM обновит схему БД.
	db.AutoMigrate(&models.User{})

	// Инициализация слоёв приложения
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализация маршрутизатора и регистрация роутов
	router := gin.Default()
	routes.RegisterRoutes(router, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
