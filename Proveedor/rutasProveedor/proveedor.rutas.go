package rutasProveedor

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola mundo"))
}

func GetProveedoresHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los proveedores
	w.Write([]byte("ola mundo proveedores"))
}

func GetProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo proveedor
	w.Write([]byte("ola mundo proveedor"))
}

func PostProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo proveedor
	//w.Write([]byte("ola mundo post proveedor"))
	var proveedor modelosProveedor.Proveedor

	if err := json.NewDecoder(r.Body).Decode(&proveedor); err != nil {
		http.Error(w, "Error al decodificar el proveedor: "+err.Error(), http.StatusBadRequest)
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

	//aca va la logica para modificar los datos de un proveedor
	w.Write([]byte("ola mundo put proveedor"))
}
