package models



type EstadoPagoMensual struct {
	Id                int    
	Nombre            string  
	Descripcion       string  
	CodigoAbreviacion string  
	Activo            bool    
	NumeroOrden       float64 
}