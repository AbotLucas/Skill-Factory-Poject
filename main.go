package main
 
import (
    "log"
    "github.com/abotlucas/microblogging/handlers"
    "github.com/abotlucas/microblogging/bd"
)
func main(){
    if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
	handlers.Manejadores()
 
}
