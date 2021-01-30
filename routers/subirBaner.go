package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/abotlucas/microblogging/bd"
	"github.com/abotlucas/microblogging/models"
)

/*SubirBanner suve el avatar al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")

	// Se extrae la extension del archivo original del avatar
	var extension = strings.Split(handler.Filename, ".")[1]

	/* En lugar de guardar el nombre que cada user le pone al archivo,
	los coloco en una carpeta avatars y como nombre le pongo el idUsuario,
	ya que tendremos solo uno por usuario.
	Ls archvos se guardan en una carpeta que debe estar creada para que todo funcione:
	Carpeta: uploads/banners */
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	//creamos el manejador de archivos con permiso de lectura, escritura y ejecucion
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}
	/* si no hubo problemas al abrir la imagen, vamos a hacer la copia en f de file,
	ademas de copiar lo renombra: */
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	/* Ahora vamos a grabar en la bd el cambio en el campo avatar */
	var usuario models.Usuario
	var status bool
	usuario.Banner = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el Banner en la bd "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	// y le damos un status created:
	w.WriteHeader(http.StatusCreated)

}
