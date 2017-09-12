package models

import (
	"time"
)

type EstadoContrato struct {
	NombreEstado  string    `orm:"column(nombre_estado);null"`
	FechaRegistro time.Time `orm:"column(fecha_registro);type(date);null"`
	Id            int       `orm:"column(id);pk;auto"`
}
