package modelosBitacora

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"gorm.io/gorm"
)

type Viaje struct {
	gorm.Model

	UsernameConductor string `gorm:"not null"`
	Matricula         string `gorm:"not null"`
	Estado            string `gorm:"not null"`
	FechaAsignacion   string `gorm:"not null"`
	FechaInicio       string
	Costo             float32
	FechaFinalizacion string
	Paquetes          []modelosPaquete.Paquete `gorm:"foreignKey:Id_viaje;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
