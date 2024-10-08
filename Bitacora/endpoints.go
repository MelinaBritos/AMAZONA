package Bitacora

import (
	//"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelos"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {
	//baseDeDatos.Conexiondb()

	//baseDeDatos.DB.AutoMigrate(modelos.Vehiculo{})

	r := mux.NewRouter()

	r.HandleFunc("/Vehiculos", rutas.GetVehiculosHandler).Methods("GET")
}
