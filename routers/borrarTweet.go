package routers

import (
	"encoding/json"
	"net/http"

	"github.com/abotlucas/microblogging/bd"
	//"github.com/abotlucas/microblogging/models"
)

func BorrarTweet(w http.ResponseWriter, r *http.Request) {

	idTweet := r.URL.Query().Get("id")
	if len(idTweet) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	// si tdo anduvo bien:
	tweet, err := bd.BorrarTweet(idTweet)
	if err != nil {
		http.Error(w, "Ocurrio un errro al intentar buscar el registro", 400)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)



}