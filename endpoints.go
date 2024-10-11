package main

import (
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"

	rutasUsuarios "github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutasUsuarios"
	"github.com/gorilla/mux"
)

func GenerarEndpoints() {

	r := mux.NewRouter()

	EndpointsVehiculo(r)
	EndpointsProveedores(r)
	EndpointsUsuarios(r)

	http.ListenAndServe(":3000", r)
}

func EndpointsVehiculo(r *mux.Router) {
	//HANDLERS VEHICULO
	r.HandleFunc("/vehiculos", rutas.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutas.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos/marcas", rutas.GetMarcasVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos/marcas/modelos", rutas.GetModelosVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutas.PostVehiculoHandler).Methods("POST")
	//HANDLERS TICKET
	r.HandleFunc("/ticket", rutas.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket/{id}", rutas.GetTicketHandler).Methods("GET")
	r.HandleFunc("/ticket", rutas.PostTicketHandler).Methods("POST")
	//falta put ticket

}

func EndpointsProveedores(r *mux.Router) {
	// HANDLERS PROOVEDOR
	r.HandleFunc("/proveedor/{id_proveedor}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST") //crear un usuario

}

func EndpointsUsuarios(r *mux.Router) {

	r.HandleFunc("/usuarios", rutasUsuarios.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios/{username}", rutasUsuarios.GetUsuarioByIdHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutasUsuarios.GetUsuariosByRolHandler).Methods("GET")

	r.HandleFunc("/usuarios/{username}", rutasUsuarios.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/usuarios", rutasUsuarios.CrearUsuario).Methods("POST")            //crear un usuario
	r.HandleFunc("/usuarios/login", rutasUsuarios.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}", rutasUsuarios.EliminarUsuario).Methods("DELETE")

}
