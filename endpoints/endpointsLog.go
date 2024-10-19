package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/rutasLogs"
	"github.com/gorilla/mux"
)

func EndpointsLogs(r *mux.Router){

	r.HandleFunc("/logs", rutasLogs.GetAll).Methods("GET")
	r.HandleFunc("/logs/{id}",rutasLogs.GetById).Methods("GET")
	r.HandleFunc("/logs/users/{username}", rutasLogs.GetByUsername).Methods("GET")

	r.HandleFunc("/logs", rutasLogs.EditarLog).Methods("PUT")

	r.HandleFunc("/logs", rutasLogs.BorrarLog).Methods("DELETE")

}