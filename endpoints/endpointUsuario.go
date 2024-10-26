package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/rutasUsuario"
	"github.com/gorilla/mux"
)

func EndpointsUsuario(r *mux.Router) {

	r.HandleFunc("/usuarios", rutasUsuario.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios/{username}", rutasUsuario.GetUsuarioByUsernameHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutasUsuario.GetUsuariosByRolHandler).Methods("GET")

	r.HandleFunc("/usuarios/{username}", rutasUsuario.EditarUsuario).Methods("PUT") 
	r.HandleFunc("/usuarios/modify/updateMany", rutasUsuario.EditarUsuarios).Methods("PUT")

	r.HandleFunc("/usuarios/create", rutasUsuario.CrearUsuario).Methods("POST")            
	r.HandleFunc("/usuarios/create/Many", rutasUsuario.CrearUsuarios).Methods("POST")

	r.HandleFunc("/usuarios/login", rutasUsuario.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}", rutasUsuario.EliminarUsuario).Methods("DELETE")
	r.HandleFunc("/usuarios/delete/deleteMany", rutasUsuario.EliminarUsuarios).Methods("DELETE")
}
