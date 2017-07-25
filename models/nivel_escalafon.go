package models

type NivelEscalafon struct {
	Id              int    `orm:"column(id);pk"`
	NombreEscalafon string `orm:"column(nombre_escalafon)"`
}