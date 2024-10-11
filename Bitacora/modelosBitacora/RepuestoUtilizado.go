package modelosBitacora

import "gorm.io/gorm"

type RepuestoUtilizado struct {
	gorm.Model

	IDTicket   int     `gorm:"not null"`
	IDRepuesto int     `gorm:"not null"`
	Cantidad   int     `gorm:"not null"`
	Costo      float32 `gorm:"not null"`
}
