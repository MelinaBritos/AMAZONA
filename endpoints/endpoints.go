package endpoints

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GenerarEndpoints() {

	r := mux.NewRouter()

	port, err := CargarPuertoV2()

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
	EndpointsViaje(r)
	EndpointsEntrega(r)
	EndpointsLocalidad(r)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	http.ListenAndServe(":"+port, corsHandler(r))
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
