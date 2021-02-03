package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos - Lee los usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacion conmigo.
Tiene muchos parametros:
-El *ID del usuario que esta leyendo los usuarios
-La *pagina porque si hay muchos usuarios hay que paginar
-El *search que es una palabra para la busqueda(Todos los usuarios que contengan la palabra search)
-El *tipo de busqueda que vamos a hacer, si listamos todos, los que seguimos o los que nos siguen
Devuelve un booleano y un Slice del tipo models.usuario, ya que devolveremos n usuarios que responden a nuestra busqueda  */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	/* creamos la variable para el resultado, que va a ser un slice de usuarios */
	var results []*models.Usuario

	findOptions := options.Find()
	/* Ahora le damos las propiedades a nuestro find */
	/* defino el salto para posicionarme en los resultados */
	findOptions.SetSkip((page - 1) * 20)
	//devuelvo de a 20 resultados
	findOptions.SetLimit(20)

	/* Ahora haremos la condicion con un mapeo de bson,
	con el nombre de tipo regex (expresion regular)
	?i hace que no distinga entre mayusculas y minusculas */
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	//ahora ejecutamos el find que devuelve el resultado en un cursor
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		//si no fue satisfactoria la busqueda en la BD
		fmt.Println(err.Error())
		//devuelve results vacio y false
		return results, false
	}

	//vamos a definir dos variables booleanas
	var encontrado, incluir bool

	//el  next me permite avanzar al siguiente registro
	for cur.Next(ctx) {

		//defino un usuario para trabajar con el elemento que puedo despues incluir en results
		var s models.Usuario
		//grabo lo retornado por el cursor en el modelo de usuario para leer los campos
		err := cur.Decode(&s)
		if err != nil {
			//encontro un error
			fmt.Println(err.Error())
			return results, false
		}

		//creo una variable de tipo relacion para consultar sobre la relacion con el usuario
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		//por cada iteracion tengo que ver si al usuario lo debo incluir o no en el resultado
		incluir = false

		encontrado, err = ConsultoRelacion(r)

		if tipo == "new" && encontrado == false {
			//lo tengo que incluir en la lista, porque no lo encontro en la relacion
			incluir = true
		}
		//si es de tipo follow solo quiero listar los que yo estoy siguiendo
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			//seria el caso de que me estoy siguiendo a mi mismo
			incluir = false
		}
		if incluir == true {
			//hago un blanqueo de los campos que no me interesa incluir
			//solo quiero el avatar, el nombre, los apellidos y la fecha de nacimiento
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			//agrego con append el modelo de usuario s en results
			results = append(results, &s)
		}

	}

	//me fijo si hubo un error del cursor
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	//si no hubo errr, cierro el cursor y retorno la lista de resultados
	cur.Close(ctx)
	return results, true
}
