package endpoints

import (
	rutasUsuarios "github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutasUsuarios"
	"github.com/gorilla/mux"
)

func EndpointsUsuario(r *mux.Router) {

	r.HandleFunc("/usuarios", rutasUsuarios.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios/{username}", rutasUsuarios.GetUsuarioByIdHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutasUsuarios.GetUsuariosByRolHandler).Methods("GET")

	r.HandleFunc("/usuarios/{username}", rutasUsuarios.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/usuarios", rutasUsuarios.CrearUsuario).Methods("POST")            //crear un usuario
	r.HandleFunc("/usuarios/login", rutasUsuarios.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}", rutasUsuarios.EliminarUsuario).Methods("DELETE")
}
