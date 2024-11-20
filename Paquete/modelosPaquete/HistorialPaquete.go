package modelosPaquete

import (
	"time"

	"gorm.io/gorm"
)

type HistorialPaquete struct {
	gorm.Model

	Id_paquete uint
	Estado     Estado    `gorm:"not null"`
	Fecha      time.Time `gorm:"type:timestamp;not null"`
}
