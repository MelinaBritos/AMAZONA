package modelosBitacora

import (
	"time"

	"gorm.io/gorm"
)

type IngresosViajes struct {
	gorm.Model

	IDViaje   int
	IDPaquete int
	Fecha     time.Time `gorm:"type:date"`
	Ingreso   float32
}
