package rutasLogs

import (
	"net/http"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/modelosLogs"
)

type Log = modelosLogs.Log

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func GetById(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusNotImplemented)
}

func GetByUsername(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
}

func EditarLog(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
}

func BorrarLog(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
}