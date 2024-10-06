package modelos

import (
	"gorm.io/gorm"
)

type Catalogo struct {
	gorm.Model

	Id_catalogo  int       `gorm:"unique;not null"`
	Id_proveedor int       `gorm:"not null"`
	Proveedor    Proveedor `gorm:"foreignKey:Id_proveedor;references:Id_proveedor"`
	Mes_vigencia string    `gorm:"not null"` //formado aaaa/mm
}
