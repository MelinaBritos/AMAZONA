package dataPaquete

import (
	"fmt"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func ObtenerPaquetes() []modelosPaquete.Paquete {

	var paquetes []modelosPaquete.Paquete
	baseDeDatos.DB.Find(&paquetes)
	return paquetes
}

func ObtenerPaquete(id_paquete string) (modelosPaquete.Paquete, error) {

	var paquete modelosPaquete.Paquete
	if err := baseDeDatos.DB.Where("id = ?", id_paquete).First(&paquete).Error; err != nil {
		return paquete, err
	}
	return paquete, nil

}

func CrearPaquetes(paquetes []modelosPaquete.Paquete) error {
	tx := baseDeDatos.DB.Begin()

	// Especificar el modelo para la tabla
	if err := tx.Model(&modelosPaquete.Paquete{}).Create(&paquetes).Error; err != nil {
		tx.Rollback() // En caso de error, realizar rollback
		return err
	}

	return tx.Commit().Error
}

func ActualizarPaquetes(paquetesInput []*modelosPaquete.Paquete) error {

	tx := baseDeDatos.DB.Begin()

	for _, paqueteInput := range paquetesInput {
		paqueteExistente, err := ObtenerPaquete(paqueteInput.GetIDAsString())
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

func ActualizarEstadoPaquete(id_paquete string, estado string) error {

	estadoNuevo, err := modelosPaquete.ParseEstado(estado)
	if err != nil {
		return fmt.Errorf("error al parsear el estado: %w", err)
	}

	paquete, err := ObtenerPaquete(id_paquete)
	if err != nil {
		return err
	}

	tx := baseDeDatos.DB.Begin()
	paquete.Estado = estadoNuevo

	if err := tx.Save(&paquete).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error al actualizar el estado del paquete: %w", err)
	}

	return tx.Commit().Error
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
	baseDeDatos.DB.Where("estado = ?", modelosPaquete.SIN_ASIGNAR).Find(&paquetesSinAsignar)
	return paquetesSinAsignar

}
