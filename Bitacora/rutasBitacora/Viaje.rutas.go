package rutasBitacora

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetViajesHandler(w http.ResponseWriter, r *http.Request) {
	var Viajes []modelosBitacora.Viaje

	baseDeDatos.DB.Find(&Viajes)
	json.NewEncoder(w).Encode(&Viajes)
}

func GetViajeHandler(w http.ResponseWriter, r *http.Request) {
	var Viaje modelosBitacora.Viaje
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&Viaje, parametros["id"])

	if Viaje.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // error 404
		w.Write([]byte("Viaje no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&Viaje)

}

func PostViajeHandler(w http.ResponseWriter, r *http.Request) {
	var Viaje modelosBitacora.Viaje

	if err := json.NewDecoder(r.Body).Decode(&Viaje); err != nil {
		http.Error(w, "Error al decodificar el Viaje: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validarViaje(Viaje); err != nil {
		http.Error(w, "Viaje inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	Viaje.FechaAsignacion = time.Now().Format("02-01-2006")
	ViajeCreado := tx.Create(&Viaje)

	err := ViajeCreado.Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el Viaje: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func validarViaje(viaje modelosBitacora.Viaje) error {

	var vehiculo modelosBitacora.Vehiculo
	err := baseDeDatos.DB.Where("matricula = ?", viaje.Matricula).First(&vehiculo).Error
	if err != nil {
		return errors.New("el vehiculo no existe " + viaje.Matricula)
	}
	var usuario modelosUsuario.Usuario
	err1 := baseDeDatos.DB.Where("username = ?", viaje.UsernameConductor).First(&usuario).Error
	if err1 != nil {
		return errors.New("el usuario no existe " + viaje.UsernameConductor)
	}
	var pesoTotalPaquetes float32
	var volumenTotalPaquetes float32
	for _, paqueteViaje := range viaje.Paquetes {
		var paquete modelosPaquete.Paquete
		err := baseDeDatos.DB.Where("ID = ?", paqueteViaje.ID).First(&paquete).Error
		if err != nil {
			return errors.New("el paquete no existe")
		}

		pesoTotalPaquetes += paqueteViaje.Peso_kg
		volumenTotalPaquetes += paqueteViaje.Tamaño_mts_cubicos
	}
	if pesoTotalPaquetes > vehiculo.PesoAdmitido || volumenTotalPaquetes > vehiculo.VolumenAdmitidoMtsCubicos {
		return errors.New("los paquetes rebasan la capacidad admitida por el vehiculo")
	}

	switch viaje.Estado {
	case "ASIGNADO", "EN CURSO", "FINALIZADO", "NO FINALIZADO":
	default:
		return errors.New("estado invalido")
	}

	return nil
}
