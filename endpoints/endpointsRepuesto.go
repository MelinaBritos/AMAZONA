package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func EndpointsRepuesto(r *mux.Router) {

	r.HandleFunc("/repuesto/{id}", rutasProveedor.GetRepuestoHandler).Methods("GET")
	r.HandleFunc("/repuesto", rutasProveedor.GetRepuestosHandler).Methods("GET")
	r.HandleFunc("/repuesto/{id}", rutasProveedor.PutRepuestoHandler).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/repuesto", rutasProveedor.PostRepuestoHandler).Methods("POST")    //crear un usuario
}
