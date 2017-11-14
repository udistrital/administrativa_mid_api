package models

import (
	
)

type ObjetoCargaLectiva struct {
	CargasLectivas struct {
		CargaLectiva []struct {
			Anio                  string `json:"anio"`
			HorasLectivas	      string  `json:"horas_lectivas"`
			DocDocente            string `json:"doc_docente"`
			IDFacultad            string `json:"id_facultad"`
			IDProyecto            string `json:"id_proyecto"`
			IDTipoVinculacion     string `json:"id_tipo_vinculacion"`
			NombreFacultad        string `json:"nombre_facultad"`
			NombreProyecto        string `json:"nombre_proyecto"`
			NombreTipoVinculacion string `json:"nombre_tipo_vinculacion"`
			Periodo               string `json:"periodo"`
		} `json:"carga_lectiva"`
	} `json:"cargas_lectivas"`
}