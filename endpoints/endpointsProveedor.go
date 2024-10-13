package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func EndpointsProveedor(r *mux.Router) {

	r.HandleFunc("/proveedor/{id_proveedor}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")   //Modificar datos de algun proveedor
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST") //crear un proveedor

}
