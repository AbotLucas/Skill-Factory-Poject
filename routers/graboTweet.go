package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/*GraboTweet permite grabar el tweet en la Bd */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	//decodificamos el body armamos un regitro
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	//Para insertarlo en la base de datos necsitamos mapearlo a un bson
	_, status, err := bd.InsertoTweet(registro)

	//si hay un error
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intentelo nuevamente. "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
