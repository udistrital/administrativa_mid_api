package models

import "time"

type RegistroPresupuestal struct {
	Id                         int
	Vigencia                   float64
	FechaRegistro              time.Time
	Responsable                int
	Estado                     *EstadoRegistroPresupuestal
	NumeroRegistroPresupuestal int
	Beneficiario               int
	TipoCompromiso             *Compromiso
	NumeroCompromiso           int
	Solicitud                  int
}
