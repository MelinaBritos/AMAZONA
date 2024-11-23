package modelosBitacora

import (
	"time"

	"gorm.io/gorm"
)


type IngresosViaje struct {
	gorm.Model

	IDViaje   int `gorm:"not null"`
	IDPaquete int `gorm:"not null"`
	Fecha     time.Time
	Ingreso   float32
}
