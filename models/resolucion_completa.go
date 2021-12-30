package models

type ResolucionCompleta struct {
	Vinculacion             ResolucionVinculacionDocente
	Consideracion           string
	Preambulo               string
	Vigencia                int
	Numero                  string
	Id                      int
	CuadroResponsabilidades string
	Articulos               []Articulo
	OrdenadorGasto          OrdenadorGasto
	Titulo                  string
}
