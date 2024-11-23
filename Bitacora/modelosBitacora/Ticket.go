package modelosBitacora

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	Username              string    `gorm:"not null"`
	MotivoIngreso         string    `gorm:"not null"`
	FechaCreacion         time.Time `gorm:"not null;type:date"`
	Estado                string    `gorm:"not null"`
	Tipo                  string    `gorm:"not null"`
	Matricula             string    `gorm:"not null"`
	CostoTotal            float32
	DescripcionReparacion string
	FechaFinalizacion     time.Time           `gorm:"not null;type:date"`
	Repuestos             []RepuestoUtilizado `gorm:"foreignKey:IDTicket;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
