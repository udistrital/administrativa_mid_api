package models

//FuenteFinanciacionRubroNecesidad modelo de la aplicación necesidades que guarda las fuentes asociadas a una solicitud
type FuenteFinanciacionRubroNecesidad struct {
	Id                   int
	FuenteFinanciamiento int
	Apropiacion          int
	MontoParcial         float64
	Necesidad            *Necesidad
	InfoFuente           *[]FuenteFinanciamiento
}
