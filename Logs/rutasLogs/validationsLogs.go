package rutasLogs

import (
	"errors"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Logs/modelosLogs"
)


type OPERATION = modelosLogs.OPERATION

func ValidateEdit(log Log) error {

	var err error ;

	if err = dontContainAnyField(log); err != nil {return err}
	if err = idIsPresent(log.ID); err != nil {return err}
	if err = idUserIsPresent(log.Id_usuario); err != nil {return err}
	if err = usernameIsPresent(log.Nombre_usuario); err != nil{ return err}

	if log.Descripcion != ""{
		if err = validateDescription(log.Descripcion); err != nil {return err}
	}
	if log.Relevancia != 0 { 
		if err = validateRelevancia(log.Relevancia); err != nil {return err}
	}
	
	if log.Accion != ""{
		if err = validateAction(log.Accion); err != nil {return err}
	}


	return nil
}

func CreateValidation(log Log) error {
	var err error

	if err = usernameIsPresent(log.Nombre_usuario); err == nil{
		return errors.New("no presente el username")
	}

	if err = idUserIsPresent(log.Id_usuario); err == nil{
		return errors.New("no presente el id_usuario")
	}

	if len(log.Nombre_usuario) < 3 {
		return errors.New("el username debe tener mas de 3 caracteres")
	}

	if log.Id_usuario < 0{
		return errors.New("el id del usuario debe ser no negativo")
	}

	
	if err = validateDescription(log.Descripcion); err != nil {return err}
	if err = validateRelevancia(log.Relevancia); err != nil {return err}
	if err = validateAction(log.Accion);err != nil {return err}

	return nil
}

func validateAction(oPERATION OPERATION) error {

	if modelosLogs.IsValidAction(string(oPERATION)) {return nil}
	return errors.New("accion invalida")
}

func usernameIsPresent(s string) error { //devuelve error si el username esta presente
	if s != "" {return errors.New("no puedes modificar el username")}
	return nil
}

func idIsPresent(u uint) error { //devuelve error si el id esta presente

	if u != 0 {
		err := 	errors.New("no puedes cambiar el id del log")
		return err
	}
	return nil
}

func idUserIsPresent(u int) error {

	if u != 0 {
		err := errors.New("no puedes cambiar el id_usuario del log")
		return err
	}
	return nil
}



func validateRelevancia(i int) error {
	if i < 1 || i > 100 {
		return errors.New("la relevancia debe estar entre 1 y 100")
	}
	return nil
}


func validateDescription(s string) error {

	if len(s) < 3{
		return errors.New("la descripcion debe ser mayor a 3 caracteres")
	}
	return nil
}



func dontContainAnyField(log Log) error {
	if log.Descripcion == "" && log.Accion == "" && log.Relevancia == 0 {
		return errors.New("no se ha colocado ningun field")
	}
	return nil
}

