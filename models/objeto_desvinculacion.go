package models

type Objeto_Desvinculacion struct {
	IdModificacionResolucion int
	IdNuevaResolucion        int
	DisponibilidadNueva      int
	DocentesDesvincular      []VinculacionDocente
}
