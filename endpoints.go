package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/rutas"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutasProveedor"
	"github.com/joho/godotenv"

	rutasUsuarios "github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/rutasUsuarios"
	"github.com/gorilla/mux"
)


func GenerarEndpoints() {

	r := mux.NewRouter()


	port, err := CargarPuerto()

	if err != nil {
		println(err)
	}

	EndpointsVehiculo(r)
	EndpointsProveedores(r)
	EndpointsUsuarios(r)

	http.ListenAndServe(":" + port, r)
}

func CargarPuerto() (string, error){

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("PORT"), err
	}
	return os.Getenv("PORT"), nil
}

func CargarPuertoV2() (string, error) {

	port := os.Getenv("PORT")

	if port == "" {
		return port, fmt.Errorf("no existe el puerto")
	}
	return port, nil
}

func EndpointsVehiculo(r *mux.Router)  {
	//HANDLERS VEHICULO
	r.HandleFunc("/vehiculos", rutas.GetVehiculosHandler).Methods("GET")
	r.HandleFunc("/vehiculos/{id}", rutas.GetVehiculoHandler).Methods("GET")
	r.HandleFunc("/vehiculos", rutas.PostVehiculoHandler).Methods("POST")
}

func EndpointsProveedores(r *mux.Router)  {
	// HANDLERS PROOVEDOR
	r.HandleFunc("/proveedor/{id}", rutasProveedor.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.GetProveedoresHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutasProveedor.PutProveedorHandler).Methods("PUT")   //Modificar datos de algun usuario
	r.HandleFunc("/proveedor", rutasProveedor.PostProveedorHandler).Methods("POST") //crear un usuario

}

func EndpointsUsuarios(r *mux.Router){

	r.HandleFunc("/usuarios", rutasUsuarios.GetUsuariosHandler).Methods("GET")
    r.HandleFunc("/usuarios/{username}", rutasUsuarios.GetUsuarioByIdHandler).Methods("GET")
	r.HandleFunc("/usuarios/roles/{rol}", rutasUsuarios.GetUsuariosByRolHandler).Methods("GET")
	
	r.HandleFunc("/usuarios/{username}", rutasUsuarios.EditarUsuario).Methods("PUT") //Modificar datos de algun usuario
	r.HandleFunc("/usuarios", rutasUsuarios.CrearUsuario).Methods("POST") //crear un usuario
	r.HandleFunc("/usuarios/login", rutasUsuarios.Loguearse).Methods("POST")

	r.HandleFunc("/usuarios/{username}",rutasUsuarios.EliminarUsuario).Methods("DELETE")
	
}