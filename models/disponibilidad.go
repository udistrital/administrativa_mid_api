package models


type Disponibilidad struct {
	Id                        int                          `orm:"auto;column(id);pk"`
	Vigencia                  float64                      `orm:"column(vigencia)"`
	NumeroDisponibilidad      float64                      `orm:"column(numero_disponibilidad);null"`

}


type DisponibilidadApropiacion struct {
	Id                   int                   `orm:"auto;column(id);pk"`
	Disponibilidad       *Disponibilidad       `orm:"column(disponibilidad);rel(fk)"`
	Apropiacion          *Apropiacion          `orm:"column(apropiacion);rel(fk)"`
	Valor                float64               `orm:"column(valor);null"`
	FuenteFinanciamiento *FuenteFinanciamiento `orm:"column(fuente_financiamiento);rel(fk);null"`
}

type Apropiacion struct {
	Id       int                `orm:"column(id);pk;auto"`
	Vigencia float64            `orm:"column(vigencia)"`
	Valor    float64            `orm:"column(valor)"`
	Saldo    int
}

type FuenteFinanciamiento struct {
	Id          int    `orm:"column(id);pk;auto"`
	Descripcion string `orm:"column(descripcion);null"`
	Nombre       string `orm:"column(nombre)"`
	Codigo      string `orm:"column(codigo)"`
}

type DatosApropiacion struct {

	Anulado      float64
	Comprometido float64
	Saldo        float64
}
