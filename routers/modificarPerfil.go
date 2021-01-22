package routers

import (
	"encoding/json"
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* ModificarPerfil, modifica el perfil de user */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	
	var usu models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usu)

	if err != nil {
		/* Es un json mal construido */
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	
	status, err = bd.ModificoRegistro(usu, IDUsuario)
	//IDUsuario es la variable global que setamos antes con el ID
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. intente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}