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
	c.Mapping("ValidarDatosExpedicion", c.ValidarDatosExpedicion)
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
	var temp int
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
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil {
			numeroContratos := cdve
			// for vinculaciones
			for _, vinculacion := range *v {
				numeroContratos = numeroContratos + 1
				v := vinculacion.VinculacionDocente
				idvinculaciondocente := strconv.Itoa(v.Id)
				//if 8 - Vinculacion_docente (GET)
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v); err == nil {
					contrato := vinculacion.ContratoGeneral
					var sup models.SupervisorContrato
					acta := vinculacion.ActaInicio
					aux1 := 181
					contrato.VigenciaContrato = vigencia
					contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
					contrato.FormaPago.Id = 240
					contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
					contrato.Justificacion = "Docente de Vinculacion Especial"
					contrato.UnidadEjecucion.Id = 269
					contrato.LugarEjecucion.Id = 4
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
					sup.Id = SupervisorActual(v.IdResolucion.Id)
					contrato.Supervisor = &sup
					contrato.Condiciones = "Sin condiciones"
					// If 5 - Informacion_Proveedor
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
						if proveedor != nil { //Nuevo If
							temp = proveedor[0].Id
							_, err = amazon.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante, contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp, contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, contrato.Supervisor.Id, contrato.ClaseContratista, contrato.TipoContrato.Id, contrato.LugarEjecucion.Id).Exec()
							//If insert contrato_general
							if err == nil {
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
								if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
									a := vinculacion.VinculacionDocente
									var ai models.ActaInicio
									ai.NumeroContrato = aux1
									ai.Vigencia = aux2
									ai.Descripcion = acta.Descripcion
									ai.FechaInicio = acta.FechaInicio
									ai.FechaFin = acta.FechaFin
									ai.FechaFin = CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)
									// If 3 - Acta_inicio
									if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response, &ai); err == nil {
										var cd models.ContratoDisponibilidad
										cd.NumeroContrato = aux1
										cd.Vigencia = aux2
										cd.Estado = true
										cd.FechaRegistro = time.Now()
										// If 2.5.2 - Get disponibildad_apropiacion
										if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/"+strconv.Itoa(v.Disponibilidad), &dispoap); err == nil {
											// If 2.5.1 - Get disponibildad
											if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad/"+strconv.Itoa(dispoap.Disponibilidad.Id), &disponibilidad); err == nil {
												cd.NumeroCdp = int(disponibilidad.NumeroDisponibilidad)
												cd.VigenciaCdp = int(disponibilidad.Vigencia)
												// If 2 - contrato_disponibilidad
												if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response, &cd); err == nil {
													a.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
													a.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
													v := a
													v.NumeroContrato.String = aux1
													v.NumeroContrato.Valid = true
													v.Vigencia.Int64 = int64(aux2)
													v.Vigencia.Valid = true
													// If 1 - vinculacion_docente
													if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.Id), "PUT", &response, &v); err == nil {
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
							} else { //If insert contrato_general
								fmt.Println("He fallado un poquito en insert contrato_general, solucioname!!!", err)
								amazon.Rollback()
								flyway.Rollback()
								return
							}
						} else { // Nuevo If
							fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor nuevo, solucioname!!!", err)
							amazon.Rollback()
							flyway.Rollback()
							c.Ctx.Output.SetStatus(233)
							c.Ctx.Output.Body([]byte("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
							return
						}
					} else { // If 5 - Informacion_Proveedor
						fmt.Println("He fallado un poquito en If 5 - Informacion_Proveedor, solucioname!!!", err)
						amazon.Rollback()
						flyway.Rollback()
						return
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
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
				r.FechaExpedicion = m.FechaExpedicion
				//If 10 - Resolucion (PUT)
				if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
					var e models.ResolucionEstado
					var er models.EstadoResolucion
					e.Resolucion = &r
					er.Id = 2
					e.Estado = &er
					e.FechaRegistro = time.Now()
					//If 9 - Resolucion_estado
					if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
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
	numero_dias := (semanas * 7)
	f_i := fecha_inicio
	after := f_i.AddDate(0, 0, int(numero_dias))
	return after
}

func SupervisorActual(id_resolucion int) (id_supervisor_actual int) {
	var r models.Resolucion
	var j []models.JefeDependencia
	var s []models.SupervisorContrato
	var fecha = time.Now().Format("2006-01-02")
	//If Resolucion (GET)
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(id_resolucion), &r); err == nil {
		//If Jefe_dependencia (GET)
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=DependenciaId:"+strconv.Itoa(r.IdDependencia)+",FechaFin__gte:"+fecha+",FechaInicio__lte:"+fecha, &j); err == nil {
			//If Supervisor (GET)
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/supervisor_contrato/?query=Documento:"+strconv.Itoa(j[0].TerceroId)+",FechaFin__gte:"+fecha+",FechaInicio__lte:"+fecha+"&CargoId.Cargo__startswith:DECANO|VICE", &s); err == nil {
				return s[0].Id
			} else { //If Jefe_dependencia (GET)
				fmt.Println("He fallado un poquito en If Supervisor (GET) en el método SupervisorActual, solucioname!!! ", err)
				return 0
			}
		} else { //If Jefe_dependencia (GET)
			fmt.Println("He fallado un poquito en If Jefe_dependencia (GET) en el método SupervisorActual, solucioname!!! ", err)
			return 0
		}
	} else { //If Resolucion (GET)
		fmt.Println("He fallado un poquito en If Resolucion (GET) en el método SupervisorActual, solucioname!!! ", err)
		return 0
	}
	return 0
}

// ExpedirResolucionController ...
// @Title ValidarDatosExpedicion
// @Description create ValidarDatosExpedicion
// @Success 201 {int}
// @Failure 403 body is empty
// @router /validar_datos_expedicion [post]
func (c *ExpedirResolucionController) ValidarDatosExpedicion() {
	var m models.ExpedicionResolucion
	//If Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		// for vinculaciones
		for _, vinculacion := range *v {
			v := vinculacion.VinculacionDocente
			idvinculaciondocente := strconv.Itoa(v.Id)
			//if Vinculacion_docente (GET)
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v); err == nil {
				contrato := vinculacion.ContratoGeneral
				var proveedor []models.InformacionProveedor
				//If informacion_proveedor
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
					//If proveedor nulo
					if proveedor != nil {
						var dispoap []models.DisponibilidadApropiacion
						// If Get disponibildad_apropiacion
						if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/?query=Id:"+strconv.Itoa(v.Disponibilidad), &dispoap); err == nil {
							//If disponibilidad nula
							if dispoap != nil {
								var proycur []models.Dependencia
								// If Get Dependencia
								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(v.IdProyectoCurricular), &proycur); err == nil {
									//If dependencia nula
									if proycur != nil {
										c.Ctx.Output.SetStatus(201)
										c.Data["json"] = v
									} else { //If dependencia nula
										fmt.Println("dependencia nula")
										c.Ctx.Output.SetStatus(233)
										c.Ctx.Output.Body([]byte("Dependencia incorrectamente homologada asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
										return
									}
								} else { // If Get Dependencia
									fmt.Println("dependencia error ", err.Error())
									c.Ctx.Output.SetStatus(233)
									c.Ctx.Output.Body([]byte("Dependencia incorrectamente homologada asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
									return
								}
							} else { //If disponibilidad nula
								fmt.Println("disponibilidad nula")
								c.Ctx.Output.SetStatus(233)
								c.Ctx.Output.Body([]byte("Disponibilidad no válida asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
								return
							}
						} else { // If Get disponibildad_apropiacion
							fmt.Println("disponibilidad_apropiacion error ", err.Error())
							c.Ctx.Output.SetStatus(233)
							c.Ctx.Output.Body([]byte("Disponibilidad no válida asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
							return
						}
					} else { //If proveedor nulo
						fmt.Println("proveedor nulo")
						c.Ctx.Output.SetStatus(233)
						c.Ctx.Output.Body([]byte("No existe el docente con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora"))
						return
					}
				} else { //If informacion_proveedor
					fmt.Println("informacion_proveedor error")
					c.Ctx.Output.SetStatus(233)
					c.Ctx.Output.Body([]byte("Docente no válido en Ágora, se encuentra identificado con el documento número " + strconv.Itoa(contrato.Contratista) + " : " + err.Error()))
					return
				}
			} else { //if Vinculacion_docente (GET)
				fmt.Println("informacion_proveedor error")
				c.Ctx.Output.SetStatus(233)
				c.Ctx.Output.Body([]byte("Previnculación no válida : " + err.Error()))
				return
			}
		} //for vinculaciones
	} else { //If Unmarshal
		fmt.Println("Unmarshal error")
		c.Ctx.Output.SetStatus(233)
		c.Ctx.Output.Body([]byte("La resolución no es válida: " + err.Error()))
		return
	}
	c.ServeJSON()
}

// Cancelar ...
// @Title Cancelar
// @Description create Cancelar
// @Success 201 {int} models.ExpedicionCancelacion
// @Failure 403 body is empty
// @router /cancelar [post]
func (c *ExpedirResolucionController) Cancelar() {
	amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	amazon.Using("amazonAdmin")
	flyway.Using("flywayAdmin")
	var m models.ExpedicionCancelacion
	var response interface{}
	//var datosAnular models.DatosAnular
	var contratoCancelado models.ContratoCancelado
	//If 13 - Unmarshal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		// for vinculaciones
		for _, vinculacion := range *v {
			v := vinculacion.VinculacionDocente
			idVinculacionDocente := strconv.Itoa(v.Id)
			//If vinculacion_docente (get)
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idVinculacionDocente, &v); err == nil {
				contratoCancelado.NumeroContrato = v.NumeroContrato.String
				contratoCancelado.Vigencia = int(v.Vigencia.Int64)
				contratoCancelado.FechaCancelacion = vinculacion.ContratoCancelado.FechaCancelacion
				contratoCancelado.MotivoCancelacion = vinculacion.ContratoCancelado.MotivoCancelacion
				contratoCancelado.Usuario = vinculacion.ContratoCancelado.Usuario
				contratoCancelado.FechaRegistro = time.Now()
				contratoCancelado.Estado = vinculacion.ContratoCancelado.Estado
				// if contrato_cancelado (post)
				if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_cancelado", "POST", &response, &contratoCancelado); err == nil {
					var ai []models.ActaInicio
					// if acta_inicio (get)
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+contratoCancelado.NumeroContrato+",Vigencia:"+strconv.Itoa(contratoCancelado.Vigencia), &ai); err == nil {
						ai[0].FechaFin = contratoCancelado.FechaCancelacion
						// if acta_inicio (put)
						if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(ai[0].Id), "PUT", &response, &ai[0]); err == nil {
							var ce models.ContratoEstado
							var ec models.EstadoContrato
							ce.NumeroContrato = contratoCancelado.NumeroContrato
							ce.Vigencia = contratoCancelado.Vigencia
							ce.FechaRegistro = time.Now()
							ec.Id = 7
							ce.Estado = &ec
							// If contrato_estado (post)
							if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err == nil {
								var r models.Resolucion
								r.Id = m.IdResolucion
								idResolucionDVE := strconv.Itoa(m.IdResolucion)
								//If  Resolucion (GET)
								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucionDVE, &r); err == nil {
									r.FechaExpedicion = m.FechaExpedicion
									//If Resolucion (PUT)
									if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(r.Id), "PUT", &response, &r); err == nil {
										var e models.ResolucionEstado
										var er models.EstadoResolucion
										e.Resolucion = &r
										er.Id = 2
										e.Estado = &er
										e.FechaRegistro = time.Now()
										//If  Resolucion_estado (post)
										if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &response, &e); err == nil {
											fmt.Println("Expedición exitosa, ahora va el commit :D")
											c.Data["json"] = v
										} else { //If  Resolucion_estado (post)
											fmt.Println("He fallado un poquito en If  Resolucion_estado (post), solucioname!!! ", err)
											amazon.Rollback()
											flyway.Rollback()
											return
										}
									} else { //If Resolucion (PUT)
										fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
										amazon.Rollback()
										flyway.Rollback()
										return
									}
								} else { // If Resolucion (GET)
									fmt.Println("He fallado un poquito en If Resolucion (PUT), solucioname!!! ", err)
									amazon.Rollback()
									flyway.Rollback()
									return
								}
							} else { // If contrato_estado (post)
								fmt.Println("He fallado un poquito en If Resolucion (GET), solucioname!!! ", err)
								amazon.Rollback()
								flyway.Rollback()
								return
							}
						} else { // If acta_inicio (post)
							fmt.Println("He fallado un poquito en If Acta_Inicio (POST), solucioname!!! ", err)
							amazon.Rollback()
							flyway.Rollback()
							return
						}
					} else { // if acta_inicio (get)
						fmt.Println("He fallado un poquito en if acta_inicio (GET), solucioname!!! ", err)
						amazon.Rollback()
						flyway.Rollback()
						return
					}
				} else { // if contrato_cancelado (post)
					fmt.Println("He fallado un poquito en if contrato_cancelado (post), solucioname!!! ", err)
					amazon.Rollback()
					flyway.Rollback()
					return
				}
			} else {
				//If vinculacion_docente (get)
				fmt.Println("He fallado un poquito en If vinculacion_docente (get), solucioname!!! ", err)
				amazon.Rollback()
				flyway.Rollback()
				return
			}
		} // for vinculaciones

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
