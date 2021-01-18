package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/* Declaración de la estructura usuario para MongoDB */
type Usuario struct {
	/* El ID de MongoDB no es un numero, es un tipo de dato binario
	de un tipo llamado ObjectID, es un slice de bits, que nos llevara a 
	utilizar luego algunas funciones especiales por este tipo de datos */
	/* omitempty > Si viene vacion no lo tenga en cuenta (Como el not null en sql)*/
	ID			primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	Nombre		string 				`bson:"nombre" json:"nombre,omitempty"`
	Apellidos	string 				`bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento	time.Time 		`bson:"fechaNacimiento"	json:"fechaNacimiento,omitempty"`
	Email		string 				`bson:"email" json:"email"`
	Password	string 				`bson:"password" json:"password,omitempty"`
	Avatar		string 				`bson:"avatar" json:"avatar,omitempty"`
	Banner		string 				`bson:"banner" json:"banner,omitempty"`
	Biografia	string 				`bson:"biografia" json:"biografia,omitempty"`
	Ubicacion	string 				`bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb	string 				`bson:"sitioWeb" json:"sitioWeb,omitempty"`

}