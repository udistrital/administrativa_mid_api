package models


type ContratoDependencia struct {
	Contratos struct {
		Contrato []struct {
				Vigencia string `json:"vigencia"`
				NumeroContrato string `json:"numero_contrato"`
		} `json:"contrato"`
	} `json:"contratos"`
}
