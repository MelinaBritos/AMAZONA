package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/rutasPaquete"
	"github.com/gorilla/mux"
)

func EndpointsPaquete(r *mux.Router) {

	r.HandleFunc("/paquete/{id}", rutasPaquete.GetPaqueteHandler).Methods("GET")
	r.HandleFunc("/paquetes", rutasPaquete.GetPaquetesHandler).Methods("GET")
	r.HandleFunc("/paquete", rutasPaquete.PutPaqueteHandler).Methods("PUT")
	r.HandleFunc("/paquete", rutasPaquete.PostPaqueteHandler).Methods("POST")
	r.HandleFunc("/paquete/{id}", rutasPaquete.DeletePaqueteHandler).Methods("DELETE")
	//r.HandleFunc("/paquete/conductor/{id}", rutasPaquete.GetPaquetesAsignadosAConductorHandler).Methods("GET")
	r.HandleFunc("/paquetes/sin_asignar", rutasPaquete.GetPaquetesSinAsignar).Methods("GET")
	r.HandleFunc("/paquete/historial/{id}", rutasPaquete.GetHistorialPaqueteHandler).Methods("GET")

}
