package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsEntrega(r *mux.Router) {
	r.HandleFunc("/entregas", rutasBitacora.GetEntregasHandler).Methods("GET")
}
