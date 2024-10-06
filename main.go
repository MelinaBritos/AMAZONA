package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora"
	Proveedor "github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/Endpoints"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios"
)

func main() {

	Bitacora.Iniciar()
	Proveedor.Iniciar()
	Usuarios.Iniciar()
}
