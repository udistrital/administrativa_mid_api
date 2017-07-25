package models

import (
	"time"
)

type Institucion struct {
	Id                int       `orm:"column(id);pk"`
	NombreInstitucion string    `orm:"column(nombre_institucion)"`
	Pais              string    `orm:"column(pais)"`
	Departamento      string    `orm:"column(departamento)"`
	Municipio         string    `orm:"column(municipio)"`
	Localidad         string    `orm:"column(localidad);null"`
	FechaDato         time.Time `orm:"column(fecha_dato);type(date);null"`
	Vigente           bool      `orm:"column(vigente);null"`
}