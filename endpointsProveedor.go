package main

import (
	rutasProveedor "github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedores"
	"github.com/gorilla/mux"
)

func EndpointsProveedores(r *mux.Router) {
	// HANDLERS PROOVEDOR

	endpointsDeProveedor(r)
	endpointsDeCatalogo(r)
	endpointsDeRepuesto(r)

}

func endpointsDeProveedor(r *mux.Router) {
	r.HandleFunc("/proveedor/{id_proveedor}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")   //Modificar datos de algun proveedor
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST") //crear un proveedor
}

func endpointsDeCatalogo(r *mux.Router) {
	r.HandleFunc("/catalogo/{id_catalogo}", rutasProveedor.GetCatalogoHandler).Methods("GET")
	r.HandleFunc("/catalogo", rutasProveedor.GetCatalogosHandler).Methods("GET")
	r.HandleFunc("/catalogo", rutasProveedor.PutCatalogoHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/catalogo", rutasProveedor.PostCatalogoHandler).Methods("POST") //crear un usuario
}

func endpointsDeRepuesto(r *mux.Router) {
	r.HandleFunc("/repuesto/{id_repuesto}", rutasProveedor.GetRepuestoHandler).Methods("GET")
	r.HandleFunc("/repuesto", rutasProveedor.GetRepuestosHandler).Methods("GET")
	r.HandleFunc("/repuesto", rutasProveedor.PutRepuestoHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/repuesto", rutasProveedor.PostRepuestoHandler).Methods("POST") //crear un usuario
}
