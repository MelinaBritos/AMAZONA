package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsCostosViaje(r *mux.Router) {

	r.HandleFunc("/costosViaje", rutasBitacora.GetCostosHandler).Methods("GET")
	r.HandleFunc("/costosViaje/{id}", rutasBitacora.GetCostoHandler).Methods("GET")
	r.HandleFunc("/costosViaje", rutasBitacora.PostCostoHandler).Methods("POST")
	r.HandleFunc("/costosViaje", rutasBitacora.PutCostoHandler).Methods("PUT")
	r.HandleFunc("/costosViaje/{id}", rutasBitacora.DeleteCostoHandler).Methods("DELETE")
}
