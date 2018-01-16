package models

import (

	"time"
	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
	Id              int       `orm:"column(id);pk;auto"`
	Estado          string    `orm:"column(estado)"`
	Numero          string    `orm:"column(numero)"`
	Vigencia        int       `orm:"column(vigencia)"`
	Facultad        int       `orm:"column(facultad)"`
	NivelAcademico  string    `orm:"column(nivel_academico)"`
	Dedicacion      string    `orm:"column(dedicacion)"`
	FechaExpedicion time.Time `orm:"column(fecha_expedicion);type(date)"`
	NumeroSemanas   int       `orm:"column(numero_semanas)"`
	Periodo         int       `orm:"column(periodo)"`
	TipoResolucion  string    `orm:"column(tipo_resolucion)"`
	FacultadNombre  string 
}


func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
}
