package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
	"strconv"
)

// PreliquidacionController operations for Preliquidacion
type GestionDesvinculacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionDesvinculacionesController) URLMapping() {

	c.Mapping("ActualizarVinculaciones", c.ActualizarVinculaciones)
	c.Mapping("AdicionarHoras", c.AdicionarHoras)
}

// ActualizarVinculaciones ...
// @Title ActualizarVinculaciones
// @Description create ActualizarVinculaciones
// @Success 201 {string}
// @Failure 403 body is empty
// @router actualizar_vinculaciones [post]
func (c *GestionDesvinculacionesController) ActualizarVinculaciones() {

	var v models.Objeto_Desvinculacion
	var respuesta string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("para poner en false",v)

		for _, pos := range v.DocentesDesvincular {
		if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id),"PUT", &respuesta, pos); err2 == nil {
			fmt.Println("respuesta", respuesta)
		}else{
			fmt.Println("error en json",err2)
		}
		}

		fmt.Println("Id para modificacion,res",v.IdModificacionResolucion)

		for _, pos := range v.DocentesDesvincular {
			temp:= models.ModificacionVinculacion {ModificacionResolucion: &models.ModificacionResolucion {Id: v.IdModificacionResolucion},VinculacionDocenteCancelada: &models.VinculacionDocente{Id:pos.Id}}
		if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/","POST", &respuesta, temp); err2 == nil {
			fmt.Println("respuesta", respuesta)
		}else{
			fmt.Println("error en json de modificacion vinculacion",err2)
		}
		}

		c.Data["json"] = respuesta
	} else {
		fmt.Println("ERROR")
		fmt.Println(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	c.ServeJSON()

}


// AdicionarHoras ...
// @Title AdicionarHoras
// @Description create AdicionarHoras
// @Success 201 {string}
// @Failure 403 body is empty
// @router adicionar_horas [post]
func (c *GestionDesvinculacionesController) AdicionarHoras() {

	var v models.Objeto_Desvinculacion
	var respuesta string
	var vinculacion_nueva models.VinculacionDocente
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("docentes a adicionar",v.DocentesDesvincular[0].Vigencia)

		//CAMBIAR ESTADO DE VINCULACIÓN DOCNETE
		for _, pos := range v.DocentesDesvincular {
			if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id),"PUT", &respuesta, pos); err2 == nil {
			fmt.Println("respuesta", respuesta)
		}else{
			fmt.Println("error en json",err2)
		}
		}

		//CREAR NUEVA Vinculacion
		for _, pos := range v.DocentesDesvincular {
		if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/","POST", &vinculacion_nueva, pos); err2 == nil {
			fmt.Println("vinculacion nueva",vinculacion_nueva)
		}else{
			fmt.Println("error en json de modificacion vinculacion",err2)
		}
		}
		//
		fmt.Println("Id para modificacion,res",v.IdModificacionResolucion)
		/*
		//ACTUALIZO TABLA MODIFICACION VINCULACION
		for _, pos := range v.DocentesDesvincular {
			temp:= models.ModificacionVinculacion {ModificacionResolucion: &models.ModificacionResolucion {Id: v.IdModificacionResolucion},VinculacionDocenteCancelada: &models.VinculacionDocente{Id:pos.Id},VinculacionDocenteRegistrada: &models.VinculacionDocente{Id:vinculacion_registrada},Horas: v.NumeroHorasSemanales}
		if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/","POST", &respuesta, temp); err2 == nil {
			fmt.Println("respuesta", respuesta)
		}else{
			fmt.Println("error en json de modificacion vinculacion",err2)
		}
		}
*/
		c.Data["json"] = respuesta
	} else {
		fmt.Println("ERROR")
		fmt.Println(err)
		c.Data["json"] = "Error al leer json para desvincular"
	}

	c.ServeJSON()

}
