package rutasBitacora

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func GetHistorialHandler(w http.ResponseWriter, r *http.Request) {
	var historial []modelosBitacora.HistorialCompras

	baseDeDatos.DB.Find(&historial)

	if err := baseDeDatos.DB.Preload("RepuestoComprado").Find(&historial).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&historial)
}
