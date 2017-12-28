package controllers

import (
	"fmt"
	"time"
	//"strconv"
	//"strings"
	"encoding/json"
	"github.com/astaxie/beego"
	//. "github.com/mndrix/golog"
	"github.com/udistrital/administrativa_mid_api/models"
	//. "github.com/udistrital/golog"
)

//GestionResolucionesController operations for Preliquidacion
type GestionResolucionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionResolucionesController) URLMapping() {
	c.Mapping("InsertarResolucionCompleta", c.InsertarResolucionCompleta)

}

// InsertarResolucionCompleta ...
// @Title InsertarResolucionCompleta
// @Description create InsertarResolucionCompleta
// @Success 201 {int} models.Resolucion
// @Failure 403 body is empty
// @router /insertar_resolucion_completa [post]
func (c *GestionResolucionesController) InsertarResolucionCompleta() {
	var v models.ObjetoResolucion
	var id_resolucion_creada int
	var control bool
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//****MANEJO DE TRANSACCIONES!***!//
		//Primero se inserta la resolución, si eso se realiza correctamente
		control,id_resolucion_creada = InsertarResolucion(v.Resolucion);
		if(control){
				//Si se inserta bien en resolución, se puede insertar en resolucion_vinculacion_docente y en resolucion_estado
					control=InsertarResolucionVinDocente(id_resolucion_creada,v.ResolucionVinculacionDocente);
					control=InsertarResolucionEstado(id_resolucion_creada);
					//Si todo sigue bien, se inserta en componente_resolucion
					if(control){
							InsertarTexto(id_resolucion_creada);
					}else{
						fmt.Println("enviar error al insertar en resolucion_vinculacion_docente")
					}
		}else{
			fmt.Println("envia error al insertar en resolución")
		}

	}else{
		fmt.Println("error al leer objeto resolucion",err)
	}

	if(control){
		fmt.Println("okey")
		c.Data["json"] = "OK"
	}else{
		fmt.Println("not okey")
		c.Data["json"] = "Error"
	}
	c.ServeJSON()
}

func InsertarResolucion(resolucion *models.Resolucion)(contr bool,id_cre int){
	var temp = resolucion
	var respuesta models.Resolucion
	var id_creada int;
	var cont bool;

	temp.Vigencia, _, _ = time.Now().Date()
	temp.FechaRegistro = time.Now()
	temp.Estado = true


	if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion", "POST", &respuesta, &temp); err == nil {
		id_creada = respuesta.Id;
		cont = true;
	} else {
		cont = false;
		id_creada = 0;
	}

	return cont,id_creada
}

func InsertarResolucionEstado(id_res int)(contr bool){

	var respuesta models.ResolucionEstado
	var cont bool;
	temp:= models.ResolucionEstado {
		FechaRegistro: time.Now(),
		Estado: &models.EstadoResolucion{Id:1},
		Resolucion: &models.Resolucion{Id:id_res},
	}

	if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &respuesta, &temp); err == nil {
		cont = true;
	} else {
		cont = false;
	}

	return cont
}

func InsertarResolucionVinDocente(id_res int, resvindoc *models.ResolucionVinculacionDocente)(contr bool){
	var temp = resvindoc
	var respuesta models.ResolucionVinculacionDocente

	var cont bool
	temp.Id = id_res

	if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente", "POST", &respuesta, &temp); err == nil {

		cont = true;
	} else {

		cont = false;
	}

	return cont
}

func InsertarTexto(id_res int){
	var texto_resolucion models.ResolucionCompleta

	if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/ResolucionTemplate", &texto_resolucion); err2 == nil {
			InsertarArticulos(id_res, texto_resolucion.Articulos)

	}else{
		fmt.Println("Error de consulta en vinculacion",err2)
	}
}

func InsertarArticulos(id_resolucion int, articulos []models.Articulo){
fmt.Println("Articulos y parágrafos")
var respuesta models.ComponenteResolucion

	for x, pos := range  articulos {
		temp:= models.ComponenteResolucion{
				Numero: x+1,
				ResolucionId: &models.Resolucion{Id: id_resolucion},
				Texto: pos.Texto,
				TipoComponente: "Articulo"}
		if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta, &temp); err == nil {
			for y,pos2 := range pos.Paragrafos {
				temp2:= models.ComponenteResolucion{
						Numero: y+1,
						ResolucionId: &models.Resolucion{Id: id_resolucion},
						Texto: pos2.Texto,
						TipoComponente: "Paragrafo",
						ComponentePadre: &models.ComponenteResolucion{Id: respuesta.Id},
					}

					if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta, &temp2); err == nil {

					}else{
							fmt.Println("error al insertar parágrafos",err2)
					}
			}


		}else {
			fmt.Println("error al insertar articulos",err)
		}
	}

}
