package models

type Objeto_Desvinculacion struct {
	IdModificacionResolucion int
	IdNuevaResolucion        int
	DisponibilidadNueva      int
	TipoDesvinculacion       string
	DocentesDesvincular      []VinculacionDocente
}
