package models

type DatosAnular struct {
	Anulacion      *AnulacionRegistroPresupuestal
	Rp_apropiacion *RegistroPresupuestalDisponibilidadApropiacion
	Valor          int
}
