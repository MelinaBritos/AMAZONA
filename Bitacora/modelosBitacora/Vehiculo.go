package modelosBitacora

import (
	"gorm.io/gorm"
)

type Vehiculo struct {
	gorm.Model

	Matricula                 string  `gorm:"unique;not null"`
	Marca                     string  `gorm:"not null"`
	Modelo                    string  `gorm:"not null"`
	Año                       int     `gorm:"not null"`
	FechaIngreso              string  `gorm:"not null"`
	Estado                    string  `gorm:"not null"`
	PesoAdmitido              float32 `gorm:"not null"`
	VolumenAdmitidoMtsCubicos float32 `gorm:"not null"`
	KmRecorridos              int
	EstadoVTV                 string `gorm:"not null"`
	FechaVTV                  string `gorm:"not null"`
}
