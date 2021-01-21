package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* Creamos dos var globales para poder accederlas dentro y fuera del package */
/* Con los valores de mail e id utilizados en todos los endpoints */
/* Email */
var Email string

/* IDUsuario */
var IDUsuario string

 /* ProcesoToken - proceso token para extraer sus valores */
 /* Es de las funciones mas importantes por las veces que sera llamada,
 valida el token y dice si es credencial y si los privilegios son validos. */

 /* En Go si una funcion tiene varios parametros y uno es 'error' debe ir al final */
 func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	 miClave := []byte("SkillFactoryGo_Avalith")
	 //Creo una variable claims de tipo claim del models, se indica como puntero
	 //Porque la struct donde va el token debe ser puntero
	 claims := &models.Claim{}
	 //El token que viene en el Header comienza con la palabra Bearer, es un estandar, no es parte del token en si
	 splitToken := strings.Split(tk, "Bearer")
	 if len(splitToken) != 2{
		 //tiene que devolver dos elementos
		 return claims, false, string(""), errors.New("formato de token invalido")
		 //el error se crea con un mensaje que no puede tener ni mayusculas, ni tildes, ni simbolos
	 }
	 tk = strings.TrimSpace(splitToken[1])
	 tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		// el tercer parametro es una funcion anonima que recibe un token y resuleve todo ahi validando el token 
		return miClave, nil
	 })
	 if err == nil {
		 _, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
		}
		if !tkn.Valid {
			return claims, false, string(""), errors.New("token invalido")
		}
		return claims, false, string(""), err
 }