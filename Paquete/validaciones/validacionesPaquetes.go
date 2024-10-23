package validaciones

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
)

func ValidarPaquete(paquete modelosPaquete.Paquete) error {

	// if !existeViaje(paquete.Id_viaje) {
	// 	return fmt.Errorf("no existe el viaje %d", paquete.Id_viaje)
	// }

	// if !existeConductor(paquete.Id_conductor) {
	// 	return fmt.Errorf("no existe el conductor %d", paquete.Id_conductor)
	// }

	// if !existeMatricula(paquete.Matricula) {
	// 	return fmt.Errorf("no existe la matricula %s", paquete.Matricula)
	// }

	if !estadoValido(paquete.Estado) {
		return fmt.Errorf("el estado %s no es un estado valido", paquete.Estado)
	}

	if err := validarPeso(paquete.Peso_kg); err != nil {
		return err
	}

	if err := validarNombreCliente(paquete.Nombre_cliente); err != nil {
		return err
	}

	if err := validarTamaño(paquete.Tamaño_mts_cubicos); err != nil {
		return err
	}

	if err := validarLocalidad(paquete.Localidad); err != nil {
		return err
	}

	if err := validarDireccionEntrega(paquete.Dir_entrega); err != nil {
		return err
	}

	return nil

}

// func existeViaje(id_viaje int) bool {

// 	//existeViajeDeViaje (id_viaje)
// 	var viaje modelosViaje.Viaje
// 	baseDeDatos.DB.Where("id = ?", id_viaje).First(&viaje)

// 	return viaje.ID != 0
// }

// func existeConductor(id_conductor int) bool {

// 	var conductor modelosUsuario.Usuario
// 	baseDeDatos.DB.Where("id = ? AND rol = ?", id_conductor, modelosUsuario.CONDUCTOR).First(&conductor)

// 	return conductor.ID != 0
// }

// func existeMatricula(matricula string) bool {

// 	var vehiculo modelosBitacora.Vehiculo
// 	baseDeDatos.DB.Where("matricula = ?", matricula).First(&vehiculo)

// 	return vehiculo.ID != 0
// }

func estadoValido(estado modelosPaquete.Estado) bool {

	for _, estadoValido := range modelosPaquete.ObtenerEstadosValidos() {
		if estado == estadoValido {
			return true
		}
	}
	return false
}

func validarPeso(peso float32) error {

	if peso <= 0 {
		return errors.New("el no puede ser negativo ni cero")
	}

	return nil
}

func validarNombreCliente(nombre string) error {

	if strings.TrimSpace(nombre) == "" {
		return errors.New("nombre invalido")
	}

	return nil
}

func validarTamaño(tamaño float32) error {

	if tamaño <= 0 {
		return errors.New("el tamaño no puede ser negativo ni cero")
	}

	return nil
}

func validarLocalidad(localidad string) error {

	if strings.TrimSpace(localidad) == "" {
		return errors.New("localidad invalida")
	}

	return nil
}

func validarDireccionEntrega(dir_entrega string) error {

	if strings.TrimSpace(dir_entrega) == "" {
		return errors.New("direccion de entrega invalida")
	}

	return nil
}
