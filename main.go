package main

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/modelosUsuarios"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func main() {

	baseDeDatos.Conexiondb()

	baseDeDatos.DB.AutoMigrate(modelosProveedor.Proveedor{})
	baseDeDatos.DB.AutoMigrate(modelosBitacora.Vehiculo{})
	baseDeDatos.DB.AutoMigrate(modelosUsuarios.Usuario{})

	Bitacora.Iniciar()
	Proveedor.Iniciar()
	Usuarios.Iniciar()

	
}
