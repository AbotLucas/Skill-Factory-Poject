package routers

import (
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* Alta relacion realiza el registro de la relacion entre users */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	//Obtengo el id que viene como parametro
	ID := r.URL.Query().Get("id")
	if len(ID) < 0 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	//Definimos un modelo relacion en donde guardemos lo que vamos a guardar en la bd
	var rel models.Relacion
	//Colocamos como UsuarioID al que  tenemos grabado en la var globarl, que es el que esta loguead
	rel.UsuarioID = IDUsuario
	//Colocamos en el id de la relacion con otro usuario el usuario que viene como parametro
	rel.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(rel)
	//si hubo un error
	if err != nil {
		http.Error(w, "Ocurrio un errr al intentar insertar la relacion, intentelo nuevamente. "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
