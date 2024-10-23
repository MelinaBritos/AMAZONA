package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutasBitacora"
	"github.com/gorilla/mux"
)

func EndpointsTicket(r *mux.Router) {

	r.HandleFunc("/ticket", rutasBitacora.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket/{id}", rutasBitacora.GetTicketHandler).Methods("GET")
	r.HandleFunc("/ticket", rutasBitacora.PostTicketHandler).Methods("POST")
	r.HandleFunc("/ticket", rutasBitacora.PutTicketHandler).Methods("PUT")
	r.HandleFunc("/ticket/{id}", rutasBitacora.DeleteTicketHandler).Methods("DELETE")
}
