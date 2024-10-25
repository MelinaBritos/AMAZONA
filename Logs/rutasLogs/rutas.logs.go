package rutasLogs

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/modelosLogs"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

type Log = modelosLogs.Log


func GetAll(w http.ResponseWriter, r *http.Request) {

	logs, err := findAll()

	if StatusInternalServerError(w, err, "error en el servidor"){return}
	prettyJSON, err := json.MarshalIndent(logs, "", "  ")

	if StatusInternalServerError(w, err, "Error interno en el servidor al crear json") {return}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON)
}

func GetById(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	id := parametros["id"]

	log, err := findById(id)

	
	if StatusNotFound(w, err, "no se encontro al log"){return}
	if StatusInternalServerError(w, err, "error interno en el servidor"){return}
	

	prettyJSON, err := json.MarshalIndent(log, "", "  ")
	if StatusInternalServerError(w, err, "Error interno en el servidor al crear json") {return}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON)
}

func CreateLogHandler(w http.ResponseWriter, r *http.Request){

	var log Log
	err := json.NewDecoder(r.Body).Decode(&log)

	if BadRequestError(w, err, "JSON inválido") {return}

	err = CreateValidation(log)
	if BadRequestError(w, err, "los datos son invalidos"){return}

	err = CreateLog(log)
	if StatusInternalServerError(w, err, "error al crear el log"){return}

	prettyJSON, err := json.MarshalIndent(log, "", "  ")
	if StatusInternalServerError(w, err, "Error interno en el servidor al crear json") {return}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON)

}

func CreateMany(w http.ResponseWriter, r *http.Request){
	var logs []Log
	var err error

	err = json.NewDecoder(r.Body).Decode(&logs)
	if BadRequestError(w, err, "Error al decodificar las logs: ") {
		return
	}

	tx := baseDeDatos.DB.Begin()
	for _, log := range logs {
		err = CreateValidation(log)

		if StatusInternalServerError(w, err, "error al crear log"){
			tx.Rollback()
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Actualizacion exitosa!"))
}

func EditarLog(w http.ResponseWriter, r *http.Request){
	var log Log

	err := json.NewDecoder(r.Body).Decode(&log)

	if BadRequestError(w, err, "error al decodificar el body"){return}
	if BadRequestError(w, ValidateEdit(log), "Error en el body"){return}

	params := mux.Vars(r)
	id := params["id"]

	err = editById(id)

	if StatusNotFound(w, err, "log no encontrado"){return}
	if StatusInternalServerError(w, err, "actualizacion no realizada"){return}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Actualizacion exitosa!"))
}



func BorrarLog(w http.ResponseWriter, r *http.Request){
	
	params := mux.Vars(r)
	id := params["id"]

	err := DeleteById(id)

	if StatusNotFound(w, err, "no se encontro el log"){return}
	if StatusInternalServerError(w, err, "Hubo un problema de eliminacion") {return}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Eliminacion exitosa!"))
}

func BorrarLogs(w http.ResponseWriter, r *http.Request){
	
	var logs []Log
	var err error

	err = json.NewDecoder(r.Body).Decode(&logs)
	if BadRequestError(w, err, "Error al decodificar las logs: ") {
		return
	}

	tx := baseDeDatos.DB.Begin()

	for _, log := range logs {

		err = DeleteByIdU(log.ID)

		if StatusNotFound(w, err, "no se encontro el id"){
			tx.Rollback()
			return
		}

		if StatusInternalServerError(w, err, "error al eliminar el usuario"){
			tx.Rollback()
			return
		}


	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}


