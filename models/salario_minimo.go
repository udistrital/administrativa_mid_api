package models

type SalarioMinimo struct {
	Id       int     `orm:"column(id);pk"`
	Vigencia int     `orm:"column(vigencia)"`
	Valor    int	 `orm:"column(valor)"`
	Decreto  string  `orm:"column(decreto);null"`
}
