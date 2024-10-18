package validaciones

import (
	"errors"
	"regexp"
	"strings"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
)

func ValidarProveedor(proveedor modelosProveedor.Proveedor) error {

	if err := validarNombre(proveedor.Nombre_empresa); err != nil {
		return err
	}

	if err := validarMail(proveedor.Mail); err != nil {
		return err
	}

	if err := validarTelefono(proveedor.Telefono); err != nil {
		return err
	}

	return nil
}

func validarNombre(nombre string) error {

	if strings.TrimSpace(nombre) == "" {
		return errors.New("nombre invalido")
	}

	return nil
}

func validarMail(mail string) error {

	// Expresión regular para validar un correo electrónico
	mailValido := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !mailValido.MatchString(mail) {
		return errors.New("correo electrónico inválido")
	}

	return nil
}

func validarTelefono(telefono string) error {

	// Expresión regular para teléfonos internacionales en formato E.164
	telefonoValido := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

	if !telefonoValido.MatchString(telefono) {
		return errors.New("número de teléfono invalido")
	}
	return nil
}
