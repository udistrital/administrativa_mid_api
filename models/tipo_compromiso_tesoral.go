package models

type TipoCompromisoTesoral struct {
	Id                  int
	Nombre              string
	Activo              bool
	CategoriaCompromiso *CategoriaCompromiso
	Descripcion         string
	CodigoAbreviacion   string
	NumeroOrden         float64
}
