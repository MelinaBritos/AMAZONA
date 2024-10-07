package rutas

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola mundo"))
}

func GetProveedoresHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener los proveedores
	w.Write([]byte("ola mundo proveedores"))
}

func GetProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para obtener un solo proveedor
	w.Write([]byte("ola mundo proveedor"))
}

func PostProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para modificar los datos de un proveedor
	w.Write([]byte("ola mundo post proveedor"))
}

func PutProveedorHandler(w http.ResponseWriter, r *http.Request) {
	//aca va la logica para agregar un nuevo proveedor
	w.Write([]byte("ola mundo put proveedor"))
}
