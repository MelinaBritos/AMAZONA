package modelosProveedor

import (
	"gorm.io/gorm"
)

type HistorialRepuesto struct {
	gorm.Model

	Id_repuesto int    `gorm:"primaryKey;not null"`
	Id_catalogo int    `gorm:"primaryKey;not null"`
	F_validez   string `gorm:"primaryKey;not null"`
	Nombre      string `gorm:"not null"`
}
