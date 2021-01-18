package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/jwt"
	"github.com/abotlucas/microblogging/models"
)

/* Login realiza el login de un user
Recibe como parametro lo mismo que todos los endpoints y no retorna nada,
como los otros endpoints, son practicamente metodos */
func Login(w http.ResponseWriter, r *http.Request) {
	//Vamos a setear en el header el contenido que devolveremos (w)
	//Será de tipo Json
	w.Header().Add("content-type", "application/json")
	var usu models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usu)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(usu.Email) == 0 {
		http.Error(w, "El mail de usuario es requerido ", 400)
	}
	documento, existe := bd.IntentoLogin(usu.Email, usu.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}
	//Si el token se genero:
	resp := models.RespuestaLogin {
		Token : jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Tambien generaremos una cookie
	//Generamos un campo fecha para ver la expiracion de la cookie
	expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}