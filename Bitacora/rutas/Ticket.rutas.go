package rutas

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/modelosUsuarios"
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

	if err := validarTicket(ticket); err != nil {
		http.Error(w, "ticket inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	ticket.FechaCreacion = time.Now().Format("02-01-2006")
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

func validarTicket(ticket modelosBitacora.Ticket) error {

	switch ticket.Estado {
	case "NUEVO", "EN CURSO", "CERRADO", "RESUELTO":
	default:
		return errors.New("estado invalido")
	}
	switch ticket.Tipo {
	case "MANTENIMIENTO", "REPARACION":
	default:
		return errors.New("tipo invalido")
	}
	var usuario modelosUsuarios.Usuario
	err := baseDeDatos.DB.Where("username = ?", ticket.Username).First(&usuario).Error
	if err != nil {
		return errors.New("error al encontrar usuario " + err.Error())
	}

	return nil
}
