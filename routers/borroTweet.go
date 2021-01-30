package routers

import (
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	//"github.com/abotlucas/microblogging/models"
)

/*BorroTweet permite eliminar un tweet determinado */
func BorroTweet(w http.ResponseWriter, r *http.Request) {

	idTweet := r.URL.Query().Get("id")
	if len(idTweet) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	// si tdo anduvo bien:
	err := bd.BorroTweet(idTweet, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un errro al intentar buscar el registro", 400)
		return
	}

	/* En la sig línea le decimos al navegador que si va respuesta será del tipo json
	   aunque no le mandamos respuesta, es una buena práctica, por si más adelante queremos enviarle una */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
