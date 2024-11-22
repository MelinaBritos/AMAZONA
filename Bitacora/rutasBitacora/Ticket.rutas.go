package rutasBitacora

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var tickets []modelosBitacora.Ticket

	baseDeDatos.DB.Find(&tickets)

	if err := baseDeDatos.DB.Preload("Repuestos").Find(&tickets).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	baseDeDatos.DB.Model(&ticket).Association("Repuestos").Find(&ticket.Repuestos)
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

	ticket.FechaCreacion = time.Now()
	ticket.Estado = "EN CURSO"

	ticketCreado := tx.Create(&ticket)
	tx.Save(ticket)

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

	if err := json.NewDecoder(r.Body).Decode(&ticketInput); err != nil {
		http.Error(w, "Error al decodificar el ticket: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&ticket, "id = ?", ticketInput.ID).Error; err != nil {
		http.Error(w, "Ticket no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := validarTicketCerrado(ticketInput); err != nil {
		http.Error(w, "Ticket inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	ticketInput.FechaFinalizacion = time.Now()
	tx.Save(ticketInput)

	if err := tx.Model(&ticket).Omit("FechaCreacion").Updates(ticketInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al cerrar el Ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := restarStockYcompras(ticketInput.Repuestos); err != nil {
		tx.Rollback()
		http.Error(w, "Error al restar stock de repuestos: "+err.Error(), http.StatusBadRequest)
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

func restarStockYcompras(repuestos []modelosBitacora.RepuestoUtilizado) error {

	tx := baseDeDatos.DB.Begin()
	for _, respuestoUtilizado := range repuestos {
		var repuesto modelosProveedor.Repuesto
		err := tx.Where("ID = ?", respuestoUtilizado.IDRepuesto).First(&repuesto).Error
		if err != nil {
			return errors.New("el repuesto no existe")
		}
		repuesto.Stock -= respuestoUtilizado.Cantidad

		//compra
		if repuesto.Stock <= repuesto.Stock_minimo {
			repuesto.Stock += repuesto.Cantidad_a_comprar

			var historialNuevo modelosBitacora.HistorialCompras

			historialNuevo.RepuestoComprado = repuesto
			historialNuevo.Cantidad = repuesto.Cantidad_a_comprar
			historialNuevo.Costo = float32(repuesto.Cantidad_a_comprar) * repuesto.Costo
			historialNuevo.FechaCompra = time.Now()

			historialCreado := tx.Create(&historialNuevo)
			err := historialCreado.Error
			if err != nil {
				tx.Rollback()
				return errors.New("error al crear historial de compras")
			}
			tx.Save(historialCreado)
		}
		tx.Save(repuesto)

	}
	tx.Commit()
	return nil
}

func validarTicketNuevo(ticket modelosBitacora.Ticket) error {
	var vehiculo modelosBitacora.Vehiculo
	err := baseDeDatos.DB.Where("matricula = ?", ticket.Matricula).First(&vehiculo).Error
	if err != nil {
		return errors.New("el vehiculo no existe " + ticket.Matricula)
	}
	var usuario modelosUsuario.Usuario
	err1 := baseDeDatos.DB.Where("username = ?", ticket.Username).First(&usuario).Error
	if err1 != nil {
		return errors.New("el usuario no existe " + ticket.Username)
	}
	switch ticket.Tipo {
	case "MANTENIMIENTO", "REPARACION":
	default:
		return errors.New("tipo invalido")
	}
	if len(ticket.MotivoIngreso) > 100 || len(ticket.MotivoIngreso) <= 1 {
		return errors.New("la longitud del motivo de ingreso es invalida")
	}

	return nil
}

func validarTicketCerrado(ticketInput modelosBitacora.Ticket) error {
	switch ticketInput.Estado {
	case "SIN SOLUCION", "RESUELTO":
	default:
		return errors.New("estado invalido")
	}
	if len(ticketInput.DescripcionReparacion) > 100 || len(ticketInput.DescripcionReparacion) <= 1 {
		return errors.New("la longitud de la descripcion de la reparacion es invalida")
	}
	for _, respuestoUtilizado := range ticketInput.Repuestos {
		var repuesto modelosProveedor.Repuesto
		err := baseDeDatos.DB.Where("ID = ?", respuestoUtilizado.IDRepuesto).First(&repuesto).Error
		if err != nil {
			return errors.New("el repuesto no existe")
		}
		if respuestoUtilizado.IDTicket != ticketInput.ID {
			return errors.New("el ID del ticket es incorrecto")
		}
		if respuestoUtilizado.Cantidad > repuesto.Stock {
			return errors.New("no hay suficiente stock para el repuesto: " + repuesto.Nombre)
		}
	}

	return nil
}
