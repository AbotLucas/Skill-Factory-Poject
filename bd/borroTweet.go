package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(id string, UserID string) error {
	//vamos a usar una peticion DELETE que vendra con el id del tweet por parametro
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	/* El id en string que recibimos por parametro lo pasamos a objectId de MongoDB */
	objID, _ := primitive.ObjectIDFromHex(id)
	/* Creamos nuestro filtro para encontrar el tweet que buscamos */
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	/* Vamos a devolver el tweet que acabamos de eliminar por si acaso
	errFind := col.FindOne(ctx, condicion).Decode(&tweet)
	if errFind != nil {
		fmt.Println("Registro no encontrado " + errFind.Error())
		return tweet, errFind
	} */
	deleteResult, err := col.DeleteOne(ctx, condicion)
	/* if errFind != nil {
		fmt.Println("Error al eliminar el tweet" + errFind.Error())
		return tweet, err
	}
	*/
	fmt.Println("Se eliminaron: ", deleteResult.DeletedCount, " registros.")

	return err

}
