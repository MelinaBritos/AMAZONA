package modelos

type ROL string

const (
	GERENTE       ROL = "GERENTE"
	ADMINISTRADOR ROL = "ADMINISTRADOR"
	SUPERVISOR    ROL = "SUPERVISOR"
	CONDUCTOR     ROL = "CONDUCTOR"
	MECANICO      ROL = "MECANICO"
)

type Credencial struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type Usuario struct {

	Username string  `json:"username"`
	Clave    string  `json:"clave"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Dni      string  `json:"dni"`
	Rol      ROL     `json:"rol"`
}
