package models

type TipoInvestigacion struct {
	Id                      int    `orm:"column(id);pk"`
	NombreTipoInvestigacion string `orm:"column(nombre_tipo_investigacion);null"`
}