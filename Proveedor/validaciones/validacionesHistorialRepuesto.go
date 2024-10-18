package validaciones

import (
	"errors"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
)

func ValidarHistorialRepuesto(historialRepuesto modelosProveedor.HistorialRepuesto) error {

	if err := validarIdHistorialRepuesto(historialRepuesto); err != nil {
		return err
	}

	if err := validarNombre(historialRepuesto.Nombre); err != nil {
		return err
	}

	return nil

}

func validarIdHistorialRepuesto(historialRepuesto modelosProveedor.HistorialRepuesto) error {

	return errors.New("necesito mas directrices xd")

}
