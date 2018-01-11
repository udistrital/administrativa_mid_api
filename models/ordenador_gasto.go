package models

import (
		"github.com/astaxie/beego/orm"
)


type OrdenadorGasto struct {
	Id                     int                     `orm:"column(id);pk;auto"`
  Cargo                 string                  `orm:"column(cargo)"`
	DependenciaId               int                 `orm:"column(dependencia_id)"`
	NombreOrdenador                string
}

func init() {
	orm.RegisterModel(new(OrdenadorGasto))
}
