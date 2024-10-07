package modelos

import (
	"gorm.io/gorm"
)

type Repuesto struct {
	gorm.Model

	Id_repuesto        int      `gorm:"unique;not null"`
	Id_catalogo        int      `gorm:"not null"`
	Catalogo           Catalogo `gorm:"foreignKey:Id_catalogo;references:Id_catalogo"`
	Nombre             string   `gorm:"not null"`
	Stock              int
	Stock_minimo       int
	Cantidad_a_comprar int
	Costo              float32 `gorm:"not null"`
	Descripcion        string
}
