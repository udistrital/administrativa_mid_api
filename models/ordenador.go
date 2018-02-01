package models

import (
	"time"
)

type Ordenador struct {
	Id              int
	IdOrdenador     int
	InfoResolucion  string
	IdCiudad		    int
  FechaInicio     time.Time
  FechaFin        time.Time
  Estado          bool
  Documento       int
  NombreOrdenador string
  RolOrdenador    string

}
