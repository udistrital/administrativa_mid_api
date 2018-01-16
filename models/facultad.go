package models

type Facultad struct {
	Id             int              `orm:"column(id);pk"`
	Nombre         string            `orm:"column(nombre)"`
}
