package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func EndpointsCatalogo(r *mux.Router) {

	r.HandleFunc("/catalogo/{id_catalogo}", rutasProveedor.GetCatalogoHandler).Methods("GET")
	r.HandleFunc("/catalogo", rutasProveedor.GetCatalogosHandler).Methods("GET")
	r.HandleFunc("/catalogo", rutasProveedor.PutCatalogoHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/catalogo", rutasProveedor.PostCatalogoHandler).Methods("POST") //crear un usuario

}
