package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ContratoGeneral struct {
	Id                           string              `orm:"column(numero_contrato);pk"`
	VigenciaContrato             int                 `orm:"column(vigencia)"`
	ObjetoContrato               string              `orm:"column(objeto_contrato);null"`
	PlazoEjecucion               int                 `orm:"column(plazo_ejecucion)"`
	FormaPago                    *Parametros         `orm:"column(forma_pago);rel(fk)"`
	OrdenadorGasto               int                 `orm:"column(ordenador_gasto)"`
	ClausulaRegistroPresupuestal bool                `orm:"column(clausula_registro_presupuestal);null"`
	SedeSolicitante              string              `orm:"column(sede_solicitante);null"`
	DependenciaSolicitante       string              `orm:"column(dependencia_solicitante);null"`
	Contratista                  int                 `orm:"column(contratista)"`
	ValorContrato                float64             `orm:"column(valor_contrato)"`
	Justificacion                string              `orm:"column(justificacion)"`
	DescripcionFormaPago         string              `orm:"column(descripcion_forma_pago)"`
	Condiciones                  string              `orm:"column(condiciones)"`
	FechaRegistro                time.Time           `orm:"column(fecha_registro);type(date)"`
	TipologiaContrato            int                 `orm:"column(tipologia_contrato)"`
	TipoCompromiso               int                 `orm:"column(tipo_compromiso)"`
	ModalidadSeleccion           int                 `orm:"column(modalidad_seleccion)"`
	Procedimiento                int                 `orm:"column(procedimiento)"`
	RegimenContratacion          int                 `orm:"column(regimen_contratacion)"`
	TipoGasto                    int                 `orm:"column(tipo_gasto)"`
	TemaGastoInversion           int                 `orm:"column(tema_gasto_inversion)"`
	OrigenPresupueso             int                 `orm:"column(origen_presupueso)"`
	OrigenRecursos               int                 `orm:"column(origen_recursos)"`
	TipoMoneda                   int                 `orm:"column(tipo_moneda)"`
	ValorContratoMe              float64             `orm:"column(valor_contrato_me);null"`
	ValorTasaCambio              float64             `orm:"column(valor_tasa_cambio);null"`
	TipoControl                  int                 `orm:"column(tipo_control);null"`
	Observaciones                string              `orm:"column(observaciones);null"`
	Supervisor                   *SupervisorContrato `orm:"column(supervisor);rel(fk)"`
	ClaseContratista             int                 `orm:"column(clase_contratista)"`
	Convenio                     string              `orm:"column(convenio);null"`
	NumeroConstancia             int                 `orm:"column(numero_constancia);null"`
	//RegistroPresupuestal         int                 `orm:"column(resgistro_presupuestal);null"`
	Estado          bool            `orm:"column(estado);null"`
	TipoContrato    *TipoContrato   `orm:"column(tipo_contrato);rel(fk)"`
	LugarEjecucion  *LugarEjecucion `orm:"column(lugar_ejecucion);rel(fk)"`
	UnidadEjecucion *Parametros     `orm:"column(unidad_ejecucion);rel(fk)"`
	UnidadEjecutora int             `orm:"column(unidad_ejecutora)"`
	//NumeroCdp                int             `orm:"column(numero_cdp)"`
	//NumeroSolicitudNecesidad int             `orm:"column(numero_solicitud_necesidad)"`
}

type TotalContratos struct {
	NumeroTotal int `orm:"column(total);null"`
}

type ContratoVinculacion struct {
	ContratoGeneral    *ContratoGeneral
	VinculacionDocente *VinculacionDocente
	ActaInicio         *ActaInicio
	Cdp                *ContratoDisponibilidad
}

type ContratoDisponibilidad struct {
	Id             int       `orm:"column(id);pk;auto"`
	NumeroCdp      int       `orm:"column(numero_cdp);null"`
	NumeroContrato string    `orm:"column(numero_contrato);null"`
	Vigencia       int       `orm:"column(vigencia);null"`
	Estado         bool      `orm:"column(estado);null"`
	FechaRegistro  time.Time `orm:"column(fecha_registro);type(date);null"`
	VigenciaCdp    int       `orm:"column(vigencia_cdp);null"`
}

type ExpedicionResolucion struct {
	Vinculaciones *[]ContratoVinculacion
	IdResolucion  int
}

type CancelarVinculacion struct {
	Vinculaciones *[]ContratoVinculacion
	IdResolucion  int
}

type EstadoAnulacion struct {
	Id                int    `orm:"column(id);pk;auto"`
	Nombre            string `orm:"column(nombre)"`
	Descripcion       string `orm:"column(descripcion);null"`
	Activo            bool   `orm:"column(activo)"`
	CodigoAbreviacion string `orm:"column(codigo_abreviacion);null"`
	NumeroOrden       string `orm:"column(numero_orden);null"`
}

type AnulacionRegistroPresupuestal struct {
	Id                   int              `orm:"auto;column(id);pk"`
	Consecutivo          int              `orm:"column(consecutivo)"`
	Motivo               string           `orm:"column(motivo)"`
	FechaRegistro        time.Time        `orm:"column(fecha_registro);type(date)"`
	TipoAnulacion        string           `orm:"column(tipo_anulacion)"`
	EstadoAnulacion      *EstadoAnulacion `orm:"column(estado_anulacion);rel(fk)"`
	JustificacionRechazo string           `orm:"column(justificacion_rechazo);null"`
	Responsable          int              `orm:"column(responsable)"`
	Solicitante          int              `orm:"column(solicitante)"`
	Expidio              int              `orm:"column(expidio)"`
}

type CategoriaCompromiso struct {
	Id          int    `orm:"column(id);pk;auto"`
	Nombre      string `orm:"column(nombre)"`
	Descripcion string `orm:"column(descripcion);null"`
}

type TipoCompromisoTesoral struct {
	Id                  int                  `orm:"column(id);pk;auto"`
	Nombre              string               `orm:"column(nombre)"`
	Activo              bool                 `orm:"column(activo);"`
	CategoriaCompromiso *CategoriaCompromiso `orm:"column(categoria_compromiso);rel(fk)"`
	Descripcion         string               `orm:"column(descripcion);null"`
	CodigoAbreviacion   string               `orm:"column(codigo_abreviacion);null"`
	NumeroOrden         float64              `orm:"column(numero_orden);null"`
}

type Compromiso struct {
	Id                    int                    `orm:"column(id);pk;auto"`
	Objeto                string                 `orm:"column(objeto)"`
	Vigencia              float64                `orm:"column(vigencia)"`
	FechaInicio           time.Time              `orm:"column(fecha_inicio);type(date)"`
	FechaFin              time.Time              `orm:"column(fecha_fin);type(date)"`
	FechaModificacion     time.Time              `orm:"column(fecha_modificacion);type(date)"`
	EstadoCompromiso      *EstadoCompromiso      `orm:"column(estado_compromiso);rel(fk)"`
	TipoCompromisoTesoral *TipoCompromisoTesoral `orm:"column(tipo_compromiso_financiero);rel(fk)"`
	UnidadEjecutora       int                    `orm:"column(unidad_ejecutora)"`
}

type EstadoCompromiso struct {
	Id                int     `orm:"column(id);pk;auto"`
	Nombre            string  `orm:"column(nombre)"`
	Descripcion       string  `orm:"column(descripcion);null"`
	CodigoAbreviacion string  `orm:"column(codigo_abreviacion);null"`
	Activo            bool    `orm:"column(activo);null"`
	NumeroOrden       float64 `orm:"column(numero_orden);null"`
}

type EstadoRegistroPresupuestal struct {
	Id                int    `orm:"column(id);pk"`
	Nombre            string `orm:"column(nombre)"`
	Descripcion       string `orm:"column(descripcion)"`
	Activo            bool   `orm:"column(activo)"`
	CodigoAbreviacion string `orm:"column(codigo_abreviacion);null"`
	NumeroOrden       string `orm:"column(numero_orden);null"`
}

type RegistroPresupuestal struct {
	Id                         int                         `orm:"column(id);pk;auto"`
	Vigencia                   float64                     `orm:"column(vigencia)"`
	FechaRegistro              time.Time                   `orm:"column(fecha_registro);type(date);null"`
	Responsable                int                         `orm:"column(responsable);null"`
	Estado                     *EstadoRegistroPresupuestal `orm:"column(estado);rel(fk)"`
	NumeroRegistroPresupuestal int                         `orm:"column(numero_registro_presupuestal)"`
	Beneficiario               int                         `orm:"column(beneficiario);null"`
	TipoCompromiso             *Compromiso                 `orm:"column(tipo_compromiso);rel(fk)"`
	NumeroCompromiso           int                         `orm:"column(numero_compromiso)"`
	Solicitud                  int                         `orm:"column(solicitud)"`
}

type RegistroPresupuestalDisponibilidadApropiacion struct {
	Id                        int                        `orm:"auto;column(id);pk"`
	RegistroPresupuestal      *RegistroPresupuestal      `orm:"column(registro_presupuestal);rel(fk)"`
	DisponibilidadApropiacion *DisponibilidadApropiacion `orm:"column(disponibilidad_apropiacion);rel(fk)"`
	Valor                     float64                    `orm:"column(valor);null"`
}

type DatosAnular struct {
	Anulacion      *AnulacionRegistroPresupuestal
	Rp_apropiacion *RegistroPresupuestalDisponibilidadApropiacion
	Valor          int
}

func GetNumeroTotalContratoGeneralDVE(vigencia int) (n int) {
	o := orm.NewOrm()
	var temp []TotalContratos
	_, err := o.Raw("SELECT COALESCE(MAX(substring(numero_contrato,4)::INTEGER), 0) total FROM argo.contrato_general WHERE numero_contrato LIKE 'DVE%' AND vigencia = " + strconv.Itoa(vigencia) + ";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return int(temp[0].NumeroTotal)
}

func (t *ContratoGeneral) TableName() string {
	return "contrato_general"
}

func init() {
	orm.RegisterModel(new(ContratoGeneral))
}

// AddContratoGeneral insert a new ContratoGeneral into database and returns
// last inserted Id on success.
func AddContratoGeneral(m *ContratoGeneral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratoGeneralById retrieves ContratoGeneral by Id. Returns error if
// Id doesn't exist
func GetContratoGeneralById(id string) (v *ContratoGeneral, err error) {
	o := orm.NewOrm()
	v = &ContratoGeneral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContratoGeneral retrieves all ContratoGeneral matches certain condition. Returns empty list if
// no records exist
func GetAllContratoGeneral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContratoGeneral))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ContratoGeneral
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateContratoGeneral updates ContratoGeneral by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratoGeneralById(m *ContratoGeneral) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContratoGeneral deletes ContratoGeneral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContratoGeneral(id string) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContratoGeneral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

/*

func AddContratosVinculcionEspecial(m ExpedicionResolucion) (alerta []string, err error) {
	o := orm.NewOrm()
	v := m.Vinculaciones
	var temp []int
	o.Begin()
	alerta = append(alerta, "success")
	vigencia, _, _ := time.Now().Date()
	numeroContratos := GetNumeroTotalContratoGeneralDVE(vigencia)
	for _, vinculacion := range v {
		numeroContratos = numeroContratos + 1
		v := vinculacion.VinculacionDocente

		//actaIni := vinculacion.ActaInicio
		if err = o.Read(&v); err == nil {
			if v.NumeroContrato == "" && v.Vigencia == 0 {
				contrato := vinculacion.ContratoGeneral
				acta := vinculacion.ActaInicio
				fmt.Println(contrato.Contratista)
				aux1 := 181
				contrato.VigenciaContrato = vigencia
				contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
				contrato.FormaPago = &Parametros{Id: 240}
				contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
				contrato.Justificacion = "Docente de Vinculacion Especial"
				contrato.UnidadEjecucion = &Parametros{Id: 205}
				contrato.LugarEjecucion = &LugarEjecucion{Id: 4}
				contrato.TipoControl = aux1
				contrato.ClaseContratista = 33
				contrato.TipoMoneda = 137
				contrato.OrigenRecursos = 149
				contrato.OrigenPresupueso = 156
				contrato.TemaGastoInversion = 166
				contrato.TipoGasto = 146
				contrato.RegimenContratacion = 136
				contrato.Procedimiento = 132
				contrato.ModalidadSeleccion = 123
				contrato.TipoCompromiso = 35
				contrato.TipologiaContrato = 46
				contrato.FechaRegistro = time.Now()
				contrato.UnidadEjecutora = 1
				contrato.Condiciones = "Sin condiciones"
				contrato.Supervisor = &SupervisorContrato{Id: 192}

				_, err1 := o.Raw("SELECT id_proveedor FROM agora.informacion_proveedor WHERE num_documento = CAST(" + strconv.Itoa(contrato.Contratista) + " as text);").QueryRows(&temp)
				if err1 == nil {
					fmt.Println("Consulta exitosa")
				}
				fmt.Println(contrato)

				_, err = o.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante,                                         contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp[0], contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, contrato.Supervisor.Id, contrato.ClaseContratista, contrato.TipoContrato, contrato.LugarEjecucion.Id).Exec()
				fmt.Println("Consulta realizada")
				if err == nil {
					aux1 := contrato.Id
					aux2 := contrato.VigenciaContrato
					ce := ContratoEstado{}
					ce.NumeroContrato = aux1
					ce.Vigencia = aux2
					ce.FechaRegistro = time.Now()
					ce.Estado = &EstadoContrato{Id: 4}

					if _, err = o.Insert(&ce); err != nil {
						o.Rollback()
						alerta[0] = "error"
						alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el estado del contrato!")
						return
					}
					if err == nil {
						a := vinculacion.VinculacionDocente
						ai := ActaInicio{}
						ai.NumeroContrato = aux1
						ai.Vigencia = aux2
						ai.Descripcion = acta.Descripcion
						ai.FechaInicio = acta.FechaInicio
						ai.FechaFin = acta.FechaFin
						ai.FechaFin = CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)
						if _, err = o.Insert(&ai); err != nil {
							o.Rollback()
							alerta[0] = "error"
							alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el acta inicio!")
							return
						} //AcA
						cd := ContratoDisponibilidad{}
						cd.NumeroCdp = vinculacion.Cdp.NumeroCdp
						cd.NumeroContrato = aux1
						cd.Vigencia = aux2
						cd.Estado = true
						cd.FechaRegistro = time.Now()
						cd.VigenciaCdp = vinculacion.Cdp.VigenciaCdp
						if _, err = o.Insert(&cd); err != nil {
							fmt.Println(err)
							o.Rollback()
							alerta[0] = "error"
							alerta = append(alerta, "Error: ¡Ocurrió un error al insertar en Contrato Disponibilidad!")
							return
						}

						//ACA
						if err == nil {
							if err = o.Read(&a); err == nil {
								a.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
								a.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
								v := a
								v.NumeroContrato = aux1
								v.Vigencia = aux2
								_, err = o.Update(&v)
								if err != nil {
									alerta[0] = "error"
									alerta = append(alerta, "Error: ¡Ocurrió un error al actualizar vinculación docente!")
									o.Rollback()
									return
								}
							} else {
								alerta[0] = "error"
								alerta = append(alerta, "Error: ¡Ocurrió un error al leer vinculacion docente!")
								o.Rollback()
								return
							}
						} //xD

						//Aca
					} else {
						alerta[0] = "error"
						alerta = append(alerta, "Error: ¡Ocurrió un error!")
						o.Rollback()
						return
					}
				} else {
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el contrato!")
					o.Rollback()
					return
				}
			} else {
				aux1 := v.NumeroContrato
				aux2 := v.Vigencia
				ce := ContratoEstado{}
				ce.NumeroContrato = aux1
				ce.Vigencia = aux2
				ce.FechaRegistro = time.Now()
				ce.Estado = &EstadoContrato{Id: 1}
				//_, err = o.Insert(&e)
				if err != nil {
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el estado del contrato!")
					o.Rollback()
					return
				}
			}
		} else {
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al leer vinculación docente!")
			o.Rollback()
			return
		}
	}
	r := &Resolucion{Id: m.IdResolucion}
	if err = o.Read(r); err == nil {
		fecha := time.Now()
		r.FechaExpedicion = fecha
		if _, err = o.Update(r); err == nil {
			var e ResolucionEstado
			e.Resolucion = r
			e.Estado = &EstadoResolucion{Id: 2}
			e.FechaRegistro = time.Now()
			_, err = o.Insert(&e)
			if err != nil {
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el estado de la resolución!")
				o.Rollback()
				return
			}
		} else {
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al actualizar el estado de la resolución!")
			o.Rollback()
			return
		}
	} else {
		alerta[0] = "error"
		alerta = append(alerta, "Error: ¡Ocurrió un error al leer la resolución!")
		o.Rollback()
		return
	}
	o.Commit()
	alerta = append(alerta, "Los datos de las vinculaciones realizadas han sido registradas con exito")
	return
}
*/
