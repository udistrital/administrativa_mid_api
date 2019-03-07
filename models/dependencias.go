package models

type DependenciasContrato struct {
	DependenciasSic struct {
		Dependencia []struct {
			CodigoDep    string `json:"ESFCODIGODEP"`
			NombreDep string `json:"ESFDEPENCARGADA"`
	    } `json:"Dependencia"`
	} `json:"DependenciasSic"`
} 


