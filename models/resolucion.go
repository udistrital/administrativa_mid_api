package models

import (
	"time"
)

type Resolucion struct {
	Objeto                  *string         `orm:"column(objeto);null"`
	FechaRegistro           time.Time       `orm:"column(fecha_registro);type(date)"`
	Estado                  bool            `orm:"column(estado)"`
	ConsideracionResolucion string          `orm:"column(consideracion_resolucion)"`
	PreambuloResolucion     string          `orm:"column(preambulo_resolucion)"`
	IdTipoResolucion        *TipoResolucion `orm:"column(id_tipo_resolucion);rel(fk)"`
	IdDependencia           int             `orm:"column(id_dependencia)"`
	Vigencia                int             `orm:"column(vigencia)"`
	FechaExpedicion         *time.Time      `orm:"column(fecha_expedicion);type(date);null"`
	NumeroResolucion        string          `orm:"column(numero_resolucion)"`
	Id                      int             `orm:"column(id_resolucion);pk:auto"`
}