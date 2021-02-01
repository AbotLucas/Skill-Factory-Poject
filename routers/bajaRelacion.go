package routers

import (
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* BajaRelacion - Realiza el borrado dela relacion entre users */
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	//Obtengo el ID que iene por parametros
	ID := r.URL.Query().Get("id")

	//Definimos un modelo relacion en donde guardaremos lo que vamos a borrar de la bd
	var rel models.Relacion
	//Colocamos como usuarioId al que tenemos grabado en la var global que es el que esta logueado
	rel.UsuarioID = IDUsuario
	//Colocamos como el usuario en relacion el que nos llego por post
	rel.UsuarioRelacionID = ID

	//Le paso a borrorelacion el modelo que arme
	status, err := bd.BorroRelacion(rel)
	if err != nil {
		http.Error(w, "No se ha logrado borrar la relacion. "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion. "+err.Error(), http.StatusBadRequest)
		return
	}

	//Si todo estuvo bien con el borrado, mado un StatusCreated
	w.WriteHeader(http.StatusCreated)
}
