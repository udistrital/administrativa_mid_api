package models

import (
	"time"
)

type ExperienciaDocente struct {
	Id                int             `orm:"column(id);pk"`
	InstitucionId     *Institucion    `orm:"column(institucion_id);rel(fk)"`
	TipoActividad     string          `orm:"column(tipo_actividad)"`
	CampoEnseñanza    string          `orm:"column(campo_enseñanza);null"`
	PersonaId         int             `orm:"column(persona_id)"`
	Validacion        bool            `orm:"column(validacion)"`
	FechaInicio       time.Time       `orm:"column(fecha_inicio);type(date)"`
	FechaFinalizacion time.Time       `orm:"column(fecha_finalizacion);type(date)"`
	TipoDedicacionId  *TipoDedicacion `orm:"column(tipo_dedicacion_id);rel(fk)"`
	FechaDato         time.Time       `orm:"column(fecha_dato);type(date);null"`
	Vigencia          bool            `orm:"column(vigencia);null"`
}