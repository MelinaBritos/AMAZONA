package baseDeDatos

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=amazona-melina-67b1.i.aivencloud.com user=avnadmin password=AVNS_74HfYqNF2KSJQ7aX7i_ dbname=defaultdb port=22433"

var DB *gorm.DB

func Conexiondb() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Base de datos conectada")
	}
}
