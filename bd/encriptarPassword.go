package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword es la rutina que me permite encriptar un password recibido por parametro */
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //El algoritmo de encriptación hará (2 elevado al costo) pasados por el texto
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}