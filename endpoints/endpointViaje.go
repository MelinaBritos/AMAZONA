package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsViaje(r *mux.Router) {

	r.HandleFunc("/viaje", rutasBitacora.GetViajesHandler).Methods("GET")
	r.HandleFunc("/viaje/{id}", rutasBitacora.GetViajeHandler).Methods("GET")
	r.HandleFunc("/viaje", rutasBitacora.PostViajeHandler).Methods("POST")
	r.HandleFunc("/viajeIniciado", rutasBitacora.PutViajeIniciadoHandler).Methods("PUT")
	r.HandleFunc("/viajeFinalizado", rutasBitacora.PutViajeFinalizadoHandler).Methods("PUT")
	r.HandleFunc("/viaje/{id}", rutasBitacora.DeleteViajeHandler).Methods("DELETE")
}
