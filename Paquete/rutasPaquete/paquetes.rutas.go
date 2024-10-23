package rutasPaquete

import (
	"encoding/json"
	"net/http"

	dataPaquete "github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/validaciones"

	"github.com/gorilla/mux"
)

func GetPaquetesHandler(w http.ResponseWriter, r *http.Request) {

	paquetes := dataPaquete.ObtenerPaquetes()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(paquetes); err != nil {
		http.Error(w, "Error al codificar paquetes en JSON", http.StatusInternalServerError)
		return
	}

}

func GetPaqueteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete := params["id"]

	paquete, err := dataPaquete.ObtenerPaquete(id_paquete)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El paquete no existe: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&paquete)
}

func PostPaqueteHandler(w http.ResponseWriter, r *http.Request) {

	var paquetes []modelosPaquete.Paquete

	if err := json.NewDecoder(r.Body).Decode(&paquetes); err != nil {
		http.Error(w, "Error al decodificar los paquetes: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, paquete := range paquetes {
		if err := validaciones.ValidarPaquete(paquete); err != nil {
			http.Error(w, "Datos del paquete inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := dataPaquete.CrearPaquetes(paquetes); err != nil {
		http.Error(w, "Error al crear los paquetes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&paquetes)
}

func PutPaqueteHandler(w http.ResponseWriter, r *http.Request) {
	var paquetesInput []*modelosPaquete.Paquete

	// Decodificar el cuerpo de la solicitud en un slice de paquetes
	if err := json.NewDecoder(r.Body).Decode(&paquetesInput); err != nil {
		http.Error(w, "Error al decodificar los paquetes: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, paquete := range paquetesInput {
		if err := validaciones.ValidarPaquete(*paquete); err != nil {
			http.Error(w, "Datos del paquete inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Actualizar los paquetes en la base de datos
	if err := dataPaquete.ActualizarPaquetes(paquetesInput); err != nil {
		http.Error(w, "Error al actualizar los paquetes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el código de estado a 200 OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Paquetes actualizados"))
}

func DeletePaqueteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete := params["id"]

	if err := dataPaquete.BorrarPaquete(id_paquete); err != nil {
		http.Error(w, "Error al borrar el paquete: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Paquete borrado"))
	w.WriteHeader(http.StatusOK)
}
