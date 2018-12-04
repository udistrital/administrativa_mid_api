package models

import "time"

//Necesidad modelo completo de la necesidad
type Necesidad struct {
	Id                        int
	Numero                    int
	Vigencia                  float64
	Objeto                    string
	FechaSolicitud            time.Time
	Valor                     float64
	Justificacion             string
	UnidadEjecutora           int
	DiasDuracion              float64
	UnicoPago                 bool
	AgotarPresupuesto         bool
	ModalidadSeleccion        *ModalidadSeleccion
	TipoContratoNecesidad     *TipoContratoNecesidad
	PlanAnualAdquisiciones    int
	EstudioMercado            string
	TipoFinanciacionNecesidad *TipoFinanciacionNecesidad
	Supervisor                int
	AnalisisRiesgo            string
	NumeroElaboracion         int
	FechaModificacion         time.Time
	EstadoNecesidad           *EstadoNecesidad
	JustificacionRechazo      string
	JustificacionAnulacion    string
	TipoNecesidad             *TipoNecesidad
}
