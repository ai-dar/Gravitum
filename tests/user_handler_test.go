package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite" // заменённый импорт
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"test/handlers"
	"test/models"
	"test/repository"
	"test/service"
)

// SetupRouter инициализирует маршрутизатор с использованием in-memory базы данных (SQLite) для тестирования.
func SetupRouter() *gin.Engine {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to initialize database: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	r.POST("/users/", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	return r
}

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var responseUser models.User
	json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, user.Name, responseUser.Name)
	assert.Equal(t, user.Email, responseUser.Email)
}

func TestGetUserNotFound(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
