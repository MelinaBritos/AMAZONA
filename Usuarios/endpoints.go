package Usuarios

import (
	"fmt"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	// Devolver por cada rol - ejemplo conductores
    r.HandleFunc("/usuarios", rutas.GetUsuariosHandler).Methods("GET")
    r.HandleFunc("/usuarios/{username}", rutas.GetUsuariosByUsernameHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutas.GetUsuariosByRolHandler).Methods("GET")

	r.HandleFunc("/usuarios/{username}", rutas.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/usuarios", rutas.CrearUsuario).Methods("POST") //crear un usuario

	r.HandleFunc("/usuarios/{username}", rutas.EliminarUsuario).Methods("DELETE")

	fmt.Println("listen users at port 3001")
    http.ListenAndServe(":3001", r)
}