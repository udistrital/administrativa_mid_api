package models

type ObjetoCategoriaDocente struct {
	CategoriaDocente struct {
		Anio           string `json:"anio"`
		Categoria      string `json:"categoria"`
		Identificacion string `json:"identificacion"`
		IDCategoria    string `json:"id_categoria"`
		Periodo        string `json:"periodo"`
	} `json:"categoria_docente"`
}
