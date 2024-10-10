package rutas

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

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

func PostVehiculoHandler(w http.ResponseWriter, r *http.Request) {
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

	switch vehiculo.Estado {
	case "NO APTO PARA CIRCULAR", "APTO PARA CIRCULAR", "EN VIAJE", "EN REPARACION", "MANTENIMIENTO", "DESHABILITADO":
	default:
		return errors.New("estado invalido")
	}
	if vehiculo.PesoAdmitido <= 0 {
		return errors.New("peso admitido invalido")
	}
	if vehiculo.VolumenAdmitidoMtsCubicos <= 0 {
		return errors.New("volumen admitido invalido")
	}
	if vehiculo.Año > 2024 || vehiculo.Año < 2000 {
		return errors.New("año invalido")
	}

	matriculaVieja := regexp.MustCompile(`^[A-Z]{3}[0-9]{3}$`)
	matriculaNueva := regexp.MustCompile(`^[A-Z]{2}[0-9]{3}[A-Z]{2}$`)

	if !matriculaVieja.MatchString(vehiculo.Matricula) && !matriculaNueva.MatchString(vehiculo.Matricula) {
		return errors.New("matricula invalida")
	}

	return nil
}
