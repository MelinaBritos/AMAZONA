package Usuarios

import (
	"fmt"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	
    r.HandleFunc("/usuarios", rutas.GetUsuariosHandler).Methods("GET")
    r.HandleFunc("/usuarios/{username}", rutas.GetUsuariosByUsernameHandler).Methods("GET")

   	r.HandleFunc("/usuarios", rutas.EditarContraseña).Methods("PUT") //Modificar datos de algun usuario

	r.HandleFunc("/usuarios", rutas.CrearUsuario).Methods("POST") //crear un usuario

	fmt.Println("listen users at port 3001")
    http.ListenAndServe(":3001", r)
}