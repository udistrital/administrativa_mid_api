package models

type CambioEstadoPago struct {
	Id                     int          
	FechaCreacion          string    	
	FechaModificacion      string    	
	EstadoPagoMensualId    int         
	DocumentoResponsableId string      
	CargoResponsable       string      
	Activo                 bool        
	PagoMensualId          PagoMensual
}