package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ConsultoRelacion consulta la relacion entre dos usuarios*/
func ConsultoRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	/* Mapeo a bsson la relacion que viene como parametro para hacer la consulta a la BD */
	condicion := bson.M{
		"usuarioid":         rel.UsuarioID,
		"usuariorelacionid": rel.UsuarioRelacionID,
	}

	/* Defino un modelo para contener el resultaod de la consulta */
	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	/* Imprimo por pantalla el resultado de la consulta */
	fmt.Println(resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
