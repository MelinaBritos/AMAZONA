package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/rutasUsuario"
	"github.com/gorilla/mux"
)

func EndpointsUsuario(r *mux.Router) {

	r.HandleFunc("/usuarios", rutasUsuario.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios/{username}", rutasUsuario.GetUsuarioByUsernameHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutasUsuario.GetUsuariosByRolHandler).Methods("GET")

	r.HandleFunc("/usuarios/{username}", rutasUsuario.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/usuarios", rutasUsuario.CrearUsuario).Methods("POST")            //crear un usuario
	r.HandleFunc("/usuarios/login", rutasUsuario.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}", rutasUsuario.EliminarUsuario).Methods("DELETE")
}
