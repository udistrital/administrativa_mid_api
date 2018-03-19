package models


type InformacionContratosPersona struct {
	ContratosPersonas struct {
		ContratoPersona []struct {
			TipoContrato  struct {
				Nombre string `json:"nombre"`
				Id int `json:"id"`
			} `json:"tipo_contrato"`
			Vigencia 		int `json:"vigencia"`
			NumeroContrato  int `json:"numero_contrato"`
			EstadoContrato  struct {
				Nombre string `json:"nombre"`
				Id int `json:"id"`
			} `json:"estado_contrato"`

		} `json:"contrato_persona"`
	} `json:"contratos_personas"`
}
