package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora"
	Proveedor "github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos/Proveedor/Endpoints"
)

func main() {

	Bitacora.Iniciar()
	Proveedor.Iniciar()
}
