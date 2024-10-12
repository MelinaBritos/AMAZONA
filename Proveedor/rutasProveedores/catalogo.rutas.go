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
	idCatalogo := params["id_catalogo"]

	baseDeDatos.DB.Where("id_catalogo = ?", idCatalogo).First(&catalogo)

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

	//aca va la logica para modificar los datos de un catalogo
	w.Write([]byte("ola mundo put catalogo"))
}
