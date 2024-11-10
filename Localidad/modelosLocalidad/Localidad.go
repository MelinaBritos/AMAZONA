package modelosLocalidad

import (
	"strconv"

	"gorm.io/gorm"
)

type Zona string

// Slice estático que contiene todos los estados válidos
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

// Método para obtener el ID como string
func (p *Localidad) GetIDAsString() string {
	return strconv.Itoa(int(p.ID))
}

// Función para obtener los estados válidos
func ObtenerZonasValidas() []Zona {
	return zonasValidas
}
