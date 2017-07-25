package models

import (
	"time"
)

type PuntoSalarial struct {
	Decreto       string    `orm:"column(decreto);null"`
	ValorPunto    int	   `orm:"column(valor_punto)"`
	FechaRegistro time.Time `orm:"column(fecha_registro);type(date)"`
	Vigencia      int       `orm:"column(vigencia)"`
	Id            int       `orm:"column(id);pk"`
}