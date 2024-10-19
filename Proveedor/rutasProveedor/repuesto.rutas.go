package rutasProveedor

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/validaciones"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetRepuestosHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los Repuestos
	var repuestos []modelosProveedor.Repuesto
	baseDeDatos.DB.Find(&repuestos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&repuestos)

}

func GetRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo repuesto
	var repuesto modelosProveedor.Repuesto
	params := mux.Vars(r)
	idRepuesto := params["id_repuesto"]

	baseDeDatos.DB.Where("id_repuesto = ?", idRepuesto).First(&repuesto)

	if repuesto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El repuesto no existe"))
		return
	}
	json.NewEncoder(w).Encode(&repuesto)
}

func PostRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo repuesto
	var repuesto modelosProveedor.Repuesto

	if err := json.NewDecoder(r.Body).Decode(&repuesto); err != nil {
		http.Error(w, "Error al decodificar el repuesto: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validaciones.ValidarRepuesto(repuesto); err != nil {
		http.Error(w, "Datos del repuesto invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	if err := tx.Create(&repuesto); err.Error != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el repuesto: "+err.Error.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&repuesto)
}

func PutRepuestoHandler(w http.ResponseWriter, r *http.Request) {

	// Obtener el ID del catálogo desde los parámetros de la URL
	var repuestoInput modelosProveedor.Repuesto
	if err := json.NewDecoder(r.Body).Decode(&repuestoInput); err != nil {
		http.Error(w, "Error al decodificar el catálogo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Asegurarse de que el ID esté presente en el cuerpo del request
	if repuestoInput.ID == 0 {
		http.Error(w, "ID del catálogo es requerido", http.StatusBadRequest)
		return
	}

	// Buscar el catálogo en la base de datos por el ID
	var repuesto modelosProveedor.Repuesto
	if err := baseDeDatos.DB.First(&repuesto, "id = ?", repuestoInput.ID).Error; err != nil {
		http.Error(w, "Catálogo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := validaciones.ValidarRepuesto(repuestoInput); err != nil {
		http.Error(w, "Repuesto inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&repuesto, "id = ?", repuestoInput.ID).Error; err != nil {
		http.Error(w, "Repuesto no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := tx.Model(&repuesto).Updates(repuestoInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar el repuesto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.Write([]byte("Repuesto actualizado"))
	w.WriteHeader(http.StatusOK)
}

func DeleteRepuestoHandler(w http.ResponseWriter, r *http.Request) {
	var repuesto modelosProveedor.Repuesto
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&repuesto, parametros["id"])

	if repuesto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Repuesto no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&repuesto)
	w.Write([]byte("Repuesto borrado"))
	w.WriteHeader(http.StatusOK)

}
