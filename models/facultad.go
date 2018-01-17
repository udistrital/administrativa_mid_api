package models

type Facultad struct {
	Id             int              `orm:"column(id);pk"`
	Nombre         string            `orm:"column(nombre)"`
}

type ObjetoFacultad struct {
	Homologacion struct {
		IdOikos        string `json:"id_oikos"`
		IdGeDep        string `json:"id_gedep"`
	} `json:"homologacion"`
}
