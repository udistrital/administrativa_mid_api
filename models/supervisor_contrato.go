package models

type SupervisorContrato struct {
	Id                    int
	Nombre                string
	Documento             int
	Cargo                 string
	SedeSupervisor        string
	DependenciaSupervisor string
	Tipo                  int
	Estado                bool
	DigitoVerificacion    int
}
