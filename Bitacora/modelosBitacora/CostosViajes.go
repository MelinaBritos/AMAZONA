package modelosBitacora

import "gorm.io/gorm"

type CostosViaje struct {
	gorm.Model

	IDViaje              int
	KilometrosRecorridos float32 `gorm:"not null"`
	CostoCombustible     float32 `gorm:"not null"`
	Peajes               float32
	GastosVarios         float32
}
