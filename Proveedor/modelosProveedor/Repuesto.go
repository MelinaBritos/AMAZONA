package modelosProveedor

import (
	"gorm.io/gorm"
)

type Repuesto struct {
	gorm.Model

	Id_catalogo        int    `gorm:"not null"`
	Nombre             string `gorm:"not null"`
	Stock              int
	Stock_minimo       int
	Cantidad_a_comprar int
	Costo              float32 `gorm:"not null"`
	Descripcion        string
}
