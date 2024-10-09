package baseDeDatos

import (
	"log"
	"os"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/modelosUsuarios"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conexiondb() {
	var err error

	DSN, err := ObtenerDSN()

	if err != nil {
		log.Fatal(err)
		return
	}

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Base de datos conectada")
	}
}

func ObtenerDSN() (string, error) {
	err := godotenv.Load("../TP-Principal-AMAZONA/.env.example")
	if err != nil {
		return "", err
	}
	return os.Getenv("DSN"), nil
}

func CrearTablas() {

	DB.AutoMigrate(modelosProveedor.Proveedor{})
	DB.AutoMigrate(modelosBitacora.Vehiculo{})
	DB.AutoMigrate(modelosUsuarios.Usuario{})

}
