package models

type Programa struct {
	Id             int          `orm:"column(id);pk"`
	NombrePrograma string       `orm:"column(nombre_programa)"`
	InstitucionId  *Institucion `orm:"column(institucion_id);rel(fk)"`
	NivelFormacion float64      `orm:"column(nivel_formacion)"`
}