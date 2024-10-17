package baseDeDatos

import (
	"fmt"
	"log"
	"os"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
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
	DB.AutoMigrate(modelosUsuario.Usuario{})
	DB.AutoMigrate(modelosBitacora.Ticket{})
	DB.AutoMigrate(modelosBitacora.RepuestoUtilizado{})

}

func CrearFKS() {

	query := `
    ALTER TABLE tickets ADD CONSTRAINT matriculaFK FOREIGN KEY (matricula) REFERENCES vehiculos(matricula);
    ALTER TABLE tickets ADD CONSTRAINT usernameFK FOREIGN KEY (username) REFERENCES usuarios(username);
	ALTER TABLE repuesto_utilizados ADD CONSTRAINT id_RepuestoFK FOREIGN KEY (id_repuesto) REFERENCES repuestos(ID);
    `
	DB.Exec(query)

}

func CrearFKS() {

	query := `
    ALTER TABLE tickets ADD CONSTRAINT matriculaFK FOREIGN KEY (matricula) REFERENCES vehiculos(matricula);
    ALTER TABLE tickets ADD CONSTRAINT usernameFK FOREIGN KEY (username) REFERENCES usuarios(username);
    `
	DB.Exec(query)

}
