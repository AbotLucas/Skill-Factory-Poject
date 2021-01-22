package bd

import (
	"context"
	"time"

	"github.com/abotlucas/microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModificoRegistro permite modificar el perfil del user */
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblog")
	col := db.Coollection("usuarios")

	/* SUPONGO QUE ME VAN A ENVIAR UN CMPO A MODIFICAR A LA VEZ,
	por eso me fijo si lo que viene tiene valor (largo mayor a cero) 
	Creamos un mapa de interfaces para armar el registro de actualizacion a la BD
	poniendole la info que hay que modificar*/

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.fechaNacimiento) > 0 {	
		registro["fechaNacimiento"] = u.fechaNacimiento
	}
	if len(u.avatar) > 0 {	
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {	
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {	
		regitro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {	
		regitro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {	
		regitro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": regitro,
	}
	//Paso el object id fromo Hex
	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil
}