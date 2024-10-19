package rutasProveedor

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/validaciones"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetProveedoresHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los proveedores
	var proveedores []modelosProveedor.Proveedor
	baseDeDatos.DB.Find(&proveedores)
	json.NewEncoder(w).Encode(&proveedores)
	w.Header().Set("Content-Type", "application/json")

}

func GetProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo proveedor
	var proveedor modelosProveedor.Proveedor
	params := mux.Vars(r)
	idProveedor := params["id_proveedor"]

	baseDeDatos.DB.Where("id_proveedor = ?", idProveedor).First(&proveedor)

	if proveedor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El proveedor no existe"))
		return
	}
	json.NewEncoder(w).Encode(&proveedor)
}

func PostProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo proveedor
	var proveedor modelosProveedor.Proveedor

	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, "Error al decodificar el proveedor: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validaciones.ValidarProveedor(proveedor); err != nil {
		http.Error(w, "Datos del proveedor invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	if err := tx.Create(&proveedor); err.Error != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el proveedor: "+err.Error.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&proveedor)
}

func PutProveedorHandler(w http.ResponseWriter, r *http.Request) {

	// Obtener el ID del catálogo desde los parámetros de la URL
	var proveedorInput modelosProveedor.Proveedor
	if err := json.NewDecoder(r.Body).Decode(&proveedorInput); err != nil {
		http.Error(w, "Error al decodificar el catálogo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Asegurarse de que el ID esté presente en el cuerpo del request
	if proveedorInput.ID == 0 {
		http.Error(w, "ID del catálogo es requerido", http.StatusBadRequest)
		return
	}

	// Buscar el catálogo en la base de datos por el ID
	var proveedor modelosProveedor.Proveedor
	if err := baseDeDatos.DB.First(&proveedor, "id = ?", proveedorInput.ID).Error; err != nil {
		http.Error(w, "Catálogo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := validaciones.ValidarProveedor(proveedorInput); err != nil {
		http.Error(w, "Proveedor inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&proveedor, "id = ?", proveedorInput.ID).Error; err != nil {
		http.Error(w, "Proveedor no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := tx.Model(&proveedor).Updates(proveedorInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar el proveedor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.Write([]byte("Proveedor actualizado"))
	w.WriteHeader(http.StatusOK)
}

func DeleteProveedorHandler(w http.ResponseWriter, r *http.Request) {
	var proveedor modelosProveedor.Proveedor
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&proveedor, parametros["id"])

	if proveedor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Proveedor no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&proveedor)
	w.Write([]byte("Proveedor borrado"))
	w.WriteHeader(http.StatusOK)

}
