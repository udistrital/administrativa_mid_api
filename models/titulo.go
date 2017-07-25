package models

type Titulo struct {
	Id          int       `orm:"column(id);pk"`
	Nombre      string    `orm:"column(nombre)"`
	Descripcion string    `orm:"column(descripcion);null"`
	Programa    *Programa `orm:"column(programa);rel(fk)"`
}