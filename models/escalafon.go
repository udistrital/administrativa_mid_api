package models

type Escalafon struct {
	Descripcion     string `orm:"column(descripcion);null"`
	NombreEscalafon string `orm:"column(nombre_escalafon)"`
	Id              int    `orm:"column(id_escalafon);pk"`
}