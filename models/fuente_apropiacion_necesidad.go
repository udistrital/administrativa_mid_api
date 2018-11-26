package models

//FuenteApropiacionNecesidad modelo que unifica las fuentes y los productos con la apropiación a la que pertenecen según la necesidad asociada
type FuenteApropiacionNecesidad struct {
	Apropiacion ApropiacionRubro
	Fuentes     []FuenteFinanciacionRubroNecesidad
	Productos   []ProductoRubroNecesidad
	Monto       float64
}
