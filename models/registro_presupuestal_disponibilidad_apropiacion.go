package models

type RegistroPresupuestalDisponibilidadApropiacion struct {
	Id                        int
	RegistroPresupuestal      *RegistroPresupuestal
	DisponibilidadApropiacion *DisponibilidadApropiacion
	Valor                     float64
}
