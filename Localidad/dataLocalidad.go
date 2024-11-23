package dataLocalidad

import (
	"fmt"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Bitacora/modelosBitacora"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/modelosLocalidad"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
	"gorm.io/gorm"
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

func obtenerLocalidadPorNombre(nombre_localidad string) (modelosLocalidad.Localidad, error) {
	var localidad modelosLocalidad.Localidad
	if err := baseDeDatos.DB.Where("nombre_localidad = ?", nombre_localidad).First(&localidad).Error; err != nil {
		return localidad, err
	}
	return localidad, nil
}

func CrearLocalidades(Localidades []modelosLocalidad.Localidad) error {
	tx := baseDeDatos.DB.Begin()

	if err := tx.Model(&modelosLocalidad.Localidad{}).Create(&Localidades).Error; err != nil {
		tx.Rollback()
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

func ObtenerZonas() []modelosLocalidad.Zona {

	return modelosLocalidad.ObtenerZonasValidas()
}

func ObtenerPrecioLocalidad(localidad string) float32 {

	loc, err := obtenerLocalidadPorNombre(localidad)
	if err != nil {
		return 0
	}

	return loc.Costo_localidad
}

func CargarIngreso(tx *gorm.DB, id_viaje int, id_paquete uint, ingreso float32) error {
	var ingresoViaje modelosBitacora.IngresosViaje
	ingresoViaje.IDViaje = id_viaje
	ingresoViaje.IDPaquete = int(id_paquete)
	ingresoViaje.Fecha = time.Now()
	ingresoViaje.Ingreso = ingreso

	if err := tx.Model(&modelosBitacora.IngresosViaje{}).Create(&ingresoViaje).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
