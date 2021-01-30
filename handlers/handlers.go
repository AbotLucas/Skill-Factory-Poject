package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/abotlucas/microblogging/middlew"
	"github.com/abotlucas/microblogging/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	/* Si alguien entro a /registro con metodo POST debe ejecutar el middleware de ChequeoBD
	   Y le paso la funcion de routers.Registro */
	/* Por cada endPoint vamos a tener un renglon de codigo que permita manejar la funcion */
	//Registro
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	//Login
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	//VerPerfil
	router.HandleFunc("/verperfil", middlew.ChequeoBD((middlew.ValidoJWT(routers.VerPerfil)))).Methods("GET")
	//ModificarPerfil
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	//Tweet
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	//BOrrarTweet
	router.HandleFunc("/borrotweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.BorroTweet))).Methods("DELETE")
	//LeoTweets
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	//SubirAvatar
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	//ObtenerAvatar
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	//SubirBanner
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	//ObtenerBanner
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
