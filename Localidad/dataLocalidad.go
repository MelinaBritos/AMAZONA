package dataLocalidad

import (
	"fmt"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/modelosLocalidad"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func ObtenerLocalidades() []modelosLocalidad.Localidad {

	var Localidades []modelosLocalidad.Localidad
	baseDeDatos.DB.Find(&Localidades)
	return Localidades
}

func ObtenerLocalidad(id_localidad string) (modelosLocalidad.Localidad, error) {

	var localidad modelosLocalidad.Localidad
	if err := baseDeDatos.DB.Where("id = ?", id_localidad).First(&localidad).Error; err != nil {
		return localidad, err
	}
	return localidad, nil

}

func CrearLocalidades(Localidades []modelosLocalidad.Localidad) error {
	tx := baseDeDatos.DB.Begin()

	// Especificar el modelo para la tabla
	if err := tx.Model(&modelosLocalidad.Localidad{}).Create(&Localidades).Error; err != nil {
		tx.Rollback() // En caso de error, realizar rollback
		return err
	}

	return tx.Commit().Error
}

func ActualizarLocalidades(LocalidadesInput []*modelosLocalidad.Localidad) error {

	tx := baseDeDatos.DB.Begin()

	for _, localidadInput := range LocalidadesInput {
		localidadExistente, err := ObtenerLocalidad(localidadInput.GetIDAsString())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("localidad no encontrada: %w", err)
		}

		if err := tx.Model(&localidadExistente).Updates(localidadInput).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error al actualizar el localidad: %w", err)
		}
	}

	return tx.Commit().Error
}

func BorrarLocalidad(id_localidad string) error {

	localidad, err := ObtenerLocalidad(id_localidad)
	if err != nil {
		return fmt.Errorf("error al obtener la localidad: %w", err)
	}

	if err := baseDeDatos.DB.Unscoped().Delete(&localidad).Error; err != nil {
		return fmt.Errorf("error al borrar el localidad: %w", err)
	}

	return nil
}

func ObtenerLocalidadesPorZona(zona string) ([]modelosLocalidad.Localidad, error) {

	var localidades []modelosLocalidad.Localidad
	if err := baseDeDatos.DB.Where("zona_pertenencia = ?", zona).Find(&localidades).Error; err != nil {
		return nil, fmt.Errorf("error al obtener las localidades: %w", err)
	}

	return localidades, nil
}
