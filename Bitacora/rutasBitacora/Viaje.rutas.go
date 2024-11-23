package rutasBitacora

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	dataPaquete "github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetViajesHandler(w http.ResponseWriter, r *http.Request) {
	var Viajes []modelosBitacora.Viaje
	baseDeDatos.DB.Find(&Viajes)

	if err := baseDeDatos.DB.Preload("Paquetes").Find(&Viajes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	baseDeDatos.DB.Model(&Viaje).Association("Paquetes").Find(&Viaje.Paquetes)
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

	Viaje.FechaAsignacion = time.Now()
	Viaje.Estado = "ASIGNADO"
	ViajeCreado := tx.Create(&Viaje)
	tx.Save(Viaje)

	err := ViajeCreado.Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el Viaje: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, paqueteViaje := range Viaje.Paquetes {
		paquete, err := dataPaquete.ObtenerPaquete(fmt.Sprint(paqueteViaje.ID))
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al encontrar paquete: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := dataPaquete.ActualizarEstadoPaquete(tx, &paquete, modelosPaquete.ASIGNADO); err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := dataPaquete.AsignarViajeAPaquete(tx, Viaje.ID, &paquete); err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		http.Error(w, "Error al confirmar la transacción: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func PutViajeIniciadoHandler(w http.ResponseWriter, r *http.Request) {
	var viaje modelosBitacora.Viaje
	parametros := mux.Vars(r)

	tx := baseDeDatos.DB.Begin()

	baseDeDatos.DB.First(&viaje, parametros["id"])
	if viaje.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Viaje no encontrado"))
		return
	}

<<<<<<< HEAD
	if viaje.FechaReservaViaje != time.Now().Format("02-01-2006") {
		tx.Rollback()
		http.Error(w, "No puede iniciarse el viaje porque no es la fecha asignada", http.StatusInternalServerError)
	}

	viaje.FechaInicio = time.Now().Format("02-01-2006")
=======
	viaje.FechaInicio = time.Now()
>>>>>>> main
	viaje.Estado = "EN CURSO"
	tx.Save(&viaje)

	// vehiculo pasa a estar en viaje
	var vehiculo modelosBitacora.Vehiculo
	err := baseDeDatos.DB.Where("matricula = ?", viaje.Matricula).First(&vehiculo).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al encontrar vehiculo: "+err.Error(), http.StatusInternalServerError)
	}

	if vehiculo.Estado == "EN VIAJE" {
		tx.Rollback()
		mensajeError := fmt.Sprintf("El vehiculo ya esta en viaje. Su estado es: %s y su matricula es : %s", vehiculo.Estado, vehiculo.Matricula)
		http.Error(w, mensajeError, http.StatusInternalServerError)
	}

	vehiculo.Estado = "EN VIAJE"
	tx.Save(&vehiculo)

	// paquetes pasan a estar en viaje
	paquetes, err := dataPaquete.ObtenerPaquetesPorViaje(viaje.ID)
	if err != nil {
		http.Error(w, "Error al obtener los paquetes por viaje: "+err.Error(), http.StatusInternalServerError)
	}

	for _, paquete := range paquetes {
		if err := dataPaquete.ActualizarEstadoPaquete(tx, &paquete, modelosPaquete.EN_VIAJE); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}

func PutViajeFinalizadoHandler(w http.ResponseWriter, r *http.Request) {
	var viaje modelosBitacora.Viaje
	parametros := mux.Vars(r)

	tx := baseDeDatos.DB.Begin()

	baseDeDatos.DB.First(&viaje, parametros["id"])
	if viaje.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Viaje no encontrado"))
		return
	}

	viaje.FechaFinalizacion = time.Now()
	tx.Save(viaje)

	// vehiculo deja de estar en viaje
	var vehiculo modelosBitacora.Vehiculo
	err := baseDeDatos.DB.Where("matricula = ?", viaje.Matricula).First(&vehiculo).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al encontrar vehiculo: "+err.Error(), http.StatusInternalServerError)
	}
	vehiculo.Estado = "APTO PARA CIRCULAR"
	tx.Save(&vehiculo)

	// paquetes pasan a no entregado si no fueron entregados en el viaje
	paquetes, err := dataPaquete.ObtenerPaquetesPorViaje(viaje.ID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al obtener los paquetes por viaje: "+err.Error(), http.StatusInternalServerError)
	}

	for _, paquete := range paquetes {
		if paquete.Estado != modelosPaquete.ENTREGADO {
			dataPaquete.ActualizarEstadoPaquete(tx, &paquete, modelosPaquete.NO_ENTREGADO)
		}
		var entrega modelosBitacora.Entrega
		entrega.IDViaje = int(viaje.ID)
		entrega.IDPaquete = int(paquete.ID)
		entrega.UsernameConductor = viaje.UsernameConductor
		entrega.DireccionEntrega = paquete.Dir_entrega
		entrega.FechaEntrega = time.Now()

		entregaCreada := tx.Create(&entrega)
		err := entregaCreada.Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al registrar entrega: "+err.Error(), http.StatusInternalServerError)
		}
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}

func DeleteViajeHandler(w http.ResponseWriter, r *http.Request) {
	var viaje modelosBitacora.Viaje
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&viaje, parametros["id"])

	if viaje.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("viaje no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&viaje)
	w.WriteHeader(http.StatusOK)

}

func validarViaje(viaje modelosBitacora.Viaje) error {

	// fecha valida
	if viaje.FechaReservaViaje.IsZero() {
		return errors.New("la fecha del viaje esta vacia")
	}
	// vehiculo existente
	var vehiculo modelosBitacora.Vehiculo
	err := baseDeDatos.DB.Where("matricula = ?", viaje.Matricula).First(&vehiculo).Error
	if err != nil {
		return errors.New("el vehiculo no existe: " + viaje.Matricula)
	}
	// estado de vehiculo valido
	if vehiculo.Estado == "NO APTO PARA CIRCULAR" || vehiculo.Estado == "REPARACION" || vehiculo.Estado == "MANTENIMIENTO" {
		return errors.New("estado de vehiculo invalido para realizar un viaje")
	}
	// usuario existente
	var usuario modelosUsuario.Usuario
	err1 := baseDeDatos.DB.Where("username = ?", viaje.UsernameConductor).First(&usuario).Error
	if err1 != nil {
		return errors.New("el usuario no existe: " + viaje.UsernameConductor)
	}
	// vehiculo y conductor disponibles para la fecha del viaje
	var viajes []modelosBitacora.Viaje
	baseDeDatos.DB.Find(&viajes)
	for _, Viaje := range viajes {
		if Viaje.Matricula == viaje.Matricula && Viaje.FechaReservaViaje == viaje.FechaReservaViaje {
			return errors.New("el vehiculo ya esta reservado para esa fecha")
		}
		if Viaje.UsernameConductor == viaje.UsernameConductor && Viaje.FechaReservaViaje == viaje.FechaReservaViaje {
			return errors.New("el conductor ya tiene reservado un viaje para esa fecha")
		}
	}
	// paquete existente , con estado valido "sin asignar"
	var pesoTotalPaquetes float32
	var volumenTotalPaquetes float32
	for _, paqueteViaje := range viaje.Paquetes {
		paquete, err := dataPaquete.ObtenerPaquete(fmt.Sprint(paqueteViaje.ID))
		if err != nil {
			return fmt.Errorf("paquete con ID %v no encontrado: %w", paqueteViaje.ID, err)
		}

		if paquete.Estado != modelosPaquete.SIN_ASIGNAR {
			return fmt.Errorf("paquete con ID %v no está disponible para asignar (estado actual: %v)", paqueteViaje.ID, paquete.Estado)
		}

		pesoTotalPaquetes += paqueteViaje.Peso_kg
		volumenTotalPaquetes += paqueteViaje.Tamaño_mts_cubicos
	}
	// peso y volumen de paquetes aptos para el vehiculo
	if pesoTotalPaquetes > vehiculo.PesoAdmitido || volumenTotalPaquetes > vehiculo.VolumenAdmitidoMtsCubicos {
		return errors.New("los paquetes rebasan la capacidad admitida por el vehiculo")
	}

	return nil
}
