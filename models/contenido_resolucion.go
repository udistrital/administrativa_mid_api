package models

type Paragrafo struct {
	Id     int
	Numero int
	Texto  string
}

type Articulo struct {
	Id         int
	Numero     int
	Texto      string
	Paragrafos []Paragrafo
}

type ResolucionCompleta struct {
	Vinculacion   ResolucionVinculacionDocente
	Consideracion string
	Preambulo     string
	Vigencia      int
	Numero        string
	Id            int
	Articulos     []Articulo
	OrdenadorGasto OrdenadorGasto
	
}

type ComponenteResolucion struct {
	Id              int                   `orm:"column(id);pk;auto"`
	Numero          int                   `orm:"column(numero)"`
	ResolucionId    *Resolucion           `orm:"column(resolucion_id);rel(fk)"`
	Texto           string                `orm:"column(texto)"`
	TipoComponente  string                `orm:"column(tipo_componente)"`
	ComponentePadre *ComponenteResolucion `orm:"column(componente_padre);rel(fk);null"`
}
