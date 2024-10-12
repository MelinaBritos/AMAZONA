package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/gorilla/mux"
)

func EndpointsTicket(r *mux.Router) {

	r.HandleFunc("/ticket", rutas.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket/{id}", rutas.GetTicketHandler).Methods("GET")
	r.HandleFunc("/ticket", rutas.PostTicketHandler).Methods("POST")
}
