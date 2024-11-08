package modelosBitacora

import "gorm.io/gorm"

type Entrega struct {
	gorm.Model

	IDViaje           int    `gorm:"not null"`
	IDPaquete         int    `gorm:"not null"`
	UsernameConductor string `gorm:"not null"`
	DireccionEntrega  string
	FechaEntrega      string `gorm:"not null"`
}
