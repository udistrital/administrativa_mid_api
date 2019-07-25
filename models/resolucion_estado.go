package models

type ResolucionEstado struct {
	Id            int
	FechaRegistro string
	Usuario       string
	Estado        *EstadoResolucion
	Resolucion    *Resolucion
}
