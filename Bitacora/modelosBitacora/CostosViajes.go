package modelosBitacora

import "gorm.io/gorm"

type CostosViaje struct {
	gorm.Model

	IDViaje                   int
	KilometrosRecorridosFinal float32 `gorm:"not null"`
	KilometrosEstimados       float32
	CostoCombustibleEstimado  float32 `gorm:"not null"`
	CostoCombustibleFinal     float32
}
