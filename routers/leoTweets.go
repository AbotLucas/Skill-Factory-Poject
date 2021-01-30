package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/*LeoTweets - lee los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	/* Envio el id de usuario para recolectar los id */
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}
	/* Vamos a trabajar con paginacion */
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	// si todo anduvo bien:
	var tweets []*models.DevuelvoTweets
	tweets, ok := bd.LeoTweets(ID, pag)
	if ok == false {
		http.Error(w, "ocurrio un error al intentar buscar los tweets"+err.Error(), 400)
		return
	}
	/* Establesco el tipo de header */
	w.Header().Set("Content-type", "application/json")
	/* Le doy un status created */
	w.WriteHeader(http.StatusCreated)
	/* le devolvemos los tweets encontrados */
	json.NewEncoder(w).Encode(tweets)
}
