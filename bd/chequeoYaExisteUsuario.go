package bd

import (
	"context"
	"time"
	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ChequeoYaExisteUsuario recibe un email x parametro y chequea si ya existe en la bd */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) */
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	/* levanto la bd */
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	/* M es una funcion que formatea  o mapea a bson lo que recibe como json */
	condicion := bson.M{"email": email}

	/* en la variable resultado modlaremos un usuario */
	var resultado models.Usuario

	/* FindOne me devuelve un solo registro que cumpla con la condicion */
	err := col.FindOne(ctx, condicion).Decode(&resultado) //En este caso un user que tenga el mail enviado y lo guarda en "resultado"
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}