package models

import (
	"time"
)

type PagoMensual struct {
	Id                int             
	NumeroContrato    string             
	VigenciaContrato  float64            
	Mes               float64            
	Persona           string             
	EstadoPagoMensual *EstadoPagoMensual
	Responsable       string             
	FechaModificacion time.Time          
	CargoResponsable  string             
	Ano               float64            
}