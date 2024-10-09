package rutas

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetVehiculosHandler(w http.ResponseWriter, r *http.Request) {
	var vehiculos []modelosBitacora.Vehiculo

	baseDeDatos.DB.Find(&vehiculos)
	json.NewEncoder(w).Encode(&vehiculos)
}

func GetVehiculoHandler(w http.ResponseWriter, r *http.Request) {
	var vehiculo modelosBitacora.Vehiculo
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&vehiculo, parametros["id"])

	if vehiculo.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // error 404
		w.Write([]byte("Producto no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&vehiculo)

}

func PostProductosHandler(w http.ResponseWriter, r *http.Request) {
	var vehiculos []modelosBitacora.Vehiculo

	if err := json.NewDecoder(r.Body).Decode(&vehiculos); err != nil {
		http.Error(w, "Error al decodificar los vehiculos: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, vehiculo := range vehiculos {
		if err := validarVehiculo(vehiculo); err != nil {
			http.Error(w, "Vehiculo inválido: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	tx := baseDeDatos.DB.Begin()
	for _, vehiculo := range vehiculos {
		vehiculoCreado := tx.Create(&vehiculo)

		err := vehiculoCreado.Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al crear los vehiculos: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func validarVehiculo(vehiculo modelosBitacora.Vehiculo) error {

	// how??
	return nil
}
