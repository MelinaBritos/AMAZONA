package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func EndpointsRepuesto(r *mux.Router) {

	r.HandleFunc("/repuesto/{id}", rutasProveedor.GetRepuestoHandler).Methods("GET")
	r.HandleFunc("/repuestos", rutasProveedor.GetRepuestosHandler).Methods("GET")
	r.HandleFunc("/repuesto", rutasProveedor.PutRepuestoHandler).Methods("PUT")
	r.HandleFunc("/repuesto", rutasProveedor.PostRepuestoHandler).Methods("POST")
	r.HandleFunc("/repuesto/{id}", rutasProveedor.DeleteRepuestoHandler).Methods("DELETE")
}
