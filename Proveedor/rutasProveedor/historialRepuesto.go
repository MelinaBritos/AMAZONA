package rutasProveedor

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/validaciones"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetHistorialRepuestosHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los HistorialRepuestos
	var historialRepuestos []modelosProveedor.HistorialRepuesto
	baseDeDatos.DB.Find(&historialRepuestos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&historialRepuestos)

}

func GetHistorialRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo HistorialRepuesto
	var historialRepuesto modelosProveedor.HistorialRepuesto
	params := mux.Vars(r)
	idRepuesto := params["id_repuesto"]
	idCatalogo := params["id_catalogo"]
	fValidez := params["f_validez"]

	baseDeDatos.DB.Where("id_repuesto = ? AND id_catalogo = ? AND f_validez = ?", idRepuesto, idCatalogo, fValidez).First(&historialRepuesto)

	if historialRepuesto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El HistorialRepuesto no existe"))
		return
	}
	json.NewEncoder(w).Encode(&historialRepuesto)
}

func PostHistorialRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo HistorialRepuesto
	var historialRepuesto modelosProveedor.HistorialRepuesto

	if err := json.NewDecoder(r.Body).Decode(&historialRepuesto); err != nil {
		http.Error(w, "Error al decodificar el HistorialRepuesto: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validaciones.ValidarHistorialRepuesto(historialRepuesto); err != nil {
		http.Error(w, "Datos del HistorialRepuesto invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	if err := tx.Create(&historialRepuesto); err.Error != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el HistorialRepuesto: "+err.Error.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&historialRepuesto)
}

func PutHistorialRepuestoHandler(w http.ResponseWriter, r *http.Request) {

	//aca va la logica para modificar los datos de un HistorialRepuesto
	w.Write([]byte("ola mundo put HistorialRepuesto"))
}
