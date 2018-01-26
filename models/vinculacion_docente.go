package models

import (
	"database/sql"
	"time"
)

type VinculacionDocente struct {
	FechaRegistro         time.Time
	Estado                bool
	IdProyectoCurricular  int
	IdDedicacion          *Dedicacion
	IdResolucion          *ResolucionVinculacionDocente
	IdSalarioMinimo       int
	IdPuntoSalarial       int
	NumeroSemanas         int
	NumeroHorasSemanales  int
	IdPersona             string
	Vigencia              sql.NullInt64
	NumeroContrato        sql.NullString
	Id                    int
	NombreCompleto        string
	Dedicacion            string
	NivelAcademico        string
	NumeroDisponibilidad  int
	ValorContrato         float64
	Categoria             string
	Disponibilidad        int
	LugarExpedicionCedula string
	NumeroHorasNuevas     int
	Periodo               int
	TipoDocumento 			  string
}
