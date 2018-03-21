package models


type InformacionContratosPersona struct {
	ContratosPersonas struct {
		ContratoPersona []struct {
			TipoContrato  struct {
				Nombre string `json:"nombre"`
				Id string `json:"id"`
			} `json:"tipo_contrato"`
			Vigencia 		string `json:"vigencia"`
			NumeroContrato  string `json:"numero_contrato"`
			EstadoContrato  struct {
				Nombre string `json:"nombre"`
				Id string `json:"id"`
			} `json:"estado_contrato"`

		} `json:"contrato_persona"`
	} `json:"contratos_personas"`
}
