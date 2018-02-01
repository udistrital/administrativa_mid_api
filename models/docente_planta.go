package models

type DocentePlanta struct {
	Nombres           string
	Apellidos         string
	Documento         string
	Id_tipo_documento string
	Tipo_documento    string
	Direccion         string
	Correo            string
	Telefono          string
	Celular           string
	Id_lugar          string
	Lugar             string
	Id_facultad       string
	Facultad          string
	Id_carrera        string
	Carrera           string
}

type ObjetoDocentePlanta struct {
	DocenteCollection struct {
		Docente []struct {
			Planta                  string `json:"planta"`
		} `json:"docentes"`
	} `json:"docentesCollection"`
}
