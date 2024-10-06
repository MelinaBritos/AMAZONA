package Proveedor

import (
	"net/http"

	//"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/endpo"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	//proveedore "github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos/Proveedor/Routes"
	"github.com/gorilla/mux"
)

func Iniciar() {
	baseDeDatos.Conexiondb()

	//baseDedatos.DB.AutoMigrate(modelos.Vehiculo{})

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ola mundo"))
	})

	//http.ListenAndServe(":3000", r)
}
