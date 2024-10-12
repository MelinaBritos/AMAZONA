package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/gorilla/mux"
)

func EndpointsVehiculo(r *mux.Router) {

	r.HandleFunc("/vehiculos", rutas.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutas.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos/marcas", rutas.GetMarcasVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos/marcas/modelos", rutas.GetModelosVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutas.PostVehiculoHandler).Methods("POST")
}
