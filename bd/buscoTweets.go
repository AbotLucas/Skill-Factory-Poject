package bd

import (
	"context"
	"time"
	
	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* BuscoTweets - busca los tweets de un id de usuario en la bd */
func BuscoTweets(id string) ([] *models.Tweet, error) {
//Vamos a usar una peticion get, ya que vendra por parametro el id
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()

db := MongoCN.Database("microblogging")
col := db.Collection("tweet")

var tweets []*models.Tweet
//objID, _ := primitive.ObjectIDFromHex(ID)

condicion := bson.M{"userid": id}

findOptions := options.Find()
//Buscar multiples registros retorna un cursor
//Iterando nuestro cursor podremos decodear todos los documentos uno a uno
cursor, err := col.Find(ctx, condicion, findOptions)

if err != nil {
	return tweets, err
}
//Iteramos el cursor
for cursor.Next(ctx) {

	//creamos un valor para ser decoded
	var elem models.Tweet
	err := cursor.Decode(&elem)
	if err != nil {
		return tweets, err
	}

	tweets = append(tweets, &elem)
}

if err := cursor.Err(); err != nil {
	return tweets, err
}

//cerramos el cursor al finalizar
cursor.Close(ctx)

return tweets, nil

}