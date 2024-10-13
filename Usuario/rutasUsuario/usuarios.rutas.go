package rutasUsuario

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Usuario = modelosUsuario.Usuario
type Credencial = modelosUsuario.Credencial

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

func GetUsuarioByIdHandler(w http.ResponseWriter, r *http.Request) {

	var usuario Usuario
	parametros := mux.Vars(r)
	username := parametros["username"]

	err := baseDeDatos.DB.Where("username = ?", username).First(&usuario).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		prettyJSON, err := json.MarshalIndent(usuario, "", "  ")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(prettyJSON)
	}
}

func GetUsuariosByRolHandler(w http.ResponseWriter, r *http.Request) {

	var usuarios []Usuario
	parametros := mux.Vars(r)
	rol := parametros["rol"]

	err := baseDeDatos.DB.Where("rol = ?", rol).Find(&usuarios).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else if len(usuarios) == 0 {
		w.WriteHeader(http.StatusNoContent)
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

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
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
	}

	errors := VerificarCamposExistentes(usuario)

	if len(errors) != 0 {
		http.Error(w, "Algun campo es invalido", http.StatusBadRequest)
		return
	}

	err = baseDeDatos.DB.Model(&usuario).Where("username = ?", username).Updates(&usuario).Error

	if err != nil {
		http.Error(w, "Hubo un problema de actualizacion", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Actualizacion exitosa!"))

}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {

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
	usuario.Clave, err = Encriptar(usuario.Clave)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = baseDeDatos.DB.Model(&usuario).Create(&usuario).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func CrearUsuarios(w http.ResponseWriter, r *http.Request){

	var usuarios []Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuarios); err != nil {
		http.Error(w, "Error al decodificar los usuarios: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, usuario := range usuarios {
		if err := verificarAtributos(usuario.Clave, usuario.Dni, usuario.Nombre,usuario.Apellido); err != nil {
			http.Error(w, "usuario inválido", http.StatusBadRequest)
			return
		}
	}

	tx := baseDeDatos.DB.Begin()
	for _, usuario := range usuarios {

		usuarioCreado := tx.Create(&usuario)

		err := usuarioCreado.Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al crear los usuarios: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario

	params := mux.Vars(r)
	username := params["username"]

	err := baseDeDatos.DB.Where("username = ?", username).Unscoped().Delete(&usuario).Error

	if err != nil {
		http.Error(w, "Hubo un problema de eliminacion", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Eliminacion exitosa!"))

}

func Loguearse(w http.ResponseWriter, r *http.Request) {

	var usuario Usuario
	var credencial Credencial

	err := json.NewDecoder(r.Body).Decode(&credencial)

	if err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	err = baseDeDatos.DB.Model(&usuario).Where("username = ?", credencial.Username).First(&usuario).Error

	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusNotFound)
		return
	}

	err = Equals(credencial.Password, usuario.Clave)

	if err != nil {
		http.Error(w, "la contraseña es incorrecta", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	prettyJSON, err := json.MarshalIndent(usuario, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(prettyJSON)

}
