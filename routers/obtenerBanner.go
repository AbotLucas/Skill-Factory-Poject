package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/abotlucas/microblogging/bd"
)

/* ObtenerBanner envia el avatar al http */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	/* DE la URL vamos a obtener nuestro parametro id del usuario
	de quien obtendremos el avatar */
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar al parametro ID", http.StatusBadRequest)
		return
	}

	/* Busco el perfil por el ID que recibi y lo cargo en un modelo */
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	/* Intento abrir el archivo */
	OpenFile, err := os.Open("uploads/banners" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}
	//Copiamos al ResponeWriter el archivo de la imagen
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
	/*En el frontend no van a chequear si hubo un status 201,
	sino que le haya llegado o no una imagen,
	por eso no hace falta enviar ningún seteo, ningún status*/

}
