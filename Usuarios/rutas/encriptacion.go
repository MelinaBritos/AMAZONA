package rutas
import (
	"golang.org/x/crypto/bcrypt"
)

func Encriptar(contraseña string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(contraseña), 4)
	
	if err != nil {
		return "", err
	}

	value := string(hashed)
	return value, nil
}

func Equals(contraseña string, hashed string) error  {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(contraseña))
	return err
}