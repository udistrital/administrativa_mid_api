package models


type InformacionContratoContratista struct {
	InformacionContratista struct {
		Tiempo struct {
			Valor int `json:"valor"`
			Unidad string `json:"unidad"`
		} `json:"tiempo"`
		Supervisor struct {
			Nombre string `json:"nombre"`
			Cargo string `json:"cargo"`
		} `json:"supervisor"`
		Documento struct {
			Ciudad string `json:"ciudad"`
			Tipo string `json:"tipo"`
			Numero int `json:"numero"`
		} `json:"documento"`
		Dependencia string `json:"dependencia"`
		UnidadEjecutora struct {
			Nombre string `json:"nombre"`
			Descripcion string `json:"descripcion"`
		} `json:"unidad_ejecutora"`
		Cuenta struct {
			Banco string `json:"banco"`
			Tipo string `json:"tipo"`
			Numero int64 `json:"numero"`
		} `json:"cuenta"`
		ValorContrato int64 `json:"valor_contrato"`
		Contrato struct {
			Fecha string `json:"fecha"`
			Objeto string `json:"objeto"`
			Numero int `json:"numero"`
		}
		NombreCompleto string `json:"nombre_completo"`
	} `json:"informacion_contratista"`
}
