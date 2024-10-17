package validaciones

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func ValidarCatalogo(catalogo modelosProveedor.Catalogo) error {

	if err := validarIdCatalogo(catalogo.Id_catalogo); err != nil {
		return err
	}

	if !existeProveedor(catalogo.Id_proveedor) {
		return fmt.Errorf("no existe el catalogo %d", catalogo.Id_catalogo)
	}

	if err := validarMesVigencia(catalogo.Mes_vigencia); err != nil {
		return err
	}

	return nil
}

func validarIdCatalogo(id_catalogo int) error {

	if id_catalogo <= 0 {
		return errors.New("el ID no puede ser negativo")
	}

	var catalogo modelosProveedor.Catalogo
	catalogoResultado := baseDeDatos.DB.Where("id_catalogo = ?", id_catalogo).First(&catalogo)
	if catalogoResultado.RowsAffected > 0 { //esta funcion calcula la cantidad de ocurrencias de la consulta
		return fmt.Errorf("el catalogo con id %d ya existe", id_catalogo)
	}

	return nil
}

func existeProveedor(id_proveedor int) bool {

	var proveedor modelosProveedor.Proveedor
	baseDeDatos.DB.Where("id_proveedor = ?", id_proveedor).First(&proveedor)

	return proveedor.Id_proveedor != 0
}

func validarMesVigencia(mes_vigencia string) error {

	if len(mes_vigencia) != 7 {
		return errors.New("longitud invalida para la fecha. El formato es mm/aaaa")
	}

	fechaActual := time.Now()

	//se valida que se ingrese un valor numerico para el mes
	mesStr := mes_vigencia[:2]
	mesInt, err := strconv.Atoi(mesStr)
	if err != nil {
		return errors.New("error al ingresar el mes. intente nuevamente")
	}

	if mesInt <= 0 || mesInt >= 12 {
		return errors.New("valor invalido del mes. Los valores validos son 1-12")
	}

	if mesActual := fechaActual.Month(); mesInt != int(mesActual) {
		return fmt.Errorf("el mes no puede ser distinto al mes corriente (%d)", mesActual)
	}

	//se valida que se ingrese un valor numerico para el a;o
	anioStr := mes_vigencia[3:]
	anioInt, err := strconv.Atoi(anioStr)
	if err != nil {
		return errors.New("error al ingresar el año. Intente nuevamente")
	}

	// Validar que el año no sea mayor al año actual
	if anioActual := fechaActual.Year(); anioInt > anioActual {
		return fmt.Errorf("valor invalido del año. No puede ser mayor al año actual %d", anioActual)
	}

	//se valida el separador
	if separadorStr := string(mes_vigencia[2]); separadorStr != "/" {
		return fmt.Errorf("separador %s invalido. El separador valido es '/'", separadorStr)
	}

	return nil
}
