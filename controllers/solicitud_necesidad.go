package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

// SolicitudNecesidadController operations for SolicitudNecesidad
type SolicitudNecesidadController struct {
	beego.Controller
}

// FuentesApropiacionNecesidad ...
// @Title FuentesApropiacionNecesidad
// @Description create FuentesApropiacionNecesidad
// @Param id_necesidad path string true "necesidad a consultar"
// @Success 201 {int} models.FuenteApropiacionNecesidad
// @Failure 403 body is empty
// @router /fuente_apropiacion_necesidad/:id_necesidad [get]
func (c *SolicitudNecesidadController) FuentesApropiacionNecesidad() {
	idNecesidad := c.Ctx.Input.Param(":id_necesidad")
	var fuentes []models.FuenteFinanciacionRubroNecesidad
	var fuentesApropiacionNecesidad []models.FuenteApropiacionNecesidad
	var fuenteApropiacionNecesidad models.FuenteApropiacionNecesidad
	var apropiacion []models.ApropiacionRubro
	var montoFuentes float64
	var productos []models.ProductoRubroNecesidad

	//Devuelve las fuentes de financiamiento asociadas a la necesidad
	query := "?limit=-1&query=Necesidad:" + idNecesidad
	err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/fuente_financiacion_rubro_necesidad"+query, &fuentes)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	//Devuelve los IDs de las apropiaciones asociadas a la necesidad
	apropiacionesIds := agruparPorApropiacion(fuentes)

	for _, pos := range apropiacionesIds {
		apropiacionID := strconv.Itoa(pos)
		//Consulta la información de cada apropiación en financiera
		query = "?&query=Id:" + apropiacionID
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/apropiacion"+query, &apropiacion)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

		//Consulta las fuentes asociadas a la necesidad y a cada apropiación
		query = "?limit=-1&query=Necesidad:" + idNecesidad + ",Apropiacion:" + apropiacionID
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/fuente_financiacion_rubro_necesidad"+query, &fuentes)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

		for x, fuente := range fuentes {
			montoFuentes += fuente.MontoParcial
			//Consulta la información de cada fuente en financiera
			query = "?&query=Id:" + strconv.Itoa(fuente.FuenteFinanciamiento)
			err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/fuente_financiamiento"+query, &fuentes[x].InfoFuente)
			if err != nil {
				beego.Error(err)
				c.Abort("400")
			}
		}

		//Consulta los productos asociados a la necesidad y a cada apropiación
		query = "?limit=-1&query=Necesidad:" + idNecesidad + ",Apropiacion:" + apropiacionID
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/producto_rubro_necesidad"+query, &productos)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

		for x, producto := range productos {
			//Consulta la información de cada producto en financiera
			query = "?&query=Id:" + strconv.Itoa(producto.ProductoRubro)
			err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/producto_rubro"+query, &productos[x].ProductoRubroInfo)
			if err != nil {
				beego.Error(err)
				c.Abort("400")
			}
		}

		fuenteApropiacionNecesidad.Apropiacion = apropiacion[0]
		fuenteApropiacionNecesidad.Fuentes = fuentes
		fuenteApropiacionNecesidad.Productos = productos
		fuenteApropiacionNecesidad.Monto = montoFuentes

		jsonAp, errr := json.Marshal(fuenteApropiacionNecesidad)
		if errr != nil {
			beego.Error(errr)
			c.Abort("400")
		}

		var tempApr models.FuenteApropiacionNecesidad
		errr = json.Unmarshal(jsonAp, &tempApr)
		if errr != nil {
			beego.Error(errr)
			c.Abort("400")
		}

		fuentesApropiacionNecesidad = append(fuentesApropiacionNecesidad, tempApr)
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = fuentesApropiacionNecesidad
	c.ServeJSON()

}

func existeEnArreglo(id int, arreglo []int) bool {
	for _, pos := range arreglo {
		if pos == id {
			return true
		}
	}
	return false
}

func agruparPorApropiacion(fuentes []models.FuenteFinanciacionRubroNecesidad) (apropiacionesIds []int) {
	for _, pos := range fuentes {
		if !existeEnArreglo(pos.Apropiacion, apropiacionesIds) {
			apropiacionesIds = append(apropiacionesIds, pos.Apropiacion)
		}
	}
	return apropiacionesIds
}
