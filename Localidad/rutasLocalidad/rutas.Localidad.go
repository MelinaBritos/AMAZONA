package rutasLocalidad

import (
	"encoding/json"
	"net/http"

	dataLocalidad "github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/modelosLocalidad"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/validaciones"

	"github.com/gorilla/mux"
)

func GetLocalidadesHandler(w http.ResponseWriter, r *http.Request) {
	localidades := dataLocalidad.ObtenerLocalidades()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(localidades); err != nil {
		http.Error(w, "Error al codificar localidades en JSON", http.StatusInternalServerError)
		return
	}
}

func GetLocalidadHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_localidad := params["id"]

	localidad, err := dataLocalidad.ObtenerLocalidad(id_localidad)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("La localidad no existe: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&localidad)
}

func PostLocalidadHandler(w http.ResponseWriter, r *http.Request) {

	var localidades []modelosLocalidad.Localidad

	if err := json.NewDecoder(r.Body).Decode(&localidades); err != nil {
		http.Error(w, "Error al decodificar las localidades: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, localidad := range localidades {
		if err := validaciones.ValidarLocalidad(localidad); err != nil {
			http.Error(w, "Datos de la localidad inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := dataLocalidad.CrearLocalidades(localidades); err != nil {
		http.Error(w, "Error al crear las localidades: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&localidades)
}

func PutLocalidadHandler(w http.ResponseWriter, r *http.Request) {
	var localidadesInput []*modelosLocalidad.Localidad

	// Decodificar el cuerpo de la solicitud en un slice de localidades
	if err := json.NewDecoder(r.Body).Decode(&localidadesInput); err != nil {
		http.Error(w, "Error al decodificar las localidades: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, localidad := range localidadesInput {
		if err := validaciones.ValidarLocalidad(*localidad); err != nil {
			http.Error(w, "Datos de la localidad inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Actualizar los localidades en la base de datos
	if err := dataLocalidad.ActualizarLocalidades(localidadesInput); err != nil {
		http.Error(w, "Error al actualizar las localidades: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el código de estado a 200 OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Localidades actualizadas"))
}

func DeleteLocalidadHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_localidad := params["id"]

	if err := dataLocalidad.BorrarLocalidad(id_localidad); err != nil {
		http.Error(w, "Error al borrar el localidad: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Localidad borrado"))
	w.WriteHeader(http.StatusOK)
}

func GetLocalidadesPorZonaHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	zona := params["zona"]

	localidades, err := dataLocalidad.ObtenerLocalidadesPorZona(zona)
	if err != nil {
		http.Error(w, "Error al obtener los localidades: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Si todo está bien, retornamos los localidades como JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(localidades); err != nil {
		http.Error(w, "Error al codificar las localidades a JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetZonasHandler(w http.ResponseWriter, r *http.Request) {

	zonas := dataLocalidad.ObtenerZonas()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&zonas); err != nil {
		http.Error(w, "Error al codificar zonas en JSON", http.StatusInternalServerError)
		return
	}

}
