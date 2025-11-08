package models

import (
	"time"

	"gorm.io/gorm"
)

type Venta struct {
	ID          uint    `gorm:"primaryKey"`
	UsuarioID   uint    `json:"usuario_id"`
	ProductoID  uint    `json:"producto_id"`
	Cantidad    int     `json:"cantidad"`
	Descuento   float64 `json:"descuento"`
	PrecioFinal float64 `json:"precio_final"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
