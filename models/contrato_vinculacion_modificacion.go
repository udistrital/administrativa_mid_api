package models

type ContratoVinculacionModificacion struct {
	ContratoGeneral    *ContratoGeneral
	VinculacionDocente *VinculacionDocente
	ActaInicio         *ActaInicio
	Cdp                *ContratoDisponibilidad
	ContratoCancelado  *ContratoCancelado
}
