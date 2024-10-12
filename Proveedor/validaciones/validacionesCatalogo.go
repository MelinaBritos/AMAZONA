package validaciones

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)

func ValidarCatalogo(catalogo modelosProveedor.Catalogo) error {

	if err := validarIdCatalogo(catalogo.Id_catalogo); err != nil {
		return err
	}

	if !existeProveedor(catalogo.Id_proveedor) {
		return fmt.Errorf("no existe el proveedor %d", catalogo.Id_catalogo)
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

	mesStr := mes_vigencia[:2]
	mesInt, err := strconv.Atoi(mesStr)
	if err != nil {
		return errors.New("error al ingresar el mes. intente nuevamente")
	}

	if mesInt <= 0 || mesInt >= 12 {
		return errors.New("valor invalido del mes. Los valores validos son 1-12")
	}

	if separadorStr := string(mes_vigencia[2]); separadorStr != "/" {
		return fmt.Errorf("separador %s invalido. El separador valido es '/'", separadorStr)
	}

	anioStr := mes_vigencia[3:]
	anioInt, err := strconv.Atoi(anioStr)
	if err != nil {
		return errors.New("error al ingresar el mes. intente nuevamente")
	}

	if anioInt <= 0 {
		return errors.New("valor invalido del año")
	}

	return nil
}
