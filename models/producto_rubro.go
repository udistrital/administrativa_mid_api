package models

import (
	"time"
)

//ProductoRubro modelo de la tabla de rompimiento que relaciona los productos y los rubros
type ProductoRubro struct {
	Id                int
	Rubro             *Rubro
	Producto          *ProductoFinanciera
	ValorDistribucion float64
	Activo            bool
	FechaRegistro     time.Time
}
