package models

import (
	"time"
)

type Investigacion struct {
	Id                  int                `orm:"column(id);pk"`
	PersonaId           int                `orm:"column(persona_id)"`
	NombreInvestigacion string             `orm:"column(nombre_investigacion)"`
	FechaInicio         time.Time          `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion   time.Time          `orm:"column(fecha_finalizacion);type(date)"`
	TipoInvestigacion   string             `orm:"column(tipo_investigacion);null"`
	InstitucionId       *Institucion       `orm:"column(institucion_id);rel(fk)"`
	TipoInvestigacionId *TipoInvestigacion `orm:"column(tipo_investigacion_id);rel(fk)"`
	FechaDato           time.Time          `orm:"column(fecha_dato);type(date);null"`
	Vigente             bool               `orm:"column(vigente);null"`
}