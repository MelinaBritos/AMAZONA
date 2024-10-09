package Bitacora

import (
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	r.HandleFunc("/vehiculos", rutas.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutas.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutas.PostVehiculoHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}
