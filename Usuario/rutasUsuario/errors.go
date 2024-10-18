package rutasUsuario

import (
	"net/http"
)


func BadRequestError(w http.ResponseWriter, err error, message string) bool {

	if err != nil {
		http.Error(w, message, http.StatusBadRequest)
		return true
	}
	return false
}

func StatusNotFoundError(w http.ResponseWriter, err error, message string) bool{
	if err != nil{
		http.Error(w, message, http.StatusNotFound)
		return true
	}
	return false

}

func StatusUnauthorizedError(w http.ResponseWriter, err error, message string)  bool{
	if err != nil{
		http.Error(w, message, http.StatusUnauthorized)
		return true
	}
	return false
}

func StatusInternalServerError(w http.ResponseWriter, err error, message string)  bool{
	if err != nil {
		http.Error(w, message, http.StatusInternalServerError)
		return true
	}
	return false
}

func StatusNotFound(w http.ResponseWriter, err error, message string) bool {
	if err != nil{
		http.Error(w, message, http.StatusNotFound)
		return true
	}
	return false
}