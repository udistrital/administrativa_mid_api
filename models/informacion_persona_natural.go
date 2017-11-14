package models

import (
	"time"
)

type InformacionPersonaNatural struct {
	TipoDocumento                     *ParametroEstandar `orm:"column(tipo_documento);rel(fk)"`
	Id                                string                `orm:"column(num_documento_persona);pk"`
	DigitoVerificacion                float64            `orm:"column(digito_verificacion)"`
	PrimerApellido                    string             `orm:"column(primer_apellido)"`
	SegundoApellido                   string             `orm:"column(segundo_apellido);null"`
	PrimerNombre                      string             `orm:"column(primer_nombre)"`
	SegundoNombre                     string             `orm:"column(segundo_nombre);null"`
	Cargo                             string             `orm:"column(cargo);null"`
	IdPaisNacimiento                  float64            `orm:"column(id_pais_nacimiento)"`
	Perfil                            *ParametroEstandar `orm:"column(perfil);rel(fk)"`
	Profesion                         string             `orm:"column(profesion);null"`
	Especialidad                      string             `orm:"column(especialidad);null"`
	MontoCapitalAutorizado            float64            `orm:"column(monto_capital_autorizado);null"`
	Genero                            string             `orm:"column(genero);null"`
	GrupoEtnico                       string             `orm:"column(grupo_etnico);null"`
	ComunidadLgbt                     bool               `orm:"column(comunidad_lgbt)"`
	CabezaFamilia                     bool               `orm:"column(cabeza_familia)"`
	PersonasACargo                    bool               `orm:"column(personas_a_cargo)"`
	NumeroPersonasACargo              float64            `orm:"column(numero_personas_a_cargo);null"`
	EstadoCivil                       string             `orm:"column(estado_civil)"`
	Discapacitado                     bool               `orm:"column(discapacitado)"`
	TipoDiscapacidad                  string             `orm:"column(tipo_discapacidad);null"`
	DeclaranteRenta                   bool               `orm:"column(declarante_renta)"`
	MedicinaPrepagada                 bool               `orm:"column(medicina_prepagada)"`
	ValorUvtPrepagada                 float64            `orm:"column(valor_uvt_prepagada);null"`
	CuentaAhorroAfc                   bool               `orm:"column(cuenta_ahorro_afc)"`
	NumCuentaBancariaAfc              string             `orm:"column(num_cuenta_bancaria_afc);null"`
	IdEntidadBancariaAfc              float64            `orm:"column(id_entidad_bancaria_afc);null"`
	InteresViviendaAfc                float64            `orm:"column(interes_vivienda_afc);null"`
	DependienteHijoMenorEdad          bool               `orm:"column(dependiente_hijo_menor_edad)"`
	DependienteHijoMenos23Estudiando  bool               `orm:"column(dependiente_hijo_menos23_estudiando)"`
	DependienteHijoMas23Discapacitado bool               `orm:"column(dependiente_hijo_mas23_discapacitado)"`
	DependienteConyuge                bool               `orm:"column(dependiente_conyuge)"`
	DependientePadreOHermano          bool               `orm:"column(dependiente_padre_o_hermano)"`
	IdNucleoBasico                    float64            `orm:"column(id_nucleo_basico);null"`
	IdArl                             int                `orm:"column(id_arl);null"`
	IdEps                             int                `orm:"column(id_eps);null"`
	IdFondoPension                    int                `orm:"column(id_fondo_pension);null"`
	IdCajaCompensacion                int                `orm:"column(id_caja_compensacion);null"`
	IdNitArl                          float64            `orm:"column(id_nit_arl);null"`
	IdNitEps                          float64            `orm:"column(id_nit_eps);null"`
	IdNitFondoPension                 float64            `orm:"column(id_nit_fondo_pension);null"`
	IdNitCajaCompensacion             float64            `orm:"column(id_nit_caja_compensacion);null"`
	FechaExpedicionDocumento          time.Time          `orm:"column(fecha_expedicion_documento);type(date)"`
	IdCiudadExpedicionDocumento       float64            `orm:"column(id_ciudad_expedicion_documento)"`
}
