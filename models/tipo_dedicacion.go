package models

type TipoDedicacion struct {
	Id                   int    `orm:"column(id);pk"`
	NombreTipoDedicacion string `orm:"column(nombre_tipo_dedicacion)"`
}