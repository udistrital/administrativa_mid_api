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