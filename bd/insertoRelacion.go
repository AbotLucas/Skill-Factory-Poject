package bd

import (
	"context"
	"time"

	"github.com/abotlucas/microblogging/models"
)

/*InsertoRelacion - graba la relacion en la bd */
func InsertoRelacion(relac models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	//Directamente trato de insertar en la bd el modelo que recibo
	_, err := col.InsertOne(ctx, relac)
	if err != nil {
		return false, err
	}
	return true, nil

}
