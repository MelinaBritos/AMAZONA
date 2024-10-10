package main

import (
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	//HANDLERS VEHICULO
	r.HandleFunc("/vehiculos", rutas.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutas.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutas.PostVehiculoHandler).Methods("POST")

	// HANDLERS PROOVEDOR
	r.HandleFunc("/proveedor/{id_proveedor}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST") //crear un usuario

	//HANDLERS USUARIOS

	http.ListenAndServe(":3000", r)
}
