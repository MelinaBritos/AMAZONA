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

	if err := json.NewDecoder(r.Body).Decode(&historialRepuesto); err != nil {
		http.Error(w, "Error al decodificar el HistorialRepuesto: "+err.Error(), http.StatusBadRequest)
		return
	}

	//baseDeDatos.DB.Where("id_repuesto = ? AND id_catalogo = ? AND f_validez = ?", historialRepuesto.Id_repuesto, historialRepuesto.Id_catalogo, historialRepuesto.F_validez).First(&historialRepuesto)

	if err := baseDeDatos.DB.First(&historialRepuesto, "id_repuesto = ? AND id_catalogo = ? AND f_validez = ?", historialRepuesto.Id_repuesto, historialRepuesto.Id_catalogo, historialRepuesto.F_validez).Error; err != nil {
		http.Error(w, "Catálogo no encontrado: "+err.Error(), http.StatusNotFound)
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

	// Obtener el ID del catálogo desde los parámetros de la URL
	var historialRepuestoInput modelosProveedor.HistorialRepuesto
	if err := json.NewDecoder(r.Body).Decode(&historialRepuestoInput); err != nil {
		http.Error(w, "Error al decodificar el catálogo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Asegurarse de que el ID esté presente en el cuerpo del request
	if historialRepuestoInput.ID == 0 {
		http.Error(w, "ID del catálogo es requerido", http.StatusBadRequest)
		return
	}

	// Buscar el catálogo en la base de datos por el ID
	var historialRepuesto modelosProveedor.HistorialRepuesto
	if err := baseDeDatos.DB.First(&historialRepuesto, "id = ?", historialRepuestoInput.ID).Error; err != nil {
		http.Error(w, "Catálogo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := validaciones.ValidarHistorialRepuesto(historialRepuestoInput); err != nil {
		http.Error(w, "HistorialRepuesto inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&historialRepuesto, "id = ?", historialRepuestoInput.ID).Error; err != nil {
		http.Error(w, "HistorialRepuesto no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := tx.Model(&historialRepuesto).Updates(historialRepuestoInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar el historialRepuesto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.Write([]byte("HistorialRepuesto actualizado"))
	w.WriteHeader(http.StatusOK)
}

func DeleteHistorialRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	var historialRepuesto modelosProveedor.HistorialRepuesto
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&historialRepuesto, parametros["id"])

	if historialRepuesto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("HistorialRepuesto no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&historialRepuesto)
	w.Write([]byte("HistorialRepuesto borrado"))
	w.WriteHeader(http.StatusOK)

}
