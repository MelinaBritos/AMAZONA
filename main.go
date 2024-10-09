package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelos"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func main() {

	baseDeDatos.Conexiondb()
	baseDeDatos.DB.AutoMigrate(modelos.Proveedor{})

	Bitacora.Iniciar()
	Proveedor.Iniciar()
	Usuarios.Iniciar()
}
