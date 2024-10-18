package modelosProveedor

import (
	"gorm.io/gorm"
)

type Proveedor struct {
	gorm.Model

	Nombre_empresa string `gorm:"not null"`
	Mail           string `gorm:"not null"`
	Telefono       string `gorm:"not null"`
}
