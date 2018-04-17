package models

import (
	"time"
)

type NovedadPostcontractual struct {
	Id              int      
	NumeroContrato  string    
	Vigencia        int       
	TipoNovedad     float64   
	FechaInicio     time.Time 
	FechaFin        time.Time 
	FechaRegistro   time.Time 
	Contratista     float64   
	NumeroCdp       int       
	VigenciaCdp     int    
	PlazoEjecucion  int       
	UnidadEjecucion int       
	ValorNovedad    float64  
}
