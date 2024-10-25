package baseDeDatos

import (
	"fmt"
	"log"
	"os"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/modelosLogs"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
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
	DB.AutoMigrate(modelosBitacora.RepuestoUtilizado{})
	DB.AutoMigrate(modelosBitacora.Ticket{})
	DB.AutoMigrate(modelosBitacora.HistorialCompras{})
	DB.AutoMigrate(modelosPaquete.Paquete{})
	DB.AutoMigrate(modelosLogs.Log{})
}

func CrearFKS() error {

	query := `
    ALTER TABLE tickets ADD CONSTRAINT matriculaFK FOREIGN KEY (matricula) REFERENCES vehiculos(matricula);
    ALTER TABLE tickets ADD CONSTRAINT usernameFK FOREIGN KEY (username) REFERENCES usuarios(username);
    `
	DB.Exec(query)
	// Alter para la tabla tickets
	if err := DB.Exec(`ALTER TABLE tickets ADD CONSTRAINT matriculaFK FOREIGN KEY (matricula) REFERENCES vehiculos(matricula)`).Error; err != nil {
		return fmt.Errorf("error creando FK matricula en tickets: %w", err)
	}

	if err := DB.Exec(`ALTER TABLE tickets ADD CONSTRAINT usernameFK FOREIGN KEY (username) REFERENCES usuarios(username)`).Error; err != nil {
		return fmt.Errorf("error creando FK username en tickets: %w", err)
	}

	// Alter para la tabla repuesto_utilizados
	if err := DB.Exec(`ALTER TABLE repuesto_utilizados ADD CONSTRAINT id_RepuestoFK FOREIGN KEY (id_repuesto) REFERENCES repuestos(ID)`).Error; err != nil {
		return fmt.Errorf("error creando FK id_repuesto en repuesto_utilizados: %w", err)
	}

	// Alters para la tabla paquetes
	if err := DB.Exec(`ALTER TABLE paquetes MODIFY COLUMN id_viaje INT NULL`).Error; err != nil {
		return fmt.Errorf("error modificando id_viaje en paquetes: %w", err)
	}
	if err := DB.Exec(`ALTER TABLE paquetes ADD CONSTRAINT fk_id_viaje FOREIGN KEY (id_viaje) REFERENCES viajes(id) ON UPDATE CASCADE ON DELETE SET NULL`).Error; err != nil {
		return fmt.Errorf("error creando FK id_viaje en paquetes: %w", err)
	}

	if err := DB.Exec(`ALTER TABLE paquetes MODIFY COLUMN id_conductor INT NULL`).Error; err != nil {
		return fmt.Errorf("error modificando id_conductor en paquetes: %w", err)
	}
	if err := DB.Exec(`ALTER TABLE paquetes ADD CONSTRAINT fk_id_conductor FOREIGN KEY (id_conductor) REFERENCES usuarios(id) ON UPDATE CASCADE ON DELETE SET NULL`).Error; err != nil {
		return fmt.Errorf("error creando FK id_conductor en paquetes: %w", err)
	}

	if err := DB.Exec(`ALTER TABLE paquetes MODIFY COLUMN matricula VARCHAR(255) NULL`).Error; err != nil {
		return fmt.Errorf("error modificando matricula en paquetes: %w", err)
	}
	if err := DB.Exec(`ALTER TABLE paquetes ADD CONSTRAINT fk_matricula FOREIGN KEY (matricula) REFERENCES vehiculos(matricula) ON UPDATE CASCADE ON DELETE SET NULL`).Error; err != nil {
		return fmt.Errorf("error creando FK matricula en paquetes: %w", err)
	}

	return nil
}
