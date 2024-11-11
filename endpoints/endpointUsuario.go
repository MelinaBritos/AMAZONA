package endpoints

import (
	UserService "github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/rutasUsuario"
	"github.com/gorilla/mux"
)



func EndpointsUsuario(r *mux.Router) {

	r.HandleFunc("/usuarios", UserService.GetUsuariosHandler).Methods("GET")
	r.HandleFunc("/usuarios/{username}", UserService.GetByUsername).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", UserService.GetByRol).Methods("GET")

	r.HandleFunc("/usuarios/{username}", UserService.Editar).Methods("PUT") 
	r.HandleFunc("/usuarios/modify/updateMany", UserService.EditMany).Methods("PUT")

	r.HandleFunc("/usuarios/create", UserService.Crear).Methods("POST")            
	r.HandleFunc("/usuarios/create/Many", UserService.CreateMany).Methods("POST")

	r.HandleFunc("/usuarios/login", UserService.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}", UserService.Eliminar).Methods("DELETE")
	r.HandleFunc("/usuarios/deshabilitar/{username}", UserService.Deshabilitar).Methods("DELETE")
	r.HandleFunc("/usuarios/delete/deleteMany", UserService.EliminarMuchos).Methods("DELETE")
}
