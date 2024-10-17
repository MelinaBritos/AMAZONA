package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsVehiculo(r *mux.Router) {

	r.HandleFunc("/vehiculos", rutasBitacora.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutasBitacora.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/marcas", rutasBitacora.GetMarcasVehiculoHandler).Methods("GET")
	r.HandleFunc("/modelos/{marca}", rutasBitacora.GetModelosVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutasBitacora.PostVehiculoHandler).Methods("POST")
	r.HandleFunc("/vehiculos", rutasBitacora.PutVehiculoHandler).Methods("PUT")
	r.HandleFunc("/vehiculos/{id}", rutasBitacora.DeleteVehiculoHandler).Methods("DELETE")

}
