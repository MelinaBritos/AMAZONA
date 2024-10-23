package endpoints

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GenerarEndpoints() {

	r := mux.NewRouter()

	port, err := CargarPuerto()

	if err != nil {
		println(err)
	}

	EndpointsVehiculo(r)
	EndpointsProveedor(r)
	EndpointsUsuario(r)
	EndpointsCatalogo(r)
	EndpointsRepuesto(r)
	EndpointsTicket(r)
	EndpointsLogs(r)
	EndpointsHistorialCompras(r)
	EndpointsPaquete(r)

	http.ListenAndServe(":"+port, r)
}

func CargarPuerto() (string, error) {

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
