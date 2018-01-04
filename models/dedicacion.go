package models

import "github.com/astaxie/beego/orm"

type Dedicacion struct {
	Descripcion      string `orm:"column(descripcion);null"`
	NombreDedicacion string `orm:"column(nombre_dedicacion)"`
	Id               int    `orm:"column(id_dedicacion);pk"`
}

func init() {
	orm.RegisterModel(new(Dedicacion))
}
