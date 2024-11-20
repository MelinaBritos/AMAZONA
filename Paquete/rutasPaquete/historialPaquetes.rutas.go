package rutasPaquete

import (
	"encoding/json"
	"net/http"

	dataPaquete "github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete"

	"github.com/gorilla/mux"
)

func GetHistorialPaqueteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete := params["id"]

	paquete, err := dataPaquete.ObtenerHistorialPaquete(id_paquete)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El paquete no tiene un historial: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&paquete)
}
