package models

import (
	"time"
)

type FormacionAcademica struct {
	Id                int          `orm:"column(id);pk"`
	PersonaId         int          `orm:"column(persona_id)"`
	InstitucionId     *Institucion `orm:"column(institucion_id);rel(fk)"`
	ProgramaId        *Programa    `orm:"column(programa_id);rel(fk)"`
	NombreProyecto    string       `orm:"column(nombre_proyecto);null"`
	Validacion        bool         `orm:"column(validacion)"`
	FechaInicio       time.Time    `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion time.Time    `orm:"column(fecha_finalizacion);type(date)"`
	AreaConocimiento  string       `orm:"column(area_conocimiento);null"`
	FechaDato         time.Time    `orm:"column(fecha_dato);type(date);null"`
	Vigente           bool         `orm:"column(vigente);null"`
	Titulo            *Titulo      `orm:"column(titulo);rel(fk)"`
}