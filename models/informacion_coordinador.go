package models

type InformacionCoordinador struct {
	CarreraSniesCollection struct {
		CarreraSnies []struct {
			NombreCoordinador          string `json:"nombre_coordinador"`
			NumeroDocumentoCoordinador string `json:"numero_documento_coordinador"`
			NombreProyectoCurricular   string `json:"nombre_proyecto_curricular"`
		} `json:"carreraSnies"`
	} `json:"carreraSniesCollection"`
}
