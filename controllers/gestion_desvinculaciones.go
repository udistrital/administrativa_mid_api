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

// GestionDesvinculacionesController ...
// @Title ListarDocentesDesvinculados
// @Description create ListarDocentesDesvinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_desvinculados [get]
func (c *GestionDesvinculacionesController) ListarDocentesDesvinculados() {
	id_resolucion := c.GetString("id_resolucion")
	fmt.Println("resolucion a consultar")
	fmt.Println(id_resolucion)
	query := "?limit=-1&query=IdResolucion.Id:" + id_resolucion + ",Estado:false"
	var v []models.VinculacionDocente

	if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v); err2 == nil {
		for x, pos := range v {
			documento_identidad, _ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto = BuscarNombreProveedor(documento_identidad)
			v[x].NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			v[x].Dedicacion = BuscarNombreDedicacion(pos.IdDedicacion.Id)
			v[x].LugarExpedicionCedula = BuscarLugarExpedicion(pos.IdPersona)
		}

	} else {
		fmt.Println("Error de cosulta en vinculacion", err2)
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()
	//fmt.Println(v)

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
	var respuesta_mod_vin models.ModificacionVinculacion
	var respuesta string
	var vinculacion_nueva int;
	var temp_vinculacion [1]models.VinculacionDocente;

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


		//CAMBIAR ESTADO DE VINCULACIÓN DOCNETE
		for _, pos := range v.DocentesDesvincular {
			if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.Id),"PUT", &respuesta, pos); err2 == nil {
			fmt.Println("respuesta", respuesta)
		}else{
			fmt.Println("error en json",err2)
		}
		}


		temp_vinculacion[0] = models.VinculacionDocente {
				IdPersona: v.DocentesDesvincular[0].IdPersona,
				NumeroHorasSemanales:  v.DocentesDesvincular[0].NumeroHorasNuevas,
				NumeroSemanas:  v.DocentesDesvincular[0].NumeroSemanas,
				IdResolucion: &models.ResolucionVinculacionDocente {Id: v.IdNuevaResolucion},
				IdDedicacion: v.DocentesDesvincular[0].IdDedicacion,
				IdProyectoCurricular:  v.DocentesDesvincular[0].IdProyectoCurricular,
				Categoria:  v.DocentesDesvincular[0].Categoria ,
				Dedicacion:  v.DocentesDesvincular[0].Dedicacion,
				NivelAcademico: v.DocentesDesvincular[0].NivelAcademico ,
				Disponibilidad:  v.DisponibilidadNueva,

		};
		//CREAR NUEVA Vinculacion

		if err := sendJson("http://localhost:8082/v1/gestion_previnculacion/Precontratacion/insertar_previnculaciones","POST", &vinculacion_nueva, temp_vinculacion); err == nil {
			fmt.Println("vinculacion nueva",vinculacion_nueva)
		}else{
			fmt.Println("error en json de modificacion vinculacion",err)
		}

		//
		fmt.Println("Id para modificacion,res",v.IdModificacionResolucion)

		//ACTUALIZO TABLA MODIFICACION VINCULACION
		for _, pos := range v.DocentesDesvincular {
			temp:= models.ModificacionVinculacion {ModificacionResolucion: &models.ModificacionResolucion {Id: v.IdModificacionResolucion},VinculacionDocenteCancelada: &models.VinculacionDocente{Id:pos.Id},VinculacionDocenteRegistrada: &models.VinculacionDocente{Id:vinculacion_nueva},Horas: pos.NumeroHorasNuevas}
		if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/","POST", &respuesta_mod_vin, temp); err2 == nil {
			fmt.Println("respuesta modificacion vin", respuesta_mod_vin)
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
