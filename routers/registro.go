package routers

import (
	"enconding/json"
	"net/http"
	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* Registro es la funcion para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(),400)
		return
	}
	/* Si no hubo error con el Body hago unas validaciones */
	if len(user.Email == 0) {
		http.Error(w, "El mail de usuario es requerido", 400 )
		return
	}
	if len(user.Password < 6) {
		http.Error(w, "Debe especificar un password de al menos 6 caracteres", 400 )
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(user.Email) 
	if encontrado == true{
		http.Error(w, "Ya existe un usuario registrado con ese Email",400)
		return
	}
	
	_, status, err := bd.InsertoRegistro(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario")
		return
	}

	/* Si llego hasta aca todo anduvo bien */
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHEader(http.StatusCreated)
}