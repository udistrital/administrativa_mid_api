package models

import (

)

type ObjetoCargaLectiva struct {
	CargasLectivas struct {
		CargaLectiva []struct {
			Anio                  string `json:"anio"`
			HorasLectivas	      string  `json:"horas_lectivas"`
			DocDocente            string `json:"docente_documento"`
			IDFacultad            string `json:"id_facultad"`
			IDProyecto            string `json:"id_proyecto"`
			IDTipoVinculacion     string `json:"id_tipo_vinculacion"`
			NombreFacultad        string `json:"facultad_nombre"`
			NombreProyecto        string `json:"proyecto_nombre"`
			NombreTipoVinculacion string `json:"tipo_vinculacion_nombre"`
			Periodo               string `json:"periodo"`
			DocenteApellido       string `json:"docente_apellido"`
			DocenteNombre       string `json:"docente_nombre"`
			CategoriaNombre 		string
			IDCategoria					string
			IdProveedor					string
		} `json:"carga_lectiva"`
	} `json:"cargas_lectivas"`
}

type ObjetoCategoriaDocente struct {
	CategoriaDocente struct {
		Anio           string `json:"anio"`
		Categoria      string `json:"categoria"`
		Identificacion string `json:"identificacion"`
		IDCategoria    string `json:"id_categoria"`
		Periodo        string `json:"periodo"`
	} `json:"categoria_docente"`
}

type Homologacion struct {
	Old     string
	New 		string
}

type HomologacionDedicacion struct {
	Nombre  string
	Old     string
	New 		string
}
