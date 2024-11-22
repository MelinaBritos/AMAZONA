package modelosUsuario

import "gorm.io/gorm"

type ROL string

const (
	GERENTE       ROL = "GERENTE"
	ADMINISTRADOR ROL = "ADMINISTRADOR"
	SUPERVISOR    ROL = "SUPERVISOR"
	CONDUCTOR     ROL = "CONDUCTOR"
	MECANICO      ROL = "MECANICO"
)

type Credencial struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Usuario struct {
	gorm.Model

	Username string `gorm:"primaryKey" json:"username"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Clave    string `json:"clave"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni      string `gorm:"unique;not null" json:"dni"`
	Rol      ROL    `json:"rol"`
}
