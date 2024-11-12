package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func EndpointsProveedor(r *mux.Router) {

	r.HandleFunc("/proveedor/{id}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedores", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")            //Modificar datos de algun proveedor
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST")          //crear un proveedor
	r.HandleFunc("/proveedor/{id}", rutasProveedor.DeleteProveedorHandler).Methods("DELETE") //crear un proveedor

}
