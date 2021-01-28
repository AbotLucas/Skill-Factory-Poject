package bd

import (
	"context"
	"log"
	"time"

	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* BuscoTweets - busca los tweets de un id de usuario en la bd */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	//Vamos a usar una peticion get, ya que vendra por parametro el id
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	var tweets []*models.DevuelvoTweets

	condicion := bson.M{"userid": ID}

	/* EL paquete options permite filtrar y dar un comportamiento a una consulta */
	findOptions := options.Find()

	/* Le diremos que como maximo nos traiga 20 registros */
	findOptions.SetLimit(20)

	/* Con sort le digo como va a ir ordenando lo que traiga, en este caso ordenados por fecha
	en orden descendente (se indica con el -1 en Value)*/
	findOptions.SetSort(bson.D{{Key: "fecha", Value: -1}})

	/* Ahora le indicaremos cuantos ira salteando, al principio ninguno, despues 20, despues 40, y asi */
	findOptions.SetSkip((pagina - 1) * 20)

	//Buscar multiples registros retorna un cursor
	//Iterando nuestro cursor podremos decodear todos los documentos uno a uno
	cursor, err := col.Find(ctx, condicion, findOptions)

	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	//Iteramos el cursor
	for cursor.Next(context.TODO()) {

		//creamos una variable para alojar el decoded decada registro del cursor
		var elem models.DevuelvoTweets
		err := cursor.Decode(&elem)
		if err != nil {
			/* si hubo error retorno resultado vacio */
			return tweets, false
		}
		/* Agrego un elemento al slice */
		tweets = append(tweets, &elem)
	}

	if err := cursor.Err(); err != nil {
		return tweets, false
	}

	//cerramos el cursor al finalizar
	cursor.Close(ctx)

	return tweets, true

}
