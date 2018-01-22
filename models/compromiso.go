package models

import "time"

type Compromiso struct {
	Id                    int
	Objeto                string
	Vigencia              float64
	FechaInicio           time.Time
	FechaFin              time.Time
	FechaModificacion     time.Time
	EstadoCompromiso      *EstadoCompromiso
	TipoCompromisoTesoral *TipoCompromisoTesoral
	UnidadEjecutora       int
}
