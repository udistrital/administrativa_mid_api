package models

//ProductoRubroNecesidad modelo de la aplicación necesidades que relaciona los productos con la necesidad asociada
type ProductoRubroNecesidad struct {
	Id                int
	ProductoRubro     int
	Apropiacion       int
	Necesidad         *Necesidad
	ProductoRubroInfo *[]ProductoRubro
}
