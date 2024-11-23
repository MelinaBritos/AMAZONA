package modelosBitacora

import (
	"time"

	"gorm.io/gorm"
)

type Entrega struct {
	gorm.Model

	IDViaje           int    `gorm:"not null"`
	IDPaquete         int    `gorm:"not null"`
	UsernameConductor string `gorm:"not null"`
	DireccionEntrega  string
	FechaEntrega      time.Time `gorm:"not null;type:date"`
}
