package models

import (
	"time"
)

/*Grabotweet es el frmato que tendra nustro tweet en la BD */
type GraboTweet struct {
	UserID  string    `bson: "userid" json:"userid,omitempty"`
	Mensaje string    `bson: "mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson: "fecha" json:"fecha,omitempty"`
}
