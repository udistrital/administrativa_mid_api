package models

type Dedicacion struct {
	Descripcion      string `orm:"column(descripcion);null"`
	NombreDedicacion string `orm:"column(nombre_dedicacion)"`
	Id               int    `orm:"column(id_dedicacion);pk"`
}