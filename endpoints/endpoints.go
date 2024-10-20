package endpoints

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	corsObj := handlers.AllowedOrigins([]string{"*"}) // Permitir todos los orígenes
    corsHeaders := handlers.AllowedHeaders([]string{
        "Content-Type",
        "X-Requested-With",
        "Authorization", // Agregar Authorization en los headers permitidos
    })
    corsMethods := handlers.AllowedMethods([]string{
        "GET",
        "POST",
        "PUT",      // Agregar PUT
        "DELETE",   // Agregar DELETE
        "OPTIONS",  // Agregar OPTIONS
    })

	http.ListenAndServe(":"+ port, handlers.CORS(corsObj, corsHeaders, corsMethods)(r))
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
