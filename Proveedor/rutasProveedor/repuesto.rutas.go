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

	params := mux.Vars(r)
	idRepuesto := params["id_repuesto"]

	// Buscar el repuesto en la base de datos
	var repuesto modelosProveedor.Repuesto
	if err := baseDeDatos.DB.Where("id_repuesto = ?", idRepuesto).First(&repuesto).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El repuesto no existe"))
		return
	}

	// Decodificar el cuerpo de la solicitud con los nuevos datos del repuesto
	var repuestoActualizado modelosProveedor.Repuesto
	if err := json.NewDecoder(r.Body).Decode(&repuestoActualizado); err != nil {
		http.Error(w, "Error al decodificar el repuesto: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validar el repuesto actualizado
	if err := validaciones.ValidarRepuesto(repuestoActualizado); err != nil {
		http.Error(w, "Datos del repuesto inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar los campos que se deseen cambiar en el repuesto original
	repuesto.Nombre = repuestoActualizado.Nombre
	repuesto.Stock = repuestoActualizado.Stock
	repuesto.Stock_minimo = repuestoActualizado.Stock_minimo
	repuesto.Cantidad_a_comprar = repuestoActualizado.Cantidad_a_comprar
	repuesto.Costo = repuestoActualizado.Costo
	repuesto.Descripcion = repuestoActualizado.Descripcion

	// Guardar los cambios en la base de datos
	if err := baseDeDatos.DB.Save(&repuesto).Error; err != nil {
		http.Error(w, "Error al actualizar el repuesto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar respuesta con el repuesto actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&repuesto)
}
