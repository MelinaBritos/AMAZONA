package rutasBitacora

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"github.com/gorilla/mux"
)

func GetCostosHandler(w http.ResponseWriter, r *http.Request) {
	var Costos []modelosBitacora.CostosViaje

	baseDeDatos.DB.Find(&Costos)
	json.NewEncoder(w).Encode(&Costos)
}

func GetCostoHandler(w http.ResponseWriter, r *http.Request) {
	var Costo modelosBitacora.CostosViaje
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&Costo, parametros["id"])

	if Costo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Costo no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&Costo)

}

func PostCostoHandler(w http.ResponseWriter, r *http.Request) {
	var Costo modelosBitacora.CostosViaje

	if err := json.NewDecoder(r.Body).Decode(&Costo); err != nil {
		http.Error(w, "Error al decodificar el costo: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validarCostoEstimado(Costo); err != nil {
		http.Error(w, "Costo inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()

	CostoCreado := tx.Create(&Costo)

	err := CostoCreado.Error
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error al crear el Costo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func PutCostoHandler(w http.ResponseWriter, r *http.Request) {
	var Costo modelosBitacora.CostosViaje
	var CostoInput modelosBitacora.CostosViaje

	if err := json.NewDecoder(r.Body).Decode(&CostoInput); err != nil {
		http.Error(w, "Error al decodificar el costo: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validarCostoFinal(CostoInput); err != nil {
		http.Error(w, "Costo inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx := baseDeDatos.DB.Begin()
	if err := tx.First(&Costo, "ID = ?", CostoInput.ID).Error; err != nil {
		http.Error(w, "Costo no encontrado: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := tx.Model(&Costo).Updates(CostoInput).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar el costo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func DeleteCostoHandler(w http.ResponseWriter, r *http.Request) {
	var Costo modelosBitacora.CostosViaje
	parametros := mux.Vars(r)

	baseDeDatos.DB.First(&Costo, parametros["id"])

	if Costo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Costos no encontrado"))
		return
	}

	baseDeDatos.DB.Unscoped().Delete(&Costo)
	w.WriteHeader(http.StatusOK)

}

func validarCostoEstimado(costo modelosBitacora.CostosViaje) error {

	var viaje modelosBitacora.Viaje
	err := baseDeDatos.DB.Where("ID = ?", costo.IDViaje).First(&viaje).Error
	if err != nil {
		return errors.New("el viaje no existe")
	}
	if costo.KilometrosEstimados < 0 || costo.KilometrosEstimados > 100000000 {
		return errors.New("los kilometros estimados no son validos")
	}
	if costo.CostoCombustibleEstimado < 0 {
		return errors.New("el costo del combustible no es valido")
	}
	return nil

}

func validarCostoFinal(costo modelosBitacora.CostosViaje) error {

	if costo.KilometrosRecorridosFinal < 0 || costo.KilometrosRecorridosFinal > 100000000 {
		return errors.New("los kilometros recorridos no son validos")
	}
	if costo.CostoCombustibleFinal < 0 {
		return errors.New("el costo del combustible no es valido")
	}
	return nil
}
