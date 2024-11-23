package modelosLocalidad

import (
	"strconv"

	"gorm.io/gorm"
)

type Zona string

var zonasValidas = []Zona{
	CABA,
	ZONA_SUR,
	ZONA_NORTE,
	ZONA_OESTE,
}

const (
	CABA       Zona = "CABA"
	ZONA_SUR   Zona = "ZONA SUR"
	ZONA_NORTE Zona = "ZONA NORTE"
	ZONA_OESTE Zona = "ZONA OESTE"
)

type Localidad struct {
	gorm.Model

	Nombre_localidad string  `gorm:"not null"`
	Zona_pertenencia Zona    `gorm:"not null"`
	Costo_localidad  float32 `gorm:"not null"`
}

func (p *Localidad) GetIDAsString() string {
	return strconv.Itoa(int(p.ID))
}

func ObtenerZonasValidas() []Zona {
	return zonasValidas
}
