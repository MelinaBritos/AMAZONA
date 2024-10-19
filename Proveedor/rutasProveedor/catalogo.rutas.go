package rutasProveedor

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/validaciones"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetCatalogosHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los catalogos
	var catalogos []modelosProveedor.Catalogo
	baseDeDatos.DB.Find(&catalogos)
	json.NewEncoder(w).Encode(&catalogos)
	w.Header().Set("Content-Type", "application/json")

}

func GetCatalogoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo catalogo
	var catalogo modelosProveedor.Catalogo
	params := mux.Vars(r)
	idCatalogo := params["id"]

	baseDeDatos.DB.Where("id = ?", idCatalogo).First(&catalogo)

	if catalogo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El catalogo no existe"))
		return
	}
	json.NewEncoder(w).Encode(&catalogo)
}

func PostCatalogoHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo catalogo
	var catalogo modelosProveedor.Catalogo

	if err := json.NewDecoder(r.Body).Decode(&catalogo); err != nil {
		http.Error(w, "Error al decodificar el catalogo: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validaciones.ValidarCatalogo(catalogo); err != nil {
		http.Error(w, "Datos del catalogo invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	if err := tx.Create(&catalogo); err.Error != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el catalogo: "+err.Error.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&catalogo)
}

func PutCatalogoHandler(w http.ResponseWriter, r *http.Request) {

	// Obtener el ID del catálogo desde los parámetros de la URL
	var catalogoInput modelosProveedor.Catalogo
	if err := json.NewDecoder(r.Body).Decode(&catalogoInput); err != nil {
		http.Error(w, "Error al decodificar el catálogo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Asegurarse de que el ID esté presente en el cuerpo del request
	if catalogoInput.ID == 0 {
		http.Error(w, "ID del catálogo es requerido", http.StatusBadRequest)
		return
	}

	// Buscar el catálogo en la base de datos por el ID
	var catalogo modelosProveedor.Catalogo
	if err := baseDeDatos.DB.First(&catalogo, "id = ?", catalogoInput.ID).Error; err != nil {
		http.Error(w, "Catálogo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := validaciones.ValidarCatalogo(catalogoInput); err != nil {
		http.Error(w, "Catalogo inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&catalogo, "id = ?", catalogoInput.ID).Error; err != nil {
		http.Error(w, "Catalogo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := tx.Model(&catalogo).Updates(catalogoInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar el catalogo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.Write([]byte("Catalogo actualizado"))
	w.WriteHeader(http.StatusOK)
}

func DeleteCatalogoHandler(w http.ResponseWriter, r *http.Request) {
	var catalogo modelosProveedor.Catalogo
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&catalogo, parametros["id"])

	if catalogo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Catalogo no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&catalogo)
	w.Write([]byte("Catalogo borrado"))
	w.WriteHeader(http.StatusOK)

}
