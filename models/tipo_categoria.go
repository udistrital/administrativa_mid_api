package models

type TipoCategoria struct {
	NombreCategoria string `orm:"column(nombre_categoria);null"`
	Id              int    `orm:"column(id);pk"`
}