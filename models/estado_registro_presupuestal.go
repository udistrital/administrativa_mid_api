package models

type EstadoRegistroPresupuestal struct {
	Id                int
	Nombre            string
	Descripcion       string
	Activo            bool
	CodigoAbreviacion string
	NumeroOrden       string
}
