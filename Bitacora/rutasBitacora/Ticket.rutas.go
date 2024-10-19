package rutasBitacora

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var tickets []modelosBitacora.Ticket

	baseDeDatos.DB.Find(&tickets)
	json.NewEncoder(w).Encode(&tickets)
}

func GetTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket modelosBitacora.Ticket
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&ticket, parametros["id"])

	if ticket.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&ticket)

}

func PostTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket modelosBitacora.Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, "Error al decodificar el ticket: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validarTicketNuevo(ticket); err != nil {
		http.Error(w, "ticket inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	ticket.FechaCreacion = time.Now().Format("02-01-2006")
	tx.Save(ticket)

	ticketCreado := tx.Create(&ticket)

	err := ticketCreado.Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//vehiculo cambia de estado
	var vehiculo modelosBitacora.Vehiculo
	err1 := tx.Where("matricula = ?", ticket.Matricula).First(&vehiculo).Error
	if err1 != nil {
		tx.Rollback()
		http.Error(w, "Error al encontrar vehiculo: "+err1.Error(), http.StatusBadRequest)
	}
	vehiculo.Estado = ticket.Tipo
	tx.Save(&vehiculo)

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func PutTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket modelosBitacora.Ticket
	var ticketInput modelosBitacora.Ticket
	// fecha finalizacion, costo total, desc reparacion, repuestos utilizados

	if err := json.NewDecoder(r.Body).Decode(&ticketInput); err != nil {
		http.Error(w, "Error al decodificar el ticket: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&ticket, "id = ?", ticketInput.ID).Error; err != nil {
		http.Error(w, "Ticket no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	ticketInput.FechaFinalizacion = time.Now().Format("02-01-2006")
	if len(ticketInput.DescripcionReparacion) > 100 || len(ticketInput.DescripcionReparacion) < 1 {
		tx.Rollback()
		http.Error(w, "La longitud de la descripcion de la reparacion es invalida", http.StatusInternalServerError)
		return
	}
	tx.Save(ticketInput)
	calcularCostoTotal(ticketInput)

	if err := tx.Model(&ticket).Updates(ticketInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al cerrar el Ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}

func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket modelosBitacora.Ticket
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&ticket, parametros["id"])

	if ticket.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&ticket)
	w.WriteHeader(http.StatusOK)

}

func calcularCostoTotal(ticket modelosBitacora.Ticket) {
	//compras automaticas con umblas minimo
	//historial de compras de repuestos
}

func validarTicketNuevo(ticket modelosBitacora.Ticket) error {

	switch ticket.Estado {
	case "EN CURSO", "SIN SOLUCION", "RESUELTO":
	default:
		return errors.New("estado invalido")
	}
	switch ticket.Tipo {
	case "MANTENIMIENTO", "REPARACION":
	default:
		return errors.New("tipo invalido")
	}
	if len(ticket.MotivoIngreso) > 100 || len(ticket.MotivoIngreso) < 1 {
		return errors.New("la longitud del motivo de ingreso es invalida")
	}

	return nil
}
