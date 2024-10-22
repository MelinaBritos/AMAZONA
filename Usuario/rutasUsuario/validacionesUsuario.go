package rutasUsuario

import (
	"errors"
	"regexp"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Usuario/modelosUsuario"
)

type COMPARATOR string

const (
	SOFT COMPARATOR = "SOFT"
	HARD COMPARATOR = "HARD"
)

func verificarAtributos(usuario Usuario, comparator COMPARATOR) []error {

	var errorList []error

	appendError := func(err error) {
		if err != nil {
			errorList = append(errorList, err)
		}
	}

	if comparator != SOFT {
		hardvalidation(appendError, usuario)
	} else {
		softvalidation(usuario, appendError)
	}

	return errorList
}

func hardvalidation(appendError func(err error), usuario Usuario) {
	appendError(verificarDni(usuario.Dni))
	appendError(verificarNombre(usuario.Nombre))
	appendError(verificarApellido(usuario.Apellido))
	appendError(verificarcontraseña(usuario.Clave))
}

func softvalidation(usuario Usuario, appendError func(err error)) {
	if usuario.Clave != "" {
		appendError(verificarcontraseña(usuario.Clave))
	}

	if usuario.Nombre != "" {
		appendError(verificarNombre(usuario.Nombre))
	}

	if usuario.Apellido != "" {
		appendError(verificarApellido(usuario.Apellido))
	}

	if usuario.Dni != "" {
		appendError(verificarDni(usuario.Dni))
	}
}

func verificarDni(dni string) error {

	if len(dni) != 8 {
		err := errors.New("el dni no tiene 8 caracteres")
		return err
	}

	return nil

}

func verificarNombre(nombre string) error {

	if !tieneSoloLetras(nombre) {
		return errors.New("el nombre no puede contener numeros ni caracteres especiales")
	}

	if len(nombre) < 3 {
		return errors.New("el nombre debe tener al menos 3 caracteres")
	}
	return nil
}

func verificarApellido(apellido string) error {

	if !tieneSoloLetras(apellido) {
		return errors.New("el apellido no puede contener numeros ni caracteres especiales")
	}

	if len(apellido) < 3 {
		return errors.New("el apellido debe tener al menos 3 caracteres")
	}

	return nil
}

func verificarcontraseña(clave string) error {
	if len(clave) < 3 {
		err := errors.New("el username debe tener al menos 3 caracteres")
		return err
	}
	return nil
}

func tieneSoloLetras(value string) bool {

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return regex.MatchString(value)
}

func DefinirUsername(usuario Usuario) Usuario {

	first_letter_name, first_letter_surname := defineFirstletter(usuario)

	usuario.Username = first_letter_name + first_letter_surname + usuario.Dni
	return usuario
}

func defineFirstletter(usuario Usuario) (string, string) {
	first_letter_name := string(usuario.Nombre[0])
	first_letter_surname := string(usuario.Apellido[0])
	return first_letter_name, first_letter_surname
}

func NoExisteNingunCampo(usuario Usuario) bool {
	return usuario.Clave == "" && usuario.Nombre == "" && usuario.Apellido == "" && usuario.Dni == "" && usuario.Rol == ""
}

func DefinirUsuarioSegunApellido(usuario modelosUsuario.Usuario, usuarioActual modelosUsuario.Usuario) Usuario {
	if usuario.Apellido != "" {

		if usuario.Dni == "" {
			usuario.Dni = usuarioActual.Dni
		}

	} else {
		usuario.Dni = usuarioActual.Dni
		usuario.Apellido = usuarioActual.Apellido
	}
	return usuario
}

func DefinirUsuarioSegunNombreVacio(usuario modelosUsuario.Usuario, usuarioActual modelosUsuario.Usuario) Usuario {
	
	usuario.Nombre = usuarioActual.Nombre
	if usuario.Apellido != "" {
		usuario.Dni = usuarioActual.Dni

	} else {
		usuario.Apellido = usuarioActual.Apellido
	}
	return usuario
}
