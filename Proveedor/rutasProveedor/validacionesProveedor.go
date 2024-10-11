package rutasProveedor

import (
	"errors"
	"regexp"
	"strings"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func validarProveedor(proveedor modelosProveedor.Proveedor) error {

	if err := validarId(proveedor); err != nil {
		return err
	}
	if err := validarNombre(proveedor.Nombre_empresa); err != nil {
		return err
	}
	if err := validarMail(proveedor.Mail); err != nil {
		return err
	}
	//validarTelefono(proveedor.Telefono)

	return nil
}

func validarId(proveedor modelosProveedor.Proveedor) error {
	if proveedor.Id_proveedor <= 0 {
		return errors.New("el ID no puede ser negativo")
	}

	if err := baseDeDatos.DB.Where("id_proveedor = ?", proveedor.Id_proveedor).First(&proveedor); err != nil {
		return errors.New("el ID ya se encuentra en uso por otro proveedor")
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

// func validarTelefono(telefono string) error {
// 	// Expresión regular para teléfonos internacionales en formato E.164
// 	telefonoValido := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

// 	if !telefonoValido.MatchString(telefono) {
// 		return errors.New("número de teléfono invalido")
// 	}
// 	return nil
// }
