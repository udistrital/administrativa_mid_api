package models

import (
	"time"
)

type VinculacionDocente struct {
	FechaRegistro        time.Time                  `orm:"column(fecha_registro);type(date)"`
	Estado               bool                       `orm:"column(estado)"`
	IdProyectoCurricular int                     `orm:"column(id_proyecto_curricular)"`
	IdDedicacion         *Dedicacion                `orm:"column(id_dedicacion);rel(fk)"`
	IdResolucion         *ResolucionVinculacionDocente `orm:"column(id_resolucion);rel(fk)"`
	IdSalarioMinimo      int                       `orm:"column(id_salario_minimo);null"`
	IdPuntoSalarial      int             			`orm:"column(id_punto_salarial);null"`
	NumeroSemanas        int                        `orm:"column(numero_semanas)"`
	NumeroHorasSemanales int                        `orm:"column(numero_horas_semanales)"`
	IdPersona            string 						`orm:"column(id_persona)"`
	Vigencia             int                       `orm:"column(vigencia);null"`
	NumeroContrato      	string                    `orm:"column(numero_contrato);null"`
	Id                   int                        `orm:"column(id);pk;auto"`
	NombreCompleto 		 string
	Categoria			 		 string
	Dedicacion				 string
	ValorContrato      float64
	NivelAcademico     string
	Disponibilidad 		   int
	NumeroDisponibilidad int
	LugarExpedicionCedula string
}

type Objeto_Desvinculacion struct{
	IdModificacionResolucion int
	DocentesDesvincular []VinculacionDocente
}

type ModificacionVinculacion struct {
	Id           								  int       			`orm:"column(id);pk;auto"`
	ModificacionResolucion        *ModificacionResolucion      `orm:"column(modificacion_resolucion);rel(fk)"`
	VinculacionDocenteCancelada     *VinculacionDocente    `orm:"column(vinculacion_docente_cancelada);rel(fk)"`
	VinculacionDocenteRegistrada    *VinculacionDocente    `orm:"column(vinculacion_docente_registrada);rel(fk);null"`
	Horas														int	   `orm:"column(horas);null"`
}
