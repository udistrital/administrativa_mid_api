package models

import (
	"time"
)

type EscalafonPersona struct {
	Estado           bool                       `orm:"column(estado)"`
	Observacion      string                     `orm:"column(observacion);null"`
	FechaRegistro    time.Time                  `orm:"column(fecha_registro);type(date)"`
	IdEscalafon      *Escalafon                 `orm:"column(id_escalafon);rel(fk)"`
	IdPersonaNatural int 						`orm:"column(id_persona_natural)"`
	Id               int                        `orm:"column(id_escalafon_persona);pk"`
}