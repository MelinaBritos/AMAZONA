package Bitacora

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func Iniciar() {
	baseDeDatos.Conexiondb()

	//baseDedatos.DB.AutoMigrate(modelos.Vehiculo{})

	r := mux.NewRouter()
	r.HandleFunc("/Vehiculos", rutas.HomeHandler)
}
