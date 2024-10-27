package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/rutasLogs"
	"github.com/gorilla/mux"
)

func EndpointsLogs(r *mux.Router){

	r.HandleFunc("/logs", rutasLogs.GetAll).Methods("GET")
	r.HandleFunc("/logs/{id}",rutasLogs.GetById).Methods("GET")
	
	r.HandleFunc("/logs/create", rutasLogs.CreateLogHandler).Methods("POST")
	r.HandleFunc("/logs/create/Many", rutasLogs.CreateMany).Methods("POST")

	r.HandleFunc("/logs/create/Many", rutasLogs.CreateMany).Methods("POST")

	r.HandleFunc("/logs/edit/{id}", rutasLogs.EditarLog).Methods("PUT")


	r.HandleFunc("/logs/delete/{id}", rutasLogs.BorrarLog).Methods("DELETE")
	r.HandleFunc("/logs/deleteByUsername/{username}", rutasLogs.BorrarLogByUsername).Methods("DELETE")
	r.HandleFunc("/logs/deleteMany", rutasLogs.BorrarLogs).Methods("DELETE")

}