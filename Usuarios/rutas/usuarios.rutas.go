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

func GetUsuariosByRolHandler(w http.ResponseWriter, r *http.Request)  {

	var usuarios []Usuario
	parametros := mux.Vars(r)
	rol := parametros["rol"]

	err := baseDeDatos.DB.Where("rol = ?", rol).Find(&usuarios).Error

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

func EditarUsuario(w http.ResponseWriter, r *http.Request){
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	params := mux.Vars(r)
	username := params["username"]

	if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    if NoExisteNingunCampo(usuario) {
        http.Error(w, "Debe proporcionar al menos un dato para actualizar", http.StatusBadRequest)
        return
    }

	errors := VerificarCamposExistentes(usuario);

	if len(errors) != 0 {
		http.Error(w, "Algun campo es invalido", http.StatusBadRequest)
        return
	}

	err = baseDeDatos.DB.Model(&usuario).Where("username = ?", username).Updates(usuario).Error

	if err != nil {
		http.Error(w, "Hubo un problema de actualizacion", http.StatusBadRequest)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Actualizacion exitosa!"))

}

func CrearUsuario(w http.ResponseWriter, r *http.Request)  {

	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

	errors := verificarAtributos(usuario.Clave, usuario.Dni, usuario.Nombre, usuario.Apellido)
	
	if len(errors) != 0 {
		http.Error(w, errors[0].Error(), http.StatusInternalServerError)
		return
	}

	usuario = DefinirUsername(usuario)
	err = baseDeDatos.DB.Model(&usuario).Create(usuario).Error

	if err != nil {
		http.Error(w, errors[0].Error(), http.StatusInternalServerError)
		return
	}
	
	prettyJSON, err := json.MarshalIndent(usuario, "", "  ")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(prettyJSON))
}