package routers

import ( 
	"encoding/json"
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/* VerTweets */
func VerTweets(w http.ResponseWriter, r *http.Request) {
	/* Envio el id de usuario para recolectar los id */
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	// si todo anduvo bien:
	var tweets []*models.Tweet
	tweets, err := bd.BuscoTweets(id)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar buscar los tweets"+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "encoding/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)
}