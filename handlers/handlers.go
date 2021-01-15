package handlers
 
import (
    "log"
    "net/http"
    "os"
    
    "github.com/gorilla/mux"
    "github.com/abotlucas/microblogging/middlew"
    "github.com/abotlucas/microblogging/routers"
    "github.com/rs/cors"
    
 
)
/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores(){
    router := mux.NewRouter()

    /* Si alguien entro a /registro con metodo POST debe ejecutar el middleware de ChequeoBD
    Y le paso la funcion de routers.Registro */
    /* Por cada endPoint vamos a tener un renglon de codigo que permita manejar la funcion */
    router.HandleFunc("/regitro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    handler := cors.AllowAll().Handler(router)
    log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
