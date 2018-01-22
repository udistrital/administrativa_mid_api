package models

type DisponibilidadApropiacion struct {
	Id                   int
	Disponibilidad       *Disponibilidad
	Apropiacion          *Apropiacion
	Valor                float64
	FuenteFinanciamiento *FuenteFinanciamiento
}
