package models

import (
	"time"
)

type ContratoEstado struct {
	NumeroContrato string          `orm:"column(numero_contrato);null"`
	Vigencia       int             `orm:"column(vigencia);null"`
	FechaRegistro  time.Time       `orm:"column(fecha_registro);type(timestamp without time zone);null"`
	Id             int             `orm:"column(id);pk;auto"`
	Estado         *EstadoContrato `orm:"column(estado);rel(fk)"`
	Usuario        string          `orm:"column(usuario);null"`
}
