package models

import (
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	ID        uint   `gorm:"primaryKey"`
	Nombre    string `json:"nombre" gorm:"unique;not null"`
	Clave     string `json:"clave" gorm:"not null"` // Hasheada con bcrypt
	Rol       string `json:"rol" gorm:"not null"`   // Validado en el controlador
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
