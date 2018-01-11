package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/administrativa_mid_api/models"
)

// ExpedirResolucionController operations for ExpedirResolucion
type ExpedirResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExpedirResolucionController) URLMapping() {
	c.Mapping("Expedir", c.Expedir)
}

// Expedir ...
// @Title Expedir
// @Description create Expedir
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedir [post]
func (c *ExpedirResolucionController) Expedir() {
	amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	amazon.Using("amazonAdmin")
	flyway.Using("flywayAdmin")
	var m models.ExpedicionResolucion
	var temp string
	var cdve int
	var proveedor []models.InformacionProveedor
	var disponibilidad models.Disponibilidad
	var dispoap models.DisponibilidadApropiacion
	var response interface{}
	vigencia, _, _ := time.Now().Date()
	//If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		//If 12 - Consecutivo contrato_general
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil {
			numeroContratos := cdve
			// for vinculaciones
			for _, vinculacion := range *v {
				numeroContratos = numeroContratos + 1
				v := vinculacion.VinculacionDocente
				idvinculaciondocente := strconv.Itoa(v.Id)
				//if 8 - Vinculacion_docente (GET)
				if err := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v); err == nil {
					//If 6 - Comprobación Contratos en vinculación docente
					if (v.NumeroContrato.String == "" && v.NumeroContrato.Valid == false) && (v.Vigencia.Int64 == 0 && v.Vigencia.Valid == false) {
						contrato := vinculacion.ContratoGeneral
						acta := vinculacion.ActaInicio
						aux1 := 181
						contrato.VigenciaContrato = vigencia
						contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
						contrato.FormaPago.Id = 240
						contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
						contrato.Justificacion = "Docente de Vinculacion Especial"
						contrato.UnidadEjecucion.Id = 205
						contrato.LugarEjecucion.Id = 4
						//contrato.Supervisor.Id = 192
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
						// If 5 - Informacion_Proveedor
						if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
							temp = proveedor[0].NumDocumento
							_, err = amazon.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante, contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp[0], contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, 192, contrato.ClaseContratista, contrato.TipoContrato, contrato.LugarEjecucion.Id).Exec()
							aux1 := contrato.Id
							aux2 := contrato.VigenciaContrato
							var ce models.ContratoEstado
							var ec models.EstadoContrato
							ce.NumeroContrato = aux1
							ce.Vigencia = aux2
							ce.FechaRegistro = time.Now()
							ec.Id = 4
							ce.Estado = &ec
							// If 4 - contrato_estado
							if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
								a := vinculacion.VinculacionDocente
								var ai models.ActaInicio
								ai.NumeroContrato = aux1
								ai.Vigencia = aux2
								ai.Descripcion = acta.Descripcion
								ai.FechaInicio = acta.FechaInicio
								ai.FechaFin = acta.FechaFin
								ai.FechaFin = CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)
								// If 3 - Acta_inicio
								if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response, &ai); err == nil {
									var cd models.ContratoDisponibilidad
									cd.NumeroContrato = aux1
									cd.Vigencia = aux2
									cd.Estado = true
									cd.FechaRegistro = time.Now()
									// If 2.5.2 - Get disponibildad_apropiacion
									if err := getJson("http://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/"+strconv.Itoa(v.Disponibilidad), &dispoap); err == nil {
										// If 2.5.1 - Get disponibildad
										if err := getJson("http://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil {
											cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
											cd.VigenciaCdp = int(disponibilidad.Vigencia)
											// If 2 - contrato_disponibilidad
											if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response, &cd); err == nil {
												a.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
												a.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
												v := a
												v.NumeroContrato.String = aux1
												v.NumeroContrato.Valid = true
												v.Vigencia.Int64 = int64(aux2)
												v.Vigencia.Valid = true
												// If 1 - vinculacion_docente
												if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.Id), "PUT", &response, &v); err == nil {
													fmt.Println("Vinculacion docente actualizada y lista, vamos por la otra")
												} else { // If 1 - vinculacion_docente
													fmt.Println("He fallado un poquito en If 1 - vinculacion_docente, solucioname!!! ", err)
													amazon.Rollback()
													flyway.Rollback()
													return
												}
											} else { // If 2 - contrato_disponibilidad
												fmt.Println("He fallado un poquito en  If 2 - contrato_disponibilidad, solucioname!!!", err)
												amazon.Rollback()
												flyway.Rollback()
												return
											}
										} else { // If 2.5.1 - Get disponibildad
											fmt.Println("He fallado un poquito en If 2.5.1 - Get disponibildad, solucioname!!!", err)
											amazon.Rollback()
											flyway.Rollback()
											return
										}
									} else { // If 2.5.2 - Get disponibildad_apropiacion
										fmt.Println("He fallado un poquito en If 2.5.2 - Get disponibildad_apropiacion, solucioname!!!", err)
										amazon.Rollback()
										flyway.Rollback()
										return
									}
								} else { // If 3 - Acta_inicio
									fmt.Println("He fallado un poquito en If 3 - Acta_inicio, solucioname!!!", err)
									amazon.Rollback()
									flyway.Rollback()
									return
								}
							} else { // If 4 - contrato_estado
								fmt.Println("He fallado un poquito en If 4 - contrato_estado, solucioname!!!", err)
								amazon.Rollback()
								flyway.Rollback()
								return
							}
						} else { // If 5 - Informacion_Proveedor
							fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
							amazon.Rollback()
							flyway.Rollback()
							return
						}

					} else { //If 6 - Comprobación Contratos en vinculación docente
						aux1 := v.NumeroContrato
						aux2 := v.Vigencia
						var ce models.ContratoEstado
						var ec models.EstadoContrato
						ce.NumeroContrato = aux1.String
						ce.Vigencia = int(aux2.Int64)
						ce.FechaRegistro = time.Now()
						ec.Id = 1
						ce.Estado = &ec
						// If 7 - Contrato_estado POST IF
						if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
							fmt.Println("Fin del proceso del for para esta vinculación, vamos bien: ", response)
						} else { // If 7 - Contrato_estado POST IF
							fmt.Println("He fallado un poquito en If 7 - Contrato_estado POST IF, solucioname!!!", err)
							amazon.Rollback()
							flyway.Rollback()
							return
						}
					}
				} else { //If 8 - Vinculacion_docente (GET)
					fmt.Println("He fallado un poquito en If 8 - Vinculacion_docente (GET), solucioname!!!", err)
					amazon.Rollback()
					flyway.Rollback()
					return
				}
			} // for vinculaciones
			var r models.Resolucion
			r.Id = m.IdResolucion
			idResolucionDVE := strconv.Itoa(m.IdResolucion)
			//If 11 - Resolucion (GET)
			if err := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
				fecha := time.Now()
				r.FechaExpedicion = fecha
				//If 10 - Resolucion (PUT)
				if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
					var e models.ResolucionEstado
					var er models.EstadoResolucion
					e.Resolucion = &r
					er.Id = 2
					e.Estado = &er
					e.FechaRegistro = time.Now()
					//If 9 - Resolucion_estado
					if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
						fmt.Println("Expedición exitosa, ahora va el commit :D")
						c.Data["json"] = v
					} else { //If 9 - Resolucion_estado
						fmt.Println("He fallado un poquito en If 9 - Resolucion_estado, solucioname!!!", err)
						amazon.Rollback()
						flyway.Rollback()
						return
					}
				} else { //If 10 - Resolucion (PUT)
					fmt.Println("He fallado un poquito en If 10 - Resolucion (PUT), solucioname!!! ", err)
					amazon.Rollback()
					flyway.Rollback()
					return
				}
			} else { //If 11 - Resolucion (GET)
				fmt.Println("He fallado un poquito en If 11 - Resolucion (GET), solucioname!!! ", err)
				amazon.Rollback()
				flyway.Rollback()
				return
			}
		} else { //If 12 - Consecutivo contrato_general
			fmt.Println("He fallado un poquito en If 12 - Consecutivo contrato_general, solucioname!!! ", err)
			amazon.Rollback()
			flyway.Rollback()
			return
		}

	} else { //If 13 - Unmarshal
		fmt.Println("He fallado un poquito en If 13 - Unmarshal, solucioname!!! ", err)
		amazon.Rollback()
		flyway.Rollback()
		return
	}
	amazon.Commit()
	flyway.Commit()
	c.ServeJSON()
}

func CalcularFechaFin(fecha_inicio time.Time, numero_semanas int) (fecha_fin time.Time) {
	semanas := float32(numero_semanas)
	meses := semanas / 4
	fmt.Println("meses", meses)
	numero_dias := (meses * 30) + 1
	f_i := fecha_inicio
	after := f_i.AddDate(0, 0, int(numero_dias))
	fmt.Println(after)
	return after
}