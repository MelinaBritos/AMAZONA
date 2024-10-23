package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsHistorialCompras(r *mux.Router) {

	r.HandleFunc("/historialCompras", rutasBitacora.GetHistorialHandler).Methods("GET")
}
