package dataPaquete

import (
	"fmt"
	"strconv"
	"time"

	dataLocalidad "github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"gorm.io/gorm"
)

func ObtenerPaquetes() []modelosPaquete.Paquete {

	var paquetes []modelosPaquete.Paquete
	baseDeDatos.DB.Find(&paquetes)
	return paquetes
}

func ObtenerPaquete(id_paquete string) (modelosPaquete.Paquete, error) {

	id_paquete_uint, err := strconv.ParseUint(id_paquete, 10, 32)
	if err != nil {
		return modelosPaquete.Paquete{}, err
	}

	return obtenerPaquete(uint(id_paquete_uint))
}

func obtenerPaquete(id_paquete uint) (modelosPaquete.Paquete, error) {

	var paquete modelosPaquete.Paquete
	if err := baseDeDatos.DB.Where("id = ?", id_paquete).First(&paquete).Error; err != nil {
		return paquete, err
	}
	return paquete, nil
}

func CrearPaquetes(paquetes []modelosPaquete.Paquete) error {
	tx := baseDeDatos.DB.Begin()

	if err := tx.Model(&modelosPaquete.Paquete{}).Create(&paquetes).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, paquete := range paquetes {
		if err := agregarHistorialPaquete(tx, paquete.ID, modelosPaquete.SIN_ASIGNAR); err != nil {
			tx.Rollback()
			return fmt.Errorf("error al actualizar el historial del paquete: %w", err)
		}
		precio := dataLocalidad.ObtenerPrecioLocalidad(paquete.Localidad)
		if err := dataLocalidad.CargarIngreso(tx, paquete.Id_viaje, paquete.ID, precio); err != nil {
			tx.Rollback()
			return fmt.Errorf("error al cargar un nuevo ingreso de dinero: %w", err)
		}
	}

	return tx.Commit().Error
}

func ActualizarPaquetes(paquetesInput []*modelosPaquete.Paquete) error {

	tx := baseDeDatos.DB.Begin()

	for _, paqueteInput := range paquetesInput {
		paqueteExistente, err := obtenerPaquete(paqueteInput.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("paquete no encontrado: %w", err)
		}

		if err := tx.Model(&paqueteExistente).Updates(paqueteInput).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error al actualizar el paquete: %w", err)
		}
	}

	return tx.Commit().Error
}

func BorrarPaquete(id_paquete string) error {

	paquete, err := ObtenerPaquete(id_paquete)
	if err != nil {
		return fmt.Errorf("error al actualizar el paquete: %w", err)
	}

	if err := baseDeDatos.DB.Unscoped().Delete(&paquete).Error; err != nil {
		return fmt.Errorf("error al borrar el paquete: %w", err)
	}

	return nil
}

func ActualizarEstadoPaquete(tx *gorm.DB, paquete *modelosPaquete.Paquete, estado modelosPaquete.Estado) error {

	paquete.Estado = estado

	if err := tx.Save(&paquete).Error; err != nil {
		return fmt.Errorf("error al actualizar el estado del paquete: %w", err)
	}

	if err := agregarHistorialPaquete(tx, paquete.ID, estado); err != nil {
		return fmt.Errorf("error al actualizar el historial del paquete: %w", err)
	}

	return tx.Error
}

func ObtenerPaquetesDeConductor(id_conductor string) ([]modelosPaquete.Paquete, error) {

	var paquetes []modelosPaquete.Paquete

	if err := baseDeDatos.DB.Where("id_conductor = ?", id_conductor).Find(&paquetes).Error; err != nil {
		return nil, fmt.Errorf("error al obtener los paquetes del conductor: %w", err)
	}

	return paquetes, nil
}

func ObtenerPaquetesSinAsignar() []modelosPaquete.Paquete {

	var paquetesSinAsignar []modelosPaquete.Paquete
	baseDeDatos.DB.Where("estado = ? OR estado = ?", modelosPaquete.SIN_ASIGNAR, modelosPaquete.NO_ENTREGADO).Find(&paquetesSinAsignar)
	return paquetesSinAsignar
}

func ObtenerHistorialPaquete(id_paquete string) ([]modelosPaquete.HistorialPaquete, error) {

	var historialPaquetes []modelosPaquete.HistorialPaquete

	if err := baseDeDatos.DB.Where("id_paquete = ?", id_paquete).Find(&historialPaquetes).Error; err != nil {
		return nil, err
	}

	return historialPaquetes, nil
}

func agregarHistorialPaquete(tx *gorm.DB, id_paquete uint, estado modelosPaquete.Estado) error {

	var historialPaquete modelosPaquete.HistorialPaquete
	historialPaquete.Id_paquete = id_paquete
	historialPaquete.Estado = estado
	historialPaquete.Fecha = time.Now()

	if err := tx.Model(&modelosPaquete.HistorialPaquete{}).Create(&historialPaquete).Error; err != nil {
		return err
	}

	return tx.Error
}

func ObtenerPaquetesPorViaje(id_viaje uint) ([]modelosPaquete.Paquete, error) {

	var paquetes []modelosPaquete.Paquete
	if err := baseDeDatos.DB.Where("id = ?", id_viaje).Find(&paquetes).Error; err != nil {
		return paquetes, err
	}
	return paquetes, nil
}

func AsignarViajeAPaquete(tx *gorm.DB, id_viaje uint, paquete *modelosPaquete.Paquete) error {

	paquete.Id_viaje = int(id_viaje)

	if err := tx.Save(&paquete).Error; err != nil {
		return fmt.Errorf("error al actualizar el viaje del paquete: %w", err)
	}

	return tx.Error
}
