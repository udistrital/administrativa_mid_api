package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

//GestionResolucionesController operations for Preliquidacion
type GestionResolucionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionResolucionesController) URLMapping() {
	c.Mapping("InsertarResolucionCompleta", c.InsertarResolucionCompleta)

}

// GestionResolucionesController ...
// @Title getResolucionesInscritas
// @Description create  getResolucionesInscritas
// @Param vigencia query string false "año a consultar"
// @Success 201 {object} []models.ResolucionVinculacion
// @Failure 403 body is empty
// @router /get_resoluciones_inscritas [get]
func (c *GestionResolucionesController) GetResolucionesInscritas() {
	var resolucion_vinculacion []models.ResolucionVinculacion
	fmt.Println(beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAdmin") + "/" + beego.AppConfig.String("NscrudAdmin") + "/resolucion_vinculacion")
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion", &resolucion_vinculacion); err2 == nil {
		for x, pos := range resolucion_vinculacion {
			resolucion_vinculacion[x].FacultadNombre = BuscarNombreFacultad(pos.Facultad)

		}

		c.Data["json"] = resolucion_vinculacion

	} else {
		c.Data["json"] = "error"
		fmt.Println("Error de consulta en resolucion_vinculacion", err2)
	}
	c.ServeJSON()
}

// GestionResolucionesController ...
// @Title getResolucionesAprobadas
// @Description create  getResolucionesAprobadas
// @Param vigencia query string false "año a consultar"
// @Success 201 {object} []models.ResolucionVinculacion
// @Failure 403 body is empty
// @router /get_resoluciones_aprobadas [get]
func (c *GestionResolucionesController) GetResolucionesAprobadas() {
	var resolucion_vinculacion_aprobada []models.ResolucionVinculacion

	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion/Aprobada", &resolucion_vinculacion_aprobada); err2 == nil {
		for x, pos := range resolucion_vinculacion_aprobada {
			resolucion_vinculacion_aprobada[x].FacultadNombre = BuscarNombreFacultad(pos.Facultad)

		}

		c.Data["json"] = resolucion_vinculacion_aprobada

	} else {
		c.Data["json"] = "error"
		fmt.Println("Error de consulta en resolucion_vinculacion_aprobada", err2)
	}
	c.ServeJSON()
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
	var texto_resolucion models.ResolucionCompleta

	var control bool
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//****MANEJO DE TRANSACCIONES!***!//

		//Se trae cuerpo de resolución según tipo
		if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/ResolucionTemplate/"+v.ResolucionVinculacionDocente.Dedicacion+"/"+v.ResolucionVinculacionDocente.NivelAcademico, &texto_resolucion); err2 == nil {
			v.Resolucion.ConsideracionResolucion = texto_resolucion.Consideracion
		} else {
			fmt.Println("Error de consulta en texto de resolucion", err2)
		}

		//Primero se inserta la resolución, si eso se realiza correctamente
		control, id_resolucion_creada = InsertarResolucion(v)
		if control {
			//Si se inserta bien en resolución, se puede insertar en resolucion_vinculacion_docente y en resolucion_estado
			control = InsertarResolucionVinDocente(id_resolucion_creada, v.ResolucionVinculacionDocente)
			control = InsertarResolucionEstado(id_resolucion_creada)
			//Si todo sigue bien, se inserta en componente_resolucion
			if control {
				InsertarTexto(id_resolucion_creada, v.ResolucionVinculacionDocente.Dedicacion, v.ResolucionVinculacionDocente.NivelAcademico)
			} else {
				fmt.Println("enviar error al insertar en resolucion_vinculacion_docente")
			}
		} else {
			fmt.Println("envia error al insertar en resolución")
		}

	} else {
		fmt.Println("error al leer objeto resolucion", err)
	}

	if control {
		fmt.Println("okey")
		c.Data["json"] = id_resolucion_creada
	} else {
		fmt.Println("not okey")
		c.Data["json"] = "Error"
	}
	c.ServeJSON()
}

func InsertarResolucion(resolucion models.ObjetoResolucion) (contr bool, id_cre int) {
	var temp = resolucion.Resolucion
	var respuesta models.Resolucion
	var id_creada int
	var cont bool
	var respuesta_modificacion_res models.ModificacionResolucion

	temp.Vigencia, _, _ = time.Now().Date()
	temp.FechaRegistro = time.Now()
	temp.Estado = true
	temp.Titulo = "Por la cual se vinculan docentes para el Primer Periodo Académico de 2018 en la modalidad de Docentes de HORA CÁTEDRA (Vinculación Especial) para la " + resolucion.NomDependencia + " en " + resolucion.ResolucionVinculacionDocente.NivelAcademico + ".”"
	temp.PreambuloResolucion = "El decano de la " + resolucion.NomDependencia + " de la Universidad Distrital Francisco José de Caldas en uso de sus facultades legales y estatuarias y"
	if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion", "POST", &respuesta, &temp); err == nil {
		id_creada = respuesta.Id
		cont = true
	} else {
		cont = false
		id_creada = 0
	}

	if temp.IdTipoResolucion.Id != 1 {
		fmt.Println(resolucion.ResolucionVieja)
		objeto_modificacion_res := models.ModificacionResolucion{
			ResolucionAnterior: resolucion.ResolucionVieja,
			ResolucionNueva:    id_creada,
		}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion", "POST", &respuesta_modificacion_res, &objeto_modificacion_res); err == nil {
			cont = true
		} else {
			fmt.Println("error al insertar en modificacion resolucion", err)
			cont = false

		}

	}

	return cont, id_creada
}

func InsertarResolucionEstado(id_res int) (contr bool) {

	var respuesta models.ResolucionEstado
	var cont bool
	temp := models.ResolucionEstado{
		FechaRegistro: time.Now(),
		Estado:        &models.EstadoResolucion{Id: 1},
		Resolucion:    &models.Resolucion{Id: id_res},
	}

	if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_estado", "POST", &respuesta, &temp); err == nil {
		cont = true
	} else {
		cont = false
	}

	return cont
}

func InsertarResolucionVinDocente(id_res int, resvindoc *models.ResolucionVinculacionDocente) (contr bool) {
	var temp = resvindoc
	var respuesta models.ResolucionVinculacionDocente

	var cont bool
	temp.Id = id_res

	if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion_vinculacion_docente", "POST", &respuesta, &temp); err == nil {

		cont = true
	} else {

		cont = false
	}

	return cont
}

func InsertarTexto(id_res int, dedicacion, nivel_academico string) {
	var texto_resolucion models.ResolucionCompleta

	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contenido_resolucion/ResolucionTemplate/"+dedicacion+"/"+nivel_academico, &texto_resolucion); err2 == nil {
		InsertarArticulos(id_res, texto_resolucion.Articulos)

	} else {
		fmt.Println("Error de consulta en vinculacion", err2)
	}
}

func InsertarArticulos(id_resolucion int, articulos []models.Articulo) {
	fmt.Println("Articulos y parágrafos")
	var respuesta models.ComponenteResolucion

	for x, pos := range articulos {
		temp := models.ComponenteResolucion{
			Numero:         x + 1,
			ResolucionId:   &models.Resolucion{Id: id_resolucion},
			Texto:          pos.Texto,
			TipoComponente: "Articulo"}
		if err := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta, &temp); err == nil {
			for y, pos2 := range pos.Paragrafos {
				temp2 := models.ComponenteResolucion{
					Numero:          y + 1,
					ResolucionId:    &models.Resolucion{Id: id_resolucion},
					Texto:           pos2.Texto,
					TipoComponente:  "Paragrafo",
					ComponentePadre: &models.ComponenteResolucion{Id: respuesta.Id},
				}

				if err2 := sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/componente_resolucion", "POST", &respuesta, &temp2); err == nil {

				} else {
					fmt.Println("error al insertar parágrafos", err2)
				}
			}

		} else {
			fmt.Println("error al insertar articulos", err)
		}
	}

}

func BuscarNombreFacultad(id_facultad int) (nombre_facultad string) {

	var facultad []models.Facultad
	var nom string
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia?query=Id:"+strconv.Itoa(id_facultad), &facultad); err2 == nil {
		nom = facultad[0].Nombre
	} else {
		nom = "N/A"
	}
	return nom
}
