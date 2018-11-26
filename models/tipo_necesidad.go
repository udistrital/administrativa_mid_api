package models

type TipoNecesidad struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Estado            bool
	NumeroOrden       string
}
