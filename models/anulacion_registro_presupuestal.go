package models

import "time"

type AnulacionRegistroPresupuestal struct {
	Id                   int
	Consecutivo          int
	Motivo               string
	FechaRegistro        time.Time
	TipoAnulacion        string
	EstadoAnulacion      *EstadoAnulacion
	JustificacionRechazo string
	Responsable          int
	Solicitante          int
	Expidio              int
}
