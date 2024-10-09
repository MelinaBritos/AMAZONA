package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func main() {

	baseDeDatos.Conexiondb()
	baseDeDatos.CrearTablas()

	Iniciar()
}
