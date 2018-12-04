package models

//EstadoNecesidad relaciona la información del estado de la necesidad
type EstadoNecesidad struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Estado            bool
	NumeroOrden       string
}
