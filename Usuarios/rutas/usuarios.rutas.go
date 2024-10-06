package rutas

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuarios/modelos"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

type Usuario = modelos.Usuario

func GetUsuariosHandler(w http.ResponseWriter, r *http.Request) {

	var usuarios []Usuario
	err := baseDeDatos.DB.Find(&usuarios).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		prettyJSON, err := json.MarshalIndent(usuarios, "", "  ")
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(prettyJSON)
	}
}

func GetUsuariosByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	var usuarios []Usuario
	parametros := mux.Vars(r)
	username := parametros["username"]

	err := baseDeDatos.DB.Where("username = ?", username).Find(&usuarios).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else if len(usuarios) == 0{
		w.WriteHeader(http.StatusNoContent)
	} else{

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		prettyJSON, err := json.MarshalIndent(usuarios, "", "  ")
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(prettyJSON)
	}
}

func EditarUsuario(w http.ResponseWriter, r *http.Request)  {
	
}