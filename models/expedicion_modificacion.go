package models

import "time"

type ExpedicionModificacion struct {
	Vinculaciones   *[]ContratoVinculacionModificacion
	IdResolucion    int
	FechaExpedicion time.Time
}
