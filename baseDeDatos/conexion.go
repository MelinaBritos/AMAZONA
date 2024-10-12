package baseDeDatos

import (
	"fmt"
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
		println("Probando otros metodos...")
	}

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Base de datos conectada")
	}
}

func ObtenerDSNV2() (string, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return "", fmt.Errorf("la variable de entorno 'DSN' no está configurada")
	}
	return dsn, nil
}
func ObtenerDSN() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("DSN"), err
	}
	return os.Getenv("DSN"), nil

}

func CrearTablas() {

	DB.AutoMigrate(modelosProveedor.Proveedor{})
	DB.AutoMigrate(modelosProveedor.Catalogo{})
	DB.AutoMigrate(modelosProveedor.Repuesto{})
	DB.AutoMigrate(modelosBitacora.Vehiculo{})
	DB.AutoMigrate(modelosUsuarios.Usuario{})
	DB.AutoMigrate(modelosBitacora.Ticket{})

}
