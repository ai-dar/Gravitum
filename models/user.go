package models

import "time"

// User представляет модель пользователя.
// Благодаря GORM теги, изменения в модели (например, добавление новых полей)
// автоматически отражаются в PostgreSQL при использовании AutoMigrate.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
