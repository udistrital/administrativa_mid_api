package models


type Disponibilidad struct {
	Id                        int                          `orm:"auto;column(id);pk"`
	Vigencia                  float64                      `orm:"column(vigencia)"`
	NumeroDisponibilidad      float64                      `orm:"column(numero_disponibilidad);null"`

}
