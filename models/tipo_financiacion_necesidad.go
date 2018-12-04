package models

type TipoFinanciacionNecesidad struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Estado            bool
	NumeroOrden       string
}
