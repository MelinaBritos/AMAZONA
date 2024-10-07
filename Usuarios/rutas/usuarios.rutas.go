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

func EditarContraseña(w http.ResponseWriter, r *http.Request)  {

	var usuario Usuario
	var datos struct {
		Clave string `json:"password"`
	}

	params := mux.Vars(r)
	username := params["username"]

	err := json.NewDecoder(r.Body).Decode(&datos)
    if err != nil {
        http.Error(w, "Debe suministrar la contraseña nueva", http.StatusBadRequest)
        return
    }
	
	
	err = baseDeDatos.DB.Model(&usuario).Where("username = ?", username).Update("clave", datos.Clave).Error
	
	if err != nil {
		http.Error(w, "Error al actualizar la contraseña", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario actualizado :)"))
}

func CrearUsuario(w http.ResponseWriter, r *http.Request)  {
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

	errors := verificarAtributos(usuario.Username, usuario.Clave, usuario.Dni, usuario.Nombre, usuario.Apellido)
	
	if len(errors) != 0 {
		http.Error(w, errors[0].Error(), http.StatusInternalServerError)
		return
	}

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