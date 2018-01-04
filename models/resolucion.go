package models

import (
	"time"
)

type Resolucion struct {
	Id                      int             `orm:"column(id_resolucion);pk;auto"`
	NumeroResolucion        string          `orm:"column(numero_resolucion)"`
	FechaExpedicion         time.Time       `orm:"column(fecha_expedicion);type(date);null"`
	Vigencia                int             `orm:"column(vigencia)"`
	IdDependencia           int             `orm:"column(id_dependencia)"`
	IdTipoResolucion        *TipoResolucion `orm:"column(id_tipo_resolucion);rel(fk)"`
	PreambuloResolucion     string          `orm:"column(preambulo_resolucion)"`
	ConsideracionResolucion string          `orm:"column(consideracion_resolucion)"`
	Estado                  bool            `orm:"column(estado)"`
	FechaRegistro           time.Time       `orm:"column(fecha_registro);type(date)"`
	Objeto                  string          `orm:"column(objeto);null"`
	NumeroSemanas           int             `orm:"column(numero_semanas)"`
	Periodo                 int             `orm:"column(periodo)"`
}


type ObjetoResolucion struct {
	Resolucion       *Resolucion
	ResolucionVinculacionDocente *ResolucionVinculacionDocente
	ResolucionVieja   int
}

type ModificacionResolucion struct {
	Id             int       			`orm:"column(id);pk;auto"`
	ResolucionNueva       int      `orm:"column(resolucion_nueva)"`
	ResolucionAnterior     int    `orm:"column(resolucion_anterior)"`
}
