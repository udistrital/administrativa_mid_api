package models

import "github.com/astaxie/beego/orm"

type ResolucionVinculacionDocente struct {
	NivelAcademico string `orm:"column(nivel_academico)"`
	Dedicacion     string `orm:"column(dedicacion)"`
	IdFacultad     int    `orm:"column(id_facultad)"`
	Id             int    `orm:"column(id_resolucion);pk"`
}

func init() {
	orm.RegisterModel(new(ResolucionVinculacionDocente))
}
