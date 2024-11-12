package endpoints

import (
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/rutasLocalidad"
	"github.com/gorilla/mux"
)

func EndpointsLocalidad(r *mux.Router) {

	r.HandleFunc("/localidad/{id}", rutasLocalidad.GetLocalidadHandler).Methods("GET")
	r.HandleFunc("/localidades", rutasLocalidad.GetLocalidadesHandler).Methods("GET")
	r.HandleFunc("/localidad", rutasLocalidad.PutLocalidadHandler).Methods("PUT")                   //Modificar datos de algun localidad
	r.HandleFunc("/localidad", rutasLocalidad.PostLocalidadHandler).Methods("POST")                 //crear un localidad
	r.HandleFunc("/localidad/{id}", rutasLocalidad.DeleteLocalidadHandler).Methods("DELETE")        //crear un localidad
	r.HandleFunc("/localidades/{zona}", rutasLocalidad.GetLocalidadesPorZonaHandler).Methods("GET") //crear un localidad

}
