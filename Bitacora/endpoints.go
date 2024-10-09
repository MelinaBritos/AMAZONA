package Bitacora

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	r.HandleFunc("/Vehiculos", rutas.GetVehiculosHandler).Methods("GET")
}
