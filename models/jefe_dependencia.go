package models

import (
	"time"
)

type JefeDependencia struct {
	Id                     int                     `orm:"column(id);pk;auto"`
  FechaInicio      time.Time                    `orm:"column(fecha_inicio)"`
  FechaFin      time.Time                    `orm:"column(fecha_fin)"`
  //Este atributo no debería ser llave foránea con terceros, diferentes esquemas.
	TerceroId                 int                     `orm:"column(tercero_id)"`
  //Cuando se una oikos al core: DependenciaId es foránea de dependencia(id).
	DependenciaId               int                 `orm:"column(dependencia_id)"`
	ActaAprobacion                 string                  `orm:"column(acta_aprobacion)"`
}
