package validaciones

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Localidad/modelosLocalidad"
)

func ValidarLocalidad(localidad modelosLocalidad.Localidad) error {

	if err := validarNombre(localidad.Nombre_localidad); err != nil {
		return err
	}

	if !validarZona(localidad.Zona_pertenencia) {
		return fmt.Errorf("la zona %s no es una zona valida", localidad.Zona_pertenencia)
	}

	if err := validarCosto(localidad.Costo_localidad); err != nil {
		return err
	}

	return nil

}

func validarNombre(nombreLocalidad string) error {

	if strings.TrimSpace(nombreLocalidad) == "" {
		return errors.New("nombre invalido")
	}

	return nil
}

func validarZona(zona modelosLocalidad.Zona) bool {

	for _, zonaValida := range modelosLocalidad.ObtenerZonasValidas() {
		if zona == zonaValida {
			return true
		}
	}
	return false
}

func validarCosto(costo float32) error {

	if costo <= 0 {
		return errors.New("el no puede ser negativo ni cero")
	}

	return nil
}
