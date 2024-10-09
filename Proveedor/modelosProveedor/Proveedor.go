package modelosProveedor

import (
	"gorm.io/gorm"
)

type Proveedor struct {
	gorm.Model

	Id_proveedor   int    `gorm:"unique;not null"`
	Nombre_empresa string `gorm:"not null"`
	Mail           string `gorm:"not null"`
	Telefono       int    `gorm:"not null"`
}
