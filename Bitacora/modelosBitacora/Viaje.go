package modelosBitacora

import (
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"gorm.io/gorm"
)

type Viaje struct {
	gorm.Model

	UsernameConductor string    `gorm:"not null"`
	Matricula         string    `gorm:"not null"`
	Estado            string    `gorm:"not null"`
	FechaAsignacion   time.Time `gorm:"type:date"`
	FechaInicio       time.Time `gorm:"type:date"`
	Costo             float32
	FechaFinalizacion time.Time                `gorm:"type:date"`
	Paquetes          []modelosPaquete.Paquete `gorm:"foreignKey:Id_viaje;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FechaReservaViaje time.Time                `gorm:"type:date"`
}
