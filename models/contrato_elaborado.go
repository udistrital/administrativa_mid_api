package models

type ContratoElaborado struct {
	Contrato struct {
			Justificacion string `json:"justificacion"`
			TipoContrato int `json:"tipo_contrato"`
			UnidadEjecucion   int `json:"unidad_ejecucion"`
			Vigencia int `json:"vigencia"`
			DescripcionFormaPago string `json:"descripcion_forma_pago"`
			FechaRegistro string `json:"fecha_registro"`
			Observaciones string `json:"observaciones"`
			ObjetoContrato string `json:"objeto_contrato"`
			Contratista int `json:"contratista"`
			Supervisor struct {
				Id int `json:"id"`
				Nombre string `json:"nombre"`
				DocumentoIdentificacion int `json:"documento_identificacion"`
				Cargo string `json:"cargo"`
			} `json:"supervisor"`
			LugarEjecucion int `json:"lugar_ejecucion"`
			Actividades string `json:"actividades"`
			UnidadEjecutora int `json:"unidad_ejecutora"`
			NumeroContrato int `json:"numero_contrato"`
			PlazoEjecucion int `json:"plazo_ejecucion"`
			ValorContrato int64 `json:"valor_contrato"`
			OrdenadorGasto int `json:"ordenador_gasto"`
	} `json:"contrato"`
}

