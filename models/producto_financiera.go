package models

import (
	"time"
)

//ProductoFinanciera modelo de los productos en financiera
type ProductoFinanciera struct {
	Id            int
	Nombre        string
	Descripcion   string
	FechaRegistro time.Time
}
