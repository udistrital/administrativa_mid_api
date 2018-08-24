package models

type ContratoOrdenadorDependencia struct {
	ContratosOrdenadorDependencia struct {
		InformacionContratos []struct {
			Documento         string `json:"Documento"`
			NombreContratista string `json:"NombreContratista"`
			NumeroContrato    string `json:"NumeroContrato"`
			Vigencia 		  string `json:"Vigencia"`
		} `json:"informacion_contratos"`
	} `json:"contratos_ordenador_dependencia"`
}
