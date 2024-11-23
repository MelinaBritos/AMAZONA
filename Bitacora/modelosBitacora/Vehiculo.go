package modelosBitacora

import (
	"time"

	"gorm.io/gorm"
)

type Vehiculo struct {
	gorm.Model

	Matricula                 string    `gorm:"unique;not null"`
	Marca                     string    `gorm:"not null"`
	Modelo                    string    `gorm:"not null"`
	Año                       int       `gorm:"not null"`
	FechaIngreso              time.Time `gorm:"not null;type:date"`
	Estado                    string    `gorm:"not null"`
	PesoAdmitido              float32   `gorm:"not null"`
	VolumenAdmitidoMtsCubicos float32   `gorm:"not null"`
	KmRecorridos              int
	EstadoVTV                 string    `gorm:"not null"`
	FechaVTV                  time.Time `gorm:"not null;type:date"`
}
