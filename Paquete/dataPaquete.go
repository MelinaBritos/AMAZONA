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

	for paquete := range paquetes {
		if err := tx.Create(paquete).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
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
