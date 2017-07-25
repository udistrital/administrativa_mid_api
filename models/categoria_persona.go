package models

import (
	"time"
)

type CategoriaPersona struct {
	Validacion      bool           `orm:"column(validacion);null"`
	PersonaId       int            `orm:"column(persona_id)"`
	IdTipoCategoria *TipoCategoria `orm:"column(id_tipo_categoria);rel(fk)"`
	Vigente         bool           `orm:"column(vigente);null"`
	FechaDato       time.Time      `orm:"column(fecha_dato);type(date);null"`
	Id              int            `orm:"column(id);pk"`
}
