package models

type ContratoElaborado struct {
	Contrato struct {
			Justificacion string `json:"justificacion"`
			TipoContrato string `json:"tipo_contrato"`
			UnidadEjecucion   string `json:"unidad_ejecucion"`
			Vigencia string `json:"vigencia"`
			DescripcionFormaPago string `json:"descripcion_forma_pago"`
			FechaRegistro string `json:"fecha_registro"`
			Observaciones string `json:"observaciones"`
			ObjetoContrato string `json:"objeto_contrato"`
			Contratista string `json:"contratista"`
			Supervisor struct {
				Id string `json:"id"`
				Nombre string `json:"nombre"`
				DocumentoIdentificacion string `json:"documento_identificacion"`
				Cargo string `json:"cargo"`
			} `json:"supervisor"`
			LugarEjecucion string `json:"lugar_ejecucion"`
			Actividades string `json:"actividades"`
			UnidadEjecutora string `json:"unidad_ejecutora"`
			NumeroContrato string `json:"numero_contrato"`
			PlazoEjecucion string `json:"plazo_ejecucion"`
			ValorContrato string `json:"valor_contrato"`
	} `json:"contrato"`
}

