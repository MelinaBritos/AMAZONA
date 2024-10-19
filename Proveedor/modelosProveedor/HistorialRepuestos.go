package modelosProveedor

import (
	"gorm.io/gorm"
)

type HistorialRepuesto struct {
	gorm.Model

	Id_repuesto int    `gorm:"not null"`
	Id_catalogo int    `gorm:"not null"`
	F_validez   string `gorm:"not null"`
	Nombre      string `gorm:"not null"`
}
