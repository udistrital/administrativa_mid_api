package models


type InformacionContrato struct {
	Contrato struct {
		FechaSuscripcion string `json:"fecha_suscripcion"`
		Justificacion    string `json:"justificacion"`
		TipoContrato int `json:"tipo_contrato"`
		UnidadEjecucion int `json:"unidad_ejecucion"`
		Vigencia int `json:"vigencia"`
		OrdenadorGasto struct {
			Id int `json:"id"`
			RolOrdenador string `json:"rol_ordenador"`
			Nombre string `json:"nombre"`
		} `json:"ordenador_gasto"`

		DescripcionFormaPago string `json:"descripcion_forma_pago"`
		FechaRegistro string `json:"fecha_registro"`
		Observaciones string `json:"observaciones"`
		ObjetoContrato string `json:"objeto_contrato"`
		Contratista int `json:"contratista"`
		NumeroContratoSuscrito int `json:"numero_contrato_suscrito"`
		Supervisor struct {
			Nombre string `json:"nombre"`
			Id int `json:"id"`
			DocumentoIdentificacion int `json:"documento_identificacion"`
			Cargo string `json:"cargo"`
		} `json:"supervisor"`
		LugarEjecucion int `json:"lugar_ejecucion"`
		Actividades string `json:"actividades"`
		UnidadEjecutora int `json:"unidad_ejecutora"`
		NumeroContrato int `json:"numero_contrato"`
		PlazoEjecucion int `json:"plazo_ejecucion"`
		ValorContrato int64 `json:"valor_contrato"` 
	} `json:"contrato"`
}
