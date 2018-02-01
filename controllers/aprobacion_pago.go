package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_homologacion_dependencias/proyecto_curricular_oikos/"+id_oikos, &temp); err == nil && temp != nil {
		json_proyecto_curricular, error_json := json.Marshal(temp)

		if error_json == nil {
			var temp_homologacion models.ObjetoProyectoCurricular
			json.Unmarshal(json_proyecto_curricular, &temp_homologacion)
			id_proyecto_snies := temp_homologacion.Homologacion.IDSnies

			if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/academicaProxy/carrera_snies/"+id_proyecto_snies, &temp_snies); err == nil && temp_snies != nil {
				json_info_coordinador, error_json := json.Marshal(temp_snies)

				if error_json == nil {
					var temp_info_coordinador models.InformacionCoordinador
					json.Unmarshal(json_info_coordinador, &temp_info_coordinador)
					fmt.Println(temp_info_coordinador)
					info_coordinador = temp_info_coordinador
				} else {
					fmt.Println(error_json.Error())
					// c.Data["json"] = error_json.Error()
				}
			}

		} else {
			fmt.Println(error_json.Error())
			// c.Data["json"] = error_json.Error()
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
	//var contrato []models.ContratoGeneral
	var contratoEstado []models.ContratoEstado
	var res models.Resolucion
	var dep models.Dependencia
	//If informacion_proveedor get
	if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=num_documento:"+numDocumento, &proveedor); err == nil {
		//If vinculacion_docente get
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=id_persona:"+numDocumento, &vinculaciones); err == nil {
			//for vinculaciones
			for _, vinculacion := range vinculaciones {
				//If dependencia get
				if err := getJson("http://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {
					//If resolucion get
					if err := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(vinculacion.IdResolucion.Id), &res); err == nil {
						//If nulo
						if vinculacion.NumeroContrato.Valid == true {
							if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado/?query=NumeroContrato:"+vinculacion.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion.Vigencia.Int64, 10)+"&sortby=FechaRegistro&order=desc&limit=1", &contratoEstado); err == nil {
								//If Estado = 4
								if contratoEstado[0].Estado.Id == 4 {
									cd.NumeroVinculacion = vinculacion.NumeroContrato.String
									cd.Vigencia = vinculacion.Vigencia.Int64
									cd.Resolucion = res.NumeroResolucion
									cd.Dependencia = dep.Nombre
									cd.IdDependencia = dep.Id
									cd.NombreDocente = proveedor[0].NomProveedor
									cd.Dedicacion = vinculacion.IdDedicacion.NombreDedicacion
									contratosDocentes = append(contratosDocentes, cd)
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
// @Success 201 {int} models.ContratoElaborado
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @router /informacion_ordenador/:numero_contrato/:vigencia [get]
func (c *AprobacionPagoController) ObtenerInfoOrdenador() {
	numero_contrato:= c.GetString(":numero_contrato")
	vigencia:= c.GetString(":vigencia")

   var temp map[string]interface{}
   //var temp_snies map[string]interface{}
   var contrato_elaborado models.ContratoElaborado

   if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/contratoSuscritoProxyService/contrato_elaborado/"+numero_contrato+"/"+vigencia, &temp); err == nil && temp != nil {
	   json_contrato_elaborado, error_json := json.Marshal(temp)

	   if error_json == nil {
		   var temp_contrato_elaborado models.ContratoElaborado
		   json.Unmarshal(json_contrato_elaborado, &temp_contrato_elaborado)
   
		   contrato_elaborado = temp_contrato_elaborado
			fmt.Println(temp)

	   } else {
		   fmt.Println(error_json.Error())
		   // c.Data["json"] = error_json.Error()
	   }
   } else {
	   fmt.Println(err)

   }

   c.Data["json"] = contrato_elaborado
   c.ServeJSON()
}