package Proveedor

import (
	

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/rutas"
	"github.com/gorilla/mux"
)

func Iniciar() {

	r := mux.NewRouter()

	r.HandleFunc("/proveedor/{id_proveedor}", rutas.GetProveedorHandler).Methods("GET")
	r.HandleFunc("/proveedor", rutas.GetProveedoresHandler).Methods("GET")

	r.HandleFunc("/proveedor", rutas.PutProveedorHandler).Methods("PUT") //Modificar datos de algun usuario

	r.HandleFunc("/proveedor", rutas.PostProveedorHandler).Methods("POST") //crear un usuario

	//fmt.Println("listen users at port 3002")
	//http.ListenAndServe(":3002", r)
}
