package modelosLogs

import "gorm.io/gorm"

type OPERATION string

const (
	ASIGNAR_PAQUETE   OPERATION = "ASIGNAR_PAQUETE"
	ASIGNAR_RUTA      OPERATION = "ASIGNAR_RUTA"
	ASIGNAR_VEHICULO  OPERATION = "ASIGNAR_VEHICULO"
	CARGAR_PAQUETE    OPERATION = "CARGAR_PAQUETE"
	COMENZAR_VIAJE    OPERATION = "COMENZAR_VIAJE"
	ENTREGAR_PAQUETE  OPERATION = "ENTREGAR_PAQUETE"
	FINALIZAR_VIAJE   OPERATION = "FINALIZAR_VIAJE"
	ASIGNAR_CONDUCTOR OPERATION = "ASIGNAR_CONDUCTOR"
	CREAR_TICKET      OPERATION = "CREAR_TICKET"
	CERRAR_TICKET     OPERATION = "CERRAR_TICKET"
	RESOLVER_TICKET   OPERATION = "RESOLVER_TICKET"
	ACTUALIZAR_TICKET OPERATION = "ACTUALIZAR_TICKET"
)

type Log struct {
	gorm.Model

	Id_usuario     int    `gorm:"not null" json:"id_usuario"`
	Nombre_usuario string `gorm:"not null" json:"username"`

	Descripcion string    `json:"descripcion"`
	Accion      OPERATION `json:"accion"`
	Relevancia  int       `json:"relevancia"`
}

func IsValidAction(s string) bool{
	var validOperations = map[OPERATION]struct{}{
		ASIGNAR_PAQUETE:   {},
		ASIGNAR_RUTA:      {},
		ASIGNAR_VEHICULO:  {},
		CARGAR_PAQUETE:    {},
		COMENZAR_VIAJE:    {},
		ENTREGAR_PAQUETE:  {},
		FINALIZAR_VIAJE:   {},
		ASIGNAR_CONDUCTOR: {},
		CREAR_TICKET:      {},
		CERRAR_TICKET:     {},
		RESOLVER_TICKET:   {},
		ACTUALIZAR_TICKET: {},
	}

	_, exists := validOperations[OPERATION(s)]
	return exists
}
