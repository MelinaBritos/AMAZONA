package rutasLogs

import (
	"strconv"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/baseDeDatos"
)



func findAll()  ([]Log, error) {
	var logs []Log
	err := baseDeDatos.DB.Find(&logs).Error
	return logs, err
}

func findById(id string)  (Log, error){
	var log Log

	_id, err := strconv.Atoi(id)

	if err != nil{
		return log , err
	}

	err = baseDeDatos.DB.Where("id = ?", _id).First(&log).Error
	return log, err
}

func editById(id string) error{

	var log Log
	_id, err := strconv.Atoi(id)

	if err != nil{
		return err
	}

	
	err = baseDeDatos.DB.Model(&log).Where("id = ?", _id).Updates(&log).Error

	return err
}

func CreateLog(log Log) error {
	
	var err error
	if err = CreateValidation(log); err != nil{
		return err
	}

	err = baseDeDatos.DB.Create(&log).Error

	if err != nil{return err}
	return nil
}

func DeleteById(id string) error {

	var log Log
	_id, err := strconv.Atoi(id)

	if err != nil{
		return err
	}

	err = baseDeDatos.DB.Where("id = ?", _id).Delete(&log).Error
	return err
}

func DeleteByUsername(username string) error{
	var log Log

	err := baseDeDatos.DB.Where("nombre_usuario = ?", username).Delete(&log).Error
	return err
}

func DeleteByIdU(id uint) error {

	var log Log
	err := baseDeDatos.DB.Where("id = ?", id).Delete(&log).Error
	return err
}