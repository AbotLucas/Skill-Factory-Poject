package models

/* REspuestaConsultaRelacion tiene el true o false que se obtiene al consultar la relacion entre dos usuarios */
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
