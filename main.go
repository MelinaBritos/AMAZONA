package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/endpoints"
)

func main() {

	baseDeDatos.Conexiondb()
	baseDeDatos.CrearTablas()
	//baseDeDatos.CrearFKS()
	endpoints.GenerarEndpoints()
}
