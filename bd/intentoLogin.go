package bd

import (
	"github.com/abotlucas/microblogging/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de Login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	//Ahora comparo la pass con la de la BD
	//Creo una var slice de bytes
	passwordEnBytes := []byte(password)
	//Creo otra variable con la pass que tengo en la BD para el user
	passEnBDEnBytes := []byte(usu.Password)
	//Ahora llamo una funcion del package bcrypt que compara las password
	err := bcrypt.CompareHashAndPassword(passEnBDEnBytes, passwordEnBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
