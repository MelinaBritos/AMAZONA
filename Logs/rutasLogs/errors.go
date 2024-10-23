package rutasLogs

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)


func BadRequestError(w http.ResponseWriter, err error, message string) bool {

	if err != nil {
		http.Error(w, message + ":" + err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

func StatusInternalServerError(w http.ResponseWriter, err error, message string)  bool{
	if err != nil {
		http.Error(w, message + ":" + err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func StatusNotFound(w http.ResponseWriter, err error, message string) bool {
	
	if errors.Is(err, gorm.ErrRecordNotFound){
		http.Error(w, message +  ":" + err.Error(), http.StatusNotFound)
		return true
	}
	return false
}