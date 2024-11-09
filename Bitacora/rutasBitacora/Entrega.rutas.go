package rutasBitacora

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func GetEntregasHandler(w http.ResponseWriter, r *http.Request) {
	var entregas []modelosBitacora.Entrega
	baseDeDatos.DB.Find(&entregas)

	json.NewEncoder(w).Encode(&entregas)
}
