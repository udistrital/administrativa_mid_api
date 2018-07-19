package controllers

import (
	"encoding/json"
	"fmt"
	"math"
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
	c.Mapping("ExpedirModificacion", c.ExpedirModificacion)
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
									// If 3 - Acta_inicio creación
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
	var entero int
	var decimal float32
	meses := float32(numero_semanas) / 4
	entero = int(meses)
	decimal = meses - float32(entero)
	numero_dias := ((decimal * 4) * 7)
	f_i := fecha_inicio
	after := f_i.AddDate(0, entero, int(numero_dias))
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

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	v := m.Vinculaciones
	beego.Info(v)

	for _, vinculacion := range *v {
		v := vinculacion.VinculacionDocente
		idvinculaciondocente := strconv.Itoa(v.Id)

		err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+idvinculaciondocente, &v)
		if err != nil {
			beego.Error("Previnculación no valida", err)
			c.Abort("233")
		}

		contrato := vinculacion.ContratoGeneral
		var proveedor []models.InformacionProveedor

		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor)
		if err != nil {
			beego.Error("Docente no válido en Ágora, se encuentra identificado con el documento número ", strconv.Itoa(contrato.Contratista), err)
			c.Abort("233")
		}

		if proveedor == nil {
			beego.Error("No existe el docente con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			c.Abort("233")
		}

		var dispoap []models.DisponibilidadApropiacion

		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad_apropiacion/?query=Id:"+strconv.Itoa(v.Disponibilidad), &dispoap)
		if err != nil {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			c.Abort("233")
		}

		if dispoap == nil {
			beego.Error("Disponibilidad no válida asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			c.Abort("233")
		}

		var proycur []models.Dependencia

		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(v.IdProyectoCurricular), &proycur)
		if err != nil {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento "+strconv.Itoa(contrato.Contratista)+" en Ágora", err)
			c.Abort("233")
		}

		if proycur == nil {
			beego.Error("Dependencia incorrectamente homologada asociada al docente identificado con número de documento " + strconv.Itoa(contrato.Contratista) + " en Ágora")
			c.Abort("233")
		}
		beego.Info(proycur)

	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
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
						ai[0].FechaFin = CalcularFechaFin(ai[0].FechaInicio, v.NumeroSemanasNuevas)
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

// ExpedirModificacion ...
// @Title ExpedirModificacion
// @Description create ExpedirModificacion
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedirModificacion [post]
func (c *ExpedirResolucionController) ExpedirModificacion() {
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
	var modVin []models.ModificacionVinculacion
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
					fechaInicioOriginal := acta.FechaInicio
					acta.FechaInicio = time.Date(fechaInicioOriginal.Year(), fechaInicioOriginal.Month(), fechaInicioOriginal.Day(), 0, 0, 0, 0, fechaInicioOriginal.Location())
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

							//If modificacion_vinculacion
							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=VinculacionDocenteCancelada:"+strconv.Itoa(v.Id), &modVin); err == nil {
								var actaInicioAnterior []models.ActaInicio
								vinculacionModificacion := modVin[0].VinculacionDocenteRegistrada
								//If get acta_inicio cancelando
								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+modVin[0].VinculacionDocenteCancelada.NumeroContrato.String+",Vigencia:"+strconv.Itoa(int(modVin[0].VinculacionDocenteCancelada.Vigencia.Int64)), &actaInicioAnterior); err == nil {
									semanasTotales := vinculacion.VinculacionDocente.NumeroSemanasNuevas
									semanasIniciales := modVin[0].VinculacionDocenteCancelada.NumeroSemanas
									semanasModificar := modVin[0].VinculacionDocenteRegistrada.NumeroSemanas
									horasTotales := vinculacion.VinculacionDocente.NumeroHorasNuevas
									horasIniciales := modVin[0].VinculacionDocenteCancelada.NumeroHorasSemanales
									fechaFinNuevoContrato := CalcularFechaFin(acta.FechaInicio, semanasModificar)
									// Sólo si es reducción cambia la fecha fin del acta anterior y el valor del nuevo contrato
									if semanasTotales < semanasIniciales || horasTotales < horasIniciales {
										var aini models.ActaInicio
										aini.Id = actaInicioAnterior[0].Id
										aini.NumeroContrato = actaInicioAnterior[0].NumeroContrato
										aini.Vigencia = actaInicioAnterior[0].Vigencia
										aini.Descripcion = actaInicioAnterior[0].Descripcion
										aini.FechaInicio = actaInicioAnterior[0].FechaInicio
										aini.FechaFin = acta.FechaInicio
										fechaFinNuevoContrato = actaInicioAnterior[0].FechaFin
										// If put acta_inicio cancelando - cambia fecha fin del acta anterior por la fecha inicio escogida por el usuario
										if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/"+strconv.Itoa(aini.Id), "PUT", &response, &aini); err == nil {
											fmt.Println("Acta anterior cancelada en la fecha indicada")
										} else {
											fmt.Println("He fallado un poquito en If put acta_inicio cancelando, solucioname!!!", err)
											amazon.Rollback()
											flyway.Rollback()
											return
										}
										// Calcula el valor del nuevo contrato con base en las semanas desde la fecha inicio escogida hasta la nueva fecha fin y las nuevas horas
										semanasTranscurridasDecimal := (acta.FechaInicio.Sub(actaInicioAnterior[0].FechaInicio).Hours()) / 24 / 7
										semanasTranscurridas, decimal := math.Modf(semanasTranscurridasDecimal)
										if decimal > 0 {
											semanasTranscurridas = semanasTranscurridas + 1
										}
										var semanasTranscurridasInt = int(semanasTranscurridas)
										semanasRestantes := semanasIniciales - semanasTranscurridasInt - semanasModificar
										var vinc [1]models.VinculacionDocente
										vinc[0] = models.VinculacionDocente{
											IdResolucion:         &models.ResolucionVinculacionDocente{Id: m.IdResolucion},
											IdPersona:            v.IdPersona,
											NumeroHorasSemanales: horasTotales,
											NumeroSemanas:        semanasModificar,
											IdDedicacion:         v.IdDedicacion,
											IdProyectoCurricular: v.IdProyectoCurricular,
											Categoria:            v.Categoria,
											Dedicacion:           v.Dedicacion,
											NivelAcademico:       v.NivelAcademico,
											Vigencia:             v.Vigencia,
											Disponibilidad:       v.Disponibilidad,
										}
										salario, err := CalcularValorContratoReduccion(vinc, semanasRestantes, horasIniciales)
										if err != nil {
											fmt.Println("He fallado en cálculo del contrato reducción, solucioname!!!", err)
											amazon.Rollback()
											flyway.Rollback()
											return
										}
										contrato.ValorContrato = salario
										beego.Info(contrato.ValorContrato)
									}

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
											var ai models.ActaInicio
											ai.NumeroContrato = aux1
											ai.Vigencia = aux2
											ai.Descripcion = acta.Descripcion
											ai.FechaInicio = acta.FechaInicio
											ai.FechaFin = fechaFinNuevoContrato
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
															vinculacionModificacion.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
															vinculacionModificacion.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
															vinculacionModificacion.NumeroContrato.String = aux1
															vinculacionModificacion.NumeroContrato.Valid = true
															vinculacionModificacion.Vigencia.Int64 = int64(aux2)
															vinculacionModificacion.Vigencia.Valid = true
															// If 1 - vinculacion_docente
															if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(vinculacionModificacion.Id), "PUT", &response, &vinculacionModificacion); err == nil {
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
								} else { //If get acta_inicio cancelando
									fmt.Println("He fallado un poquito en If get acta_inicio cancelando, solucioname!!!", err)
									amazon.Rollback()
									flyway.Rollback()
									return
								}
							} else { //If modificacion_vinculacion
								fmt.Println("He fallado un poquito en If modificacion_vinculacion, solucioname!!!", err)
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

// Calcula el valor del contrato a reversar en dos partes:
// (1) las horas a reducir durante las semanas a reducir
// (2) las horas a originales en las semanas restantes (si quedan después de la reducción)
func CalcularValorContratoReduccion(v [1]models.VinculacionDocente, semanasRestantes int, horasOriginales int) (salarioTotal float64, err error) {
	var d []models.VinculacionDocente
	var salarioSemanasReducidas float64
	var salarioSemanasRestantes float64

	jsonEjemplo, err := json.Marshal(v)
	if err != nil {
		return salarioTotal, err
	}
	err = json.Unmarshal(jsonEjemplo, &d)
	if err != nil {
		return salarioTotal, err
	}

	docentes, err := CalcularSalarioPrecontratacion(d)
	if err != nil {
		return salarioTotal, err
	}
	salarioSemanasReducidas = docentes[0].ValorContrato

	if semanasRestantes > 0 {
		d[0].NumeroSemanas = semanasRestantes
		d[0].NumeroHorasSemanales = horasOriginales
		docentes, err := CalcularSalarioPrecontratacion(d)
		if err != nil {
			return salarioTotal, err
		}
		salarioSemanasRestantes = docentes[0].ValorContrato
	}

	salarioTotal = salarioSemanasReducidas + salarioSemanasRestantes
	return salarioTotal, nil
}
