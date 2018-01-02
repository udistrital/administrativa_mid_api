package models

import (

	"time"

)

type EstadoResolucion struct {
	Id            int       `orm:"column(id);pk;auto"`
	FechaRegistro time.Time `orm:"column(fecha_registro);type(date)"`
	NombreEstado  string    `orm:"column(nombre_estado)"`
}
