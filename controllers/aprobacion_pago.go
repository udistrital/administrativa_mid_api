package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	//"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

// AprobacionPagoController operations for AprobacionPago
type AprobacionPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *AprobacionPagoController) URLMapping() {
	c.Mapping("ObtenerInfoCoordinador", c.ObtenerInfoCoordinador)
	c.Mapping("GetContratosDocente", c.GetContratosDocente)
	c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
	c.Mapping("PagoAprobado", c.PagoAprobado)
	c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
	c.Mapping("CertificacionDocumentosAprobados", c.CertificacionDocumentosAprobados)
	c.Mapping("ObtenerDependenciaOrdenador", c.ObtenerDependenciaOrdenador)

}

// AprobacionPagoController ...
// @Title ObtenerInfoCoordinador
// @Description create ObtenerInfoCoordinador
// @Param id_dependencia_oikos query int true "Proyecto a obtener información coordinador"
// @Success 201 {int} models.InformacionCoordinador
// @Failure 403 :id_dependencia_oikos is empty
// @router /informacion_coordinador/:id_dependencia_oikos [get]
func (c *AprobacionPagoController) ObtenerInfoCoordinador() {
	id_oikos := c.GetString(":id_dependencia_oikos")
	var temp map[string]interface{}
	var temp_snies map[string]interface{}
	var info_coordinador models.InformacionCoordinador

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+"proyecto_curricular_oikos/"+id_oikos, &temp); err == nil && temp != nil {
		json_proyecto_curricular, error_json := json.Marshal(temp)

		if error_json == nil {
			var temp_homologacion models.ObjetoProyectoCurricular
			json.Unmarshal(json_proyecto_curricular, &temp_homologacion)
			id_proyecto_snies := temp_homologacion.Homologacion.IDSnies

			if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAcademica")+"/"+"carrera_snies/"+id_proyecto_snies, &temp_snies); err == nil && temp_snies != nil {
				json_info_coordinador, error_json := json.Marshal(temp_snies)

				if error_json == nil {
					var temp_info_coordinador models.InformacionCoordinador
					json.Unmarshal(json_info_coordinador, &temp_info_coordinador)
					fmt.Println(temp_info_coordinador)
					info_coordinador = temp_info_coordinador
				} else {
					fmt.Println(error_json.Error())
				}
			}

		} else {
			fmt.Println(error_json.Error())
		}
	} else {
		fmt.Println(err)

	}

	c.Data["json"] = info_coordinador
	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title GetContratosDocente
// @Description create  GetContratosDocente
// @Param numDocumento query string true "Docente a consultar"
// @Success 201 {object} []models.ContratosDocentes
// @Failure 403 body is empty
// @router /get_contratos_docente/:numDocumento [get]
func (c *AprobacionPagoController) GetContratosDocente() {
	numDocumento := c.GetString(":numDocumento")
	var contratosDocentes []models.ContratosDocente
	var cd models.ContratosDocente
	var proveedor []models.InformacionProveedor
	var vinculaciones []models.VinculacionDocente
	var actasInicio []models.ActaInicio
	var res models.Resolucion
	var dep models.Dependencia
	//If informacion_proveedor get
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=num_documento:"+numDocumento, &proveedor); err == nil {
		//If vinculacion_docente get
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:TRUE,IdPersona:"+numDocumento+"&limit=-1", &vinculaciones); err == nil {
			//for vinculaciones
			for _, vinculacion := range vinculaciones {
				//If dependencia get
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {
					//If resolucion get
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(vinculacion.IdResolucion.Id), &res); err == nil {
						//If nulo

						if vinculacion.NumeroContrato.Valid == true {
							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+vinculacion.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion.Vigencia.Int64, 10), &actasInicio); err == nil {
	
								//If Estado = 4
								for _, actaInicio := range actasInicio {
									actaInicio.FechaInicio = actaInicio.FechaInicio.UTC()
									actaInicio.FechaFin = actaInicio.FechaFin.UTC()
									if int(actaInicio.FechaInicio.Month()) <= int(time.Now().Month()) && actaInicio.FechaInicio.Year() <= time.Now().Year() && int(actaInicio.FechaFin.Month()) >= int(time.Now().Month()) && actaInicio.FechaFin.Year() <= time.Now().Year() {
										cd.NumeroVinculacion = vinculacion.NumeroContrato.String
										cd.Vigencia = vinculacion.Vigencia.Int64
										cd.Resolucion = res.NumeroResolucion
										cd.Dependencia = dep.Nombre
										cd.IdDependencia = dep.Id
										cd.NombreDocente = proveedor[0].NomProveedor
										cd.Dedicacion = vinculacion.IdDedicacion.NombreDedicacion
										contratosDocentes = append(contratosDocentes, cd)
									}
								}
							} else { //If contrato_estado get
								fmt.Println("Mirenme, me morí en If contrato_estado get, solucioname!!! ", err)
								return
							}
						}
					} else { //If resolucion get
						fmt.Println("Mirenme, me morí en If resolucion get, solucioname!!! ", err)
						return
					}
				} else { //If dependencia get
					fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
					return
				}
			} //for vinculaciones
			c.Data["json"] = contratosDocentes
		} else { //If informacion_proveedor get
			fmt.Println("Mirenme, me morí en informacion proveedor, solucioname!!! ", err)
			return
		}
	} else { //If informacion_proveedor get
		fmt.Println("Mirenme, me morí en informacion proveedor, solucioname!!! ", err)
		return
	}
	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title ObtenerInfoOrdenador
// @Description create ObtenerInfoOrdenador
// @Param numero_contrato query int true "Numero de contrato en la tabla contrato general"
// @Param vigencia query int true "Vigencia del contrato en la tabla contrato general"
// @Success 201 {int} models.InformacionOrdenador
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @router /informacion_ordenador/:numero_contrato/:vigencia [get]
func (c *AprobacionPagoController) ObtenerInfoOrdenador() {
	numero_contrato := c.GetString(":numero_contrato")
	vigencia := c.GetString(":vigencia")

	var temp map[string]interface{}
	var contrato_elaborado models.ContratoElaborado
	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia
	var informacion_proveedores []models.InformacionProveedor
	var informacion_ordenador models.InformacionOrdenador
	var ordenadores []models.Ordenador

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contrato_elaborado/"+numero_contrato+"/"+vigencia, &temp); err == nil && temp != nil {
		json_contrato_elaborado, error_json := json.Marshal(temp)

		if error_json == nil {
			json.Unmarshal(json_contrato_elaborado, &contrato_elaborado)
			if contrato_elaborado.Contrato.TipoContrato == "2" || contrato_elaborado.Contrato.TipoContrato == "3" || contrato_elaborado.Contrato.TipoContrato == "18" {
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=Id:"+contrato_elaborado.Contrato.OrdenadorGasto, &ordenadores_gasto); err == nil {

					for _, ordenador_gasto := range ordenadores_gasto {

						if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=DependenciaId:"+strconv.Itoa(ordenador_gasto.DependenciaId)+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); err == nil {

							for _, jefe_dependencia := range jefes_dependencia {

								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(jefe_dependencia.TerceroId), &informacion_proveedores); err == nil {

									for _, informacion_proveedor := range informacion_proveedores {

										informacion_ordenador.NumeroDocumento = jefe_dependencia.TerceroId
										informacion_ordenador.Cargo = ordenador_gasto.Cargo
										informacion_ordenador.Nombre = informacion_proveedor.NomProveedor
										informacion_ordenador.IdDependencia = jefe_dependencia.DependenciaId
										c.Data["json"] = informacion_ordenador

									}

								} else {

									fmt.Println(err)
								}

							}

						} else {
							fmt.Println(err)
						}

					}

				} else {
					fmt.Println(err)
				}

			} else { //si no son docentes
				fmt.Println(contrato_elaborado.Contrato.OrdenadorGasto)
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/ordenadores/?query=IdOrdenador:"+contrato_elaborado.Contrato.OrdenadorGasto+"&sortby=FechaInicio&order=desc&limit=1", &ordenadores); err == nil {
					for _, ordenador := range ordenadores {
						informacion_ordenador.NumeroDocumento = ordenador.Documento
						informacion_ordenador.Cargo = ordenador.RolOrdenador
						informacion_ordenador.Nombre = ordenador.NombreOrdenador
						c.Data["json"] = informacion_ordenador

					}

				} else {

					fmt.Println(err)

				}

			}
		} else {
			fmt.Println(error_json.Error())
			return
		}
	} else {
		fmt.Println(err)

	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title PagoAprobado
// @Description create PagoAprobado
// @Param numero_contrato query int true "Numero de contrato en la tabla contrato general"
// @Param vigencia query int true "Vigencia del contrato en la tabla contrato general"
// @Param mes query int true "Mes del pago mensual"
// @Param anio query int true "Año del pago mensual"
// @Success 201
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /pago_aprobado/:numero_contrato/:vigencia/:mes/:anio [get]
func (c *AprobacionPagoController) PagoAprobado() {
	numero_contrato := c.GetString(":numero_contrato")
	vigencia := c.GetString(":vigencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	var pagos_mensuales []models.PagoMensual

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=NumeroContrato:"+numero_contrato+",VigenciaContrato:"+vigencia+",Mes:"+mes+",Ano:"+anio, &pagos_mensuales); err == nil {

		if pagos_mensuales != nil {

			for _, pago_mensual := range pagos_mensuales {

				if pago_mensual.EstadoPagoMensual.CodigoAbreviacion == "AP" {

					c.Data["json"] = "True"
				} else {

					c.Data["json"] = "False"
				}

			}
		} else {
			c.Data["json"] = "False"
		}

	} else { //If pago_mensual get
		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title CertificacionVistoBueno
// @Description create CertificacionVistoBueno
// @Param dependencia query int true "Dependencia del contrato en la tabla vinculacion"
// @Param mes query int true "Mes del pago mensual"
// @Param anio query int true "Año del pago mensual"
// @Success 201
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /certificacion_visto_bueno/:dependencia/:mes/:anio [get]
func (c *AprobacionPagoController) CertificacionVistoBueno() {
	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")
	var vinculaciones_docente []models.VinculacionDocente
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var personas []models.Persona
	var persona models.Persona
	var actasInicio []models.ActaInicio
	var mes_cer, _ = strconv.Atoi(mes)
	var anio_cer, _ = strconv.Atoi(anio)

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=Estado:TRUE,IdProyectoCurricular:"+dependencia, &vinculaciones_docente); err == nil {

		for _, vinculacion_docente := range vinculaciones_docente {
			if vinculacion_docente.NumeroContrato.Valid == true {

				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10), &actasInicio); err == nil {

					for _, actaInicio := range actasInicio {
						//If Estado = 4
						if int(actaInicio.FechaInicio.Month()) <= mes_cer && actaInicio.FechaInicio.Year() <= anio_cer && int(actaInicio.FechaFin.Month()) >= mes_cer && actaInicio.FechaFin.Year() >= anio_cer {

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=EstadoPagoMensual.CodigoAbreviacion.in:PAD|AD|AP,NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",VigenciaContrato:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10)+",Mes:"+mes+",Ano:"+anio, &pagos_mensuales); err == nil {

								if pagos_mensuales == nil {

									if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+vinculacion_docente.IdPersona, &contratistas); err == nil {

										for _, contratista := range contratistas {

											persona.NumDocumento = contratista.NumDocumento
											persona.Nombre = contratista.NomProveedor
											persona.NumeroContrato = actaInicio.NumeroContrato
											persona.Vigencia = actaInicio.Vigencia

											personas = append(personas, persona)

										}

									} else { //If informacion_proveedor get

										fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
									}

								}

							} else { //If pago_mensual get
								fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
								return
							}
						}
					}
				} else { //If contrato_estado get
					fmt.Println("Mirenme, me morí en If contrato_estado get, solucioname!!! ", err)
					return
				}
			}
		}

	} else { //If vinculacion_docente get

		fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
	}
	c.Data["json"] = personas

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title CertificacionDocumentosAprobados
// @Description create CertificacionDocumentosAprobados
// @Param dependencia query int true "Dependencia del contrato en la tabla ordenador_gasto"
// @Param mes query int true "Mes del pago mensual"
// @Param anio query int true "Año del pago mensual"
// @Success 201
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /certificacion_documentos_aprobados/:dependencia/:mes/:anio [get]
func (c *AprobacionPagoController) CertificacionDocumentosAprobados() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	var ordenadores_gasto []models.OrdenadorGasto
	var contratos_generales []models.ContratoGeneral
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var personas []models.Persona
	var persona models.Persona
	var actasInicio []models.ActaInicio
	var vinculaciones_docente []models.VinculacionDocente

	var mes_cer, _ = strconv.Atoi(mes)
	var anio_cer, _ = strconv.Atoi(anio)
	

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=DependenciaId:"+dependencia, &ordenadores_gasto); err == nil {

		for _, ordenador_gasto := range ordenadores_gasto {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/?limit=-1&query=OrdenadorGasto:"+strconv.Itoa(ordenador_gasto.Id), &contratos_generales); err == nil {

				for _, contrato_general := range contratos_generales {

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=Estado:TRUE,NumeroContrato:"+contrato_general.Id+",Vigencia:"+strconv.Itoa(contrato_general.VigenciaContrato), &vinculaciones_docente); err == nil {

						for _, vinculacion_docente := range vinculaciones_docente {
							if vinculacion_docente.NumeroContrato.Valid == true {
								
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+contrato_general.Id+",Vigencia:"+strconv.Itoa(contrato_general.VigenciaContrato), &actasInicio); err == nil {

						for _, actaInicio := range actasInicio {

							if int(actaInicio.FechaInicio.Month()) <= mes_cer && actaInicio.FechaInicio.Year() <= anio_cer && int(actaInicio.FechaFin.Month()) >= mes_cer && actaInicio.FechaFin.Year() >= anio_cer {
								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=EstadoPagoMensual.CodigoAbreviacion:AP,NumeroContrato:"+contrato_general.Id+",VigenciaContrato:"+strconv.Itoa(contrato_general.VigenciaContrato)+",Mes:"+mes+",Ano:"+anio, &pagos_mensuales); err == nil {

									if pagos_mensuales == nil {

										if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=Id:"+strconv.Itoa(contrato_general.Contratista), &contratistas); err == nil {

											for _, contratista := range contratistas {

												persona.NumDocumento = contratista.NumDocumento
												persona.Nombre = contratista.NomProveedor
												persona.NumeroContrato = contrato_general.Id
												persona.Vigencia = contrato_general.VigenciaContrato

												personas = append(personas, persona)

											}

										} else { //If informacion_proveedor get

											fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
										}

									}

								} else { //If pago_mensual get
									fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
								}

							}

						}
					} else { //If contrato_estado get

						fmt.Println("Mirenme, me morí en If contrato_estado get, solucioname!!! ", err)
					}
				}
				}
					}else { //If vinculacion_docente get

						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
					}
				}

			} else { //If contrato_general get

				fmt.Println("Mirenme, me morí en If contrato_general get, solucioname!!! ", err)
			}

		}

	} else { //If ordenador_gasto get

		fmt.Println("Mirenme, me morí en If ordenador_gasto get, solucioname!!! ", err)

	}
	c.Data["json"] = personas
	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title GetSolicitudesSupervisor
// @Description create GetSolicitudesSupervisor
// @Param docsupervisor query string true "Número del documento del supervisor"
// @Success 201
// @Failure 403 :docsupervisor is empty
// @router /solicitudes_supervisor/:docsupervisor [get]
func (c *AprobacionPagoController) GetSolicitudesSupervisor() {

	doc_supervisor := c.GetString(":docsupervisor")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_personas_proyecto []models.PagoPersonaProyecto

	var vinculaciones_docente []models.VinculacionDocente
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?limit=-1&query=EstadoPagoMensual.CodigoAbreviacion:PAD,Responsable:"+doc_supervisor, &pagos_mensuales); err == nil {

		for x, pago_mensual := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=Estado:TRUE,NumeroContrato:"+pago_mensual.NumeroContrato+",Vigencia:"+strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {

						for _, vinculacion := range vinculaciones_docente {
							var dep models.Dependencia
							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {
								var pago_personas_proyecto models.PagoPersonaProyecto

								pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
								pago_personas_proyecto.NombrePersona = contratista.NomProveedor
								pago_personas_proyecto.Dependencia = &dep
								pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)

							} else { //If dependencia get

								fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
								return

							}

						}

					} else { // If vinculacion_docente_get

						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
						return
					}
				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}

			c.Data["json"] = pagos_personas_proyecto
		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title GetSolicitudesCoordinador
// @Description create GetSolicitudesCoordinador
// @Param doccoordinador query string true "Número del documento del coordinador"
// @Success 201
// @Failure 403 :doccoordinador is empty
// @router /solicitudes_coordinador/:doccoordinador [get]
func (c *AprobacionPagoController) GetSolicitudesCoordinador() {

	doc_coordinador := c.GetString(":doccoordinador")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_personas_proyecto []models.PagoPersonaProyecto
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?limit=-1&query=EstadoPagoMensual.CodigoAbreviacion:PRC,Responsable:"+doc_coordinador, &pagos_mensuales); err == nil {

		for x, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[x].Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=Estado:TRUE,NumeroContrato:"+pagos_mensuales[x].NumeroContrato+",Vigencia:"+strconv.FormatFloat(pagos_mensuales[x].VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {

						for y, _ := range vinculaciones_docente {
							var dep []models.Dependencia

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(vinculaciones_docente[y].IdProyectoCurricular), &dep); err == nil {

								for z, _ := range dep {
									pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
									pago_personas_proyecto.NombrePersona = contratista.NomProveedor
									pago_personas_proyecto.Dependencia = &dep[z]
									pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)
								}

							} else { //If dependencia get

								fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
								return

							}

						}

					} else { // If vinculacion_docente_get

						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
						return
					}
				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}
		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}
	c.Data["json"] = pagos_personas_proyecto
	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenador
// @Description create GetSolicitudesOrdenador
// @Param docordenador query string true "Número del documento del ordenador"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /solicitudes_ordenador/:docordenador [get]
func (c *AprobacionPagoController) GetSolicitudesOrdenador() {

	doc_ordenador := c.GetString(":docordenador")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_personas_proyecto []models.PagoPersonaProyecto
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?limit=-1&query=EstadoPagoMensual.CodigoAbreviacion:AD,Responsable:"+doc_ordenador, &pagos_mensuales); err == nil {
		for x, pago_mensual := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=Estado:TRUE,NumeroContrato:"+pago_mensual.NumeroContrato+",Vigencia:"+strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {

						for _, vinculacion := range vinculaciones_docente {
							var dep models.Dependencia

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {

								pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
								pago_personas_proyecto.NombrePersona = contratista.NomProveedor
								pago_personas_proyecto.Dependencia = &dep

								pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)

							} else { //If dependencia get

								fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
								return

							}

						}

					} else { // If vinculacion_docente_get

						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
						return
					}
				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}
		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}
	c.Data["json"] = pagos_personas_proyecto
	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title ObtenerDependenciaOrdenador
// @Description create ObtenerDependenciaOrdenador
// @Param docordenador query string true "Número del documento del ordenador"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /dependencia_ordenador/:docordenador [get]
func (c *AprobacionPagoController) ObtenerDependenciaOrdenador() {

	doc_ordenador := c.GetString(":docordenador")

	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=TerceroId:"+doc_ordenador+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); err == nil {
		for _, jefe := range jefes_dependencia {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=DependenciaId:"+strconv.Itoa(jefe.DependenciaId), &ordenadores_gasto); err == nil {

				for _, ordenador := range ordenadores_gasto {

					c.Data["json"] = ordenador.DependenciaId

				}

			} else { // If ordenador_gasto get
				fmt.Println("Mirenme, me morí en If ordenador_gasto get, solucioname!!! ", err)
			}

		}

	} else { // If jefe_dependencia get
		fmt.Println("Mirenme, me morí en If jefe_dependencia get, solucioname!!! ", err)
	}
	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title AprobarMultiplesSolicitudes
// @Description create AprobarMultiplesSolicitudes
// @Success 201
// @Failure 403
// @router /aprobar_documentos [post]
func (c *AprobacionPagoController) AprobarMultiplesSolicitudes() {

	var v []models.PagoPersonaProyecto
	var response interface{}

	var pagos_mensuales []*models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var pago_mensual *models.PagoMensual
		for _, pm := range v {

			pago_mensual = pm.PagoMensual

			pagos_mensuales = append(pagos_mensuales, pago_mensual)
		}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/tr_aprobacion_masiva_documentos", "POST", &response, pagos_mensuales); err == nil {
			c.Data["json"] = "OK"
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title AprobarMultiplesPagos
// @Description create AprobarMultiplesPagos
// @Success 201
// @Failure 403
// @router /aprobar_pagos [post]
func (c *AprobacionPagoController) AprobarMultiplesPagos() {

	var v []models.PagoPersonaProyecto
	var response interface{}

	var pagos_mensuales []*models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var pago_mensual *models.PagoMensual
		for _, pm := range v {

			pago_mensual = pm.PagoMensual

			pagos_mensuales = append(pagos_mensuales, pago_mensual)
		}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/tr_aprobacion_masiva_pagos", "POST", &response, pagos_mensuales); err == nil {
			c.Data["json"] = "OK"
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title GetContratosContratista
// @Param numero_documento query string true "Número documento"
// @Success 201
// @Failure 403 :numero_cdp is empty
// @router /contratos_contratista/:numero_documento [get]
func (c *AprobacionPagoController) GetContratosContratista() {
	numero_documento := c.GetString(":numero_documento")
	var contratos_disponibilidad []models.ContratoDisponibilidad
	var contratos_disponibilidad_rp []models.ContratoDisponibilidadRp
	var novedades_postcontractuales []models.NovedadPostcontractual
	var informacion_proveedores []models.InformacionProveedor
	contratos_persona := GetContratosPersona(numero_documento)

	if contratos_persona.ContratosPersonas.ContratoPersona == nil {// Si no tiene contrato
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numero_documento,&informacion_proveedores); err == nil {
			
			for _,persona := range informacion_proveedores {
				
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=Contratista:"+ strconv.Itoa(persona.Id)+"&sortby=FechaInicio&order=desc&limit=1", &novedades_postcontractuales); err == nil {

					for _, novedad := range novedades_postcontractuales {
						var contrato models.InformacionContrato
						contrato = GetContrato(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))

						var informacion_contrato_contratista models.InformacionContratoContratista
		informacion_contrato_contratista = GetInformacionContratoContratista(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {


			for _, contrato_disponibilidad := range contratos_disponibilidad {


				var cdprp models.InformacionCdpRp
				cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

				for _, rp := range cdprp.CdpXRp.CdpRp {
					
					var contrato_disponibilidad_rp models.ContratoDisponibilidadRp

					contrato_disponibilidad_rp.NumeroContratoSuscrito = novedad.NumeroContrato
					contrato_disponibilidad_rp.Vigencia = strconv.Itoa(novedad.Vigencia)
					contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
					contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
					contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
					contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia

					contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
					contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion

					contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
				}

			}

		} else { // If contrato_disponibilidad get
			fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
		
		}
					}

				}else{ // If novedad_postcontractual get
					fmt.Println("Mirenme, me morí en If novedad_postcontractual get, solucioname!!! ", err.Error())
				}	
			}

		}else{// If informacion_proveedor get
			fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err.Error())
		}

	}else{// si tiene contrato
	for _, contrato_persona := range contratos_persona.ContratosPersonas.ContratoPersona {
		var contrato models.InformacionContrato
		contrato = GetContrato(contrato_persona.NumeroContrato, contrato_persona.Vigencia)

		//var novedad_postcontractual models.NovedadPostcontractual
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=NumeroContrato:"+contrato_persona.NumeroContrato+",Vigencia:"+contrato_persona.Vigencia+"&sortby=FechaInicio&order=desc&limit=1",&novedades_postcontractuales); err == nil {
		//var	prueba []models.NovedadPostcontractual

		//	json.NewDecoder(r.Body).Decode(prueba)
		var informacion_contrato_contratista models.InformacionContratoContratista
		informacion_contrato_contratista = GetInformacionContratoContratista(contrato_persona.NumeroContrato, contrato_persona.Vigencia)

		if novedades_postcontractuales != nil { // Si tiene novedades
	
		for _, novedad := range novedades_postcontractuales {
			if novedad.TipoNovedad == 219 {// si es una cesión
			}else{// si no es una cesión
				var cdprp models.InformacionCdpRp
				cdprp = GetRP(strconv.Itoa(novedad.NumeroCdp),strconv.Itoa(novedad.VigenciaCdp))

				for _, rp := range cdprp.CdpXRp.CdpRp {
				var contrato_disponibilidad_rp models.ContratoDisponibilidadRp

				contrato_disponibilidad_rp.NumeroContratoSuscrito = novedad.NumeroContrato
				contrato_disponibilidad_rp.Vigencia = strconv.Itoa(novedad.Vigencia)
				contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(novedad.NumeroCdp)
				contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(novedad.VigenciaCdp)
				contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
				contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia

				contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
				contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion

				contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
				}



			}

		}

		}else{// si no tiene novedades
	
	

		

		
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

			for _, contrato_disponibilidad := range contratos_disponibilidad {
				var cdprp models.InformacionCdpRp
				cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

				for _, rp := range cdprp.CdpXRp.CdpRp {
					var contrato_disponibilidad_rp models.ContratoDisponibilidadRp

					contrato_disponibilidad_rp.NumeroContratoSuscrito = contrato_persona.NumeroContrato
					contrato_disponibilidad_rp.Vigencia = contrato_persona.Vigencia
					contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
					contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
					contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
					contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia

					contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
					contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion

					contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
				}

			}

		} else { // If contrato_disponibilidad get
			fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
		
		}
	

	
	}
		} else { // If novedad_postcontractual get
			fmt.Println("Mirenme, me morí en If novedad_postcontractual get, solucioname!!! ", err.Error())
		}

	}

}

	c.Data["json"] = contratos_disponibilidad_rp

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title GetSolicitudesSupervisorContratistas
// @Description create GetSolicitudesSupervisorContratistas
// @Param docsupervisor query string true "Número del documento del supervisor"
// @Success 201
// @Failure 403 :docsupervisor is empty
// @router /solicitudes_supervisor_contratistas/:docsupervisor [get]
func (c *AprobacionPagoController) GetSolicitudesSupervisorContratistas() {

	doc_supervisor := c.GetString(":docsupervisor")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_contratista_cdp_rp []models.PagoContratistaCdpRp
	var contratos_disponibilidad []models.ContratoDisponibilidad
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?limit=-1&query=EstadoPagoMensual.CodigoAbreviacion:PRS,Responsable:"+doc_supervisor, &pagos_mensuales); err == nil {

		for v, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					var informacion_contrato_contratista models.InformacionContratoContratista
					informacion_contrato_contratista = GetInformacionContratoContratista(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))
					var contrato models.InformacionContrato
					contrato = GetContrato(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

						for _, contrato_disponibilidad := range contratos_disponibilidad {

							var cdprp models.InformacionCdpRp
							cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

							for _, rp := range cdprp.CdpXRp.CdpRp {
								var pago_contratista_cdp_rp models.PagoContratistaCdpRp

								pago_contratista_cdp_rp.PagoMensual = &pagos_mensuales[v]
								pago_contratista_cdp_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
								pago_contratista_cdp_rp.NombrePersona = contratista.NomProveedor
								pago_contratista_cdp_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
								pago_contratista_cdp_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
								pago_contratista_cdp_rp.NumeroRp = rp.RpNumeroRegistro
								pago_contratista_cdp_rp.VigenciaRp = rp.RpVigencia

								pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)

							}

						}

					} else { // If contrato_disponibilidad get
						fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
					}

				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}

		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}
	c.Data["json"] = pagos_contratista_cdp_rp

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistas
// @Description create GetSolicitudesOrdenadorContratistas
// @Param docordenador query string true "Número del documento del supervisor"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /solicitudes_ordenador_contratistas/:docordenador [get]
func (c *AprobacionPagoController) GetSolicitudesOrdenadorContratistas() {

	doc_ordenador := c.GetString(":docordenador")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_contratista_cdp_rp []models.PagoContratistaCdpRp
	var contratos_disponibilidad []models.ContratoDisponibilidad

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?limit=-1&query=EstadoPagoMensual.CodigoAbreviacion:AS,Responsable:"+doc_ordenador, &pagos_mensuales); err == nil {

		for v, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					var informacion_contrato_contratista models.InformacionContratoContratista
					informacion_contrato_contratista = GetInformacionContratoContratista(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))
					var contrato models.InformacionContrato
					contrato = GetContrato(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

						for _, contrato_disponibilidad := range contratos_disponibilidad {

							var cdprp models.InformacionCdpRp
							cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

							for _, rp := range cdprp.CdpXRp.CdpRp {
								var pago_contratista_cdp_rp models.PagoContratistaCdpRp

								pago_contratista_cdp_rp.PagoMensual = &pagos_mensuales[v]
								pago_contratista_cdp_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
								pago_contratista_cdp_rp.NombrePersona = contratista.NomProveedor
								pago_contratista_cdp_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
								pago_contratista_cdp_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
								pago_contratista_cdp_rp.NumeroRp = rp.RpNumeroRegistro
								pago_contratista_cdp_rp.VigenciaRp = rp.RpVigencia

								pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)

							}

						}

					} else { // If contrato_disponibilidad get
						fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
					}

				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}

		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}
	c.Data["json"] = pagos_contratista_cdp_rp

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title AprobarMultiplesSolicitudesContratistas
// @Description create AprobarMultiplesSolicitudesContratistas
// @Success 201
// @Failure 403
// @router /aprobar_soportes_contratistas [post]
func (c *AprobacionPagoController) AprobarMultiplesSolicitudesContratistas() {

	var v []models.PagoContratistaCdpRp
	var response interface{}

	var pagos_mensuales []*models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var pago_mensual *models.PagoMensual
		for _, pm := range v {

			pago_mensual = pm.PagoMensual

			pagos_mensuales = append(pagos_mensuales, pago_mensual)
		}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/tr_aprobacion_masiva_soportes_contratistas", "POST", &response, pagos_mensuales); err == nil {
			c.Data["json"] = "OK"
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title AprobarMultiplesPagosContratistas
// @Description create AprobarMultiplesPagosContratistas
// @Success 201
// @Failure 403
// @router /aprobar_pagos_contratistas [post]
func (c *AprobacionPagoController) AprobarMultiplesPagosContratistas() {

	var v []models.PagoContratistaCdpRp
	var response interface{}

	var pagos_mensuales []*models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var pago_mensual *models.PagoMensual
		for _, pm := range v {

			pago_mensual = pm.PagoMensual

			pagos_mensuales = append(pagos_mensuales, pago_mensual)
		}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/tr_aprobacion_masiva_pagos", "POST", &response, pagos_mensuales); err == nil {
			c.Data["json"] = "OK"
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

	c.ServeJSON()
}

func GetRP(numero_cdp string, vigencia_cdp string) (rp models.InformacionCdpRp) {

	var temp map[string]interface{}
	var temp_cdp_rp models.InformacionCdpRp

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudFinanciera")+"/"+"cdprp/"+numero_cdp+"/"+vigencia_cdp+"/01", &temp); err == nil {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			json.Unmarshal(json_cdp_rp, &temp_cdp_rp)
			rp = temp_cdp_rp
			return rp

		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}
	return rp
}

func GetContratosPersona(num_documento string) (contratos_persona models.InformacionContratosPersona) {

	var temp map[string]interface{}
	var contratos models.InformacionContratosPersona

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_persona/"+num_documento, &temp); err == nil {
		json_contratos, error_json := json.Marshal(temp)

		if error_json == nil {
			json.Unmarshal(json_contratos, &contratos)
			contratos_persona = contratos
			return contratos_persona

		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return contratos_persona

}

func GetContrato(num_contrato_suscrito string, vigencia string) (informacion_contrato models.InformacionContrato) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contrato/"+num_contrato_suscrito+"/"+vigencia, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)

		if error_json == nil {
			var contrato models.InformacionContrato
			json.Unmarshal(json_contrato, &contrato)
			informacion_contrato = contrato
			return informacion_contrato
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return informacion_contrato
}

func GetInformacionContratoContratista(num_contrato_suscrito string, vigencia string) (informacion_contrato_contratista models.InformacionContratoContratista) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"informacion_contrato_contratista/"+num_contrato_suscrito+"/"+vigencia, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)

		if error_json == nil {
			var contrato_contratista models.InformacionContratoContratista
			json.Unmarshal(json_contrato, &contrato_contratista)
			informacion_contrato_contratista = contrato_contratista
			return informacion_contrato_contratista
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return informacion_contrato_contratista
}
