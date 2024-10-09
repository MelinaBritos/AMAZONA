package Usuarios

import (
	"fmt"
	"net/http"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()
	usr:= r.PathPrefix("/usuarios").Subrouter()

    usr.HandleFunc("", rutas.GetUsuariosHandler).Methods("GET")
    usr.HandleFunc("/{id}", rutas.GetUsuarioByIdHandler).Methods("GET")
	usr.HandleFunc("/roles/{rol}", rutas.GetUsuariosByRolHandler).Methods("GET")
	
	usr.HandleFunc("/{id}", rutas.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	usr.HandleFunc("", rutas.CrearUsuario).Methods("POST") //crear un usuario
	usr.HandleFunc("/login", rutas.Loguearse).Methods("POST")

	usr.HandleFunc("/{id}", rutas.EliminarUsuario).Methods("DELETE")


	fmt.Println("listen users at port 3001")
    http.ListenAndServe(":3001", r)
}