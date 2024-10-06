package Usuarios

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	
    r.HandleFunc("/usuarios", rutas.GetUsuariosHandler).Methods("GET")
    r.HandleFunc("/usuarios/{username}", rutas.GetUsuariosByUsernameHandler).Methods("GET")

   r.HandleFunc("/usuarios", rutas.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
}