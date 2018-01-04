package models

import (
		"time"
)

type ResolucionEstado struct {
	Id            int               `orm:"column(id);pk;auto"`
	FechaRegistro time.Time         `orm:"column(fecha_registro);type(timestamp without time zone)"`
	Usuario       string            `orm:"column(usuario);null"`
	Estado        *EstadoResolucion `orm:"column(estado);rel(fk)"`
	Resolucion    *Resolucion       `orm:"column(resolucion);rel(fk)"`
}
