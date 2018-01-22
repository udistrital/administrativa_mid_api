package models

type DocentePrecontratado struct {
	Id                 int
	PrimerNombre       string
	SegundoNombre      string
	PrimerApellido     string
	SegundoApellido    string
	NombreCompleto     string
	Documento          int
	Expedicion         string
	Categoria          string
	Dedicacion         string
	HorasSemanales     int
	Semanas            int
	ProyectoCurricular int
	ValorContrato      int
}
