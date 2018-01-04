package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/utilidades"
)

// Contrato_generalController operations for Contrato_genral
type Contrato_generalController struct {
	beego.Controller
}

// URLMapping ...
func (c *Contrato_generalController) URLMapping() {
	c.Mapping("GetContratoByContratoSuscritoId", c.GetContratoByContratoSuscritoId)
}

// GetContratoByContratoSuscritoId ...
// @Title GetContratoByContratoSuscritoId
// @Description get Contrato_genral by id
// @Param	id		path 	string	true		"numero del contrato a consultar"
// @Param	vigencia		path 	string	true		"numero del contrato a consultar"
// @Success 200 {object} models.Contrato_genral
// @Failure 403
// @router GetContratoByContratoSuscritoId/:id/:vigencia [get]
func (c *Contrato_generalController) GetContratoByContratoSuscritoId() {
	idStr := c.Ctx.Input.Param(":id")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	var infoContrato []map[string]interface{}
	if _, err := strconv.Atoi(vigenciaStr); err == nil {
		if err = getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general?query=Vigencia:"+vigenciaStr+",ContratoSuscrito.Vigencia:"+vigenciaStr+",ContratoSuscrito.Id:"+idStr, &infoContrato); err == nil {

			if infoContrato != nil {
				c.Data["json"] = infoContrato[0]
			} else {
				c.Data["json"] = map[string]interface{}{"Type": "error", "Body": "No data Found", "Code": "E_003"}
			}
		} else {
			//error en crud.
			fmt.Println(err.Error())
			c.Data["json"] = map[string]interface{}{"Type": "error", "Body": err, "Code": "E_002"}
		}
	} else {
		//si la vigencia no es un numero.
		c.Data["json"] = map[string]interface{}{"Type": "error", "Body": err, "Code": "E_001"}
	}
	c.ServeJSON()
}

//funcion para armar info de los contratos de contratistas.
func FormatoInfoContratoContratoSuscrito(contratoIntfc interface{}, params ...interface{}) (res interface{}) {
	if infoContrato, e := contratoIntfc.(map[string]interface{}); e {
		idContratista := infoContrato["Contratista"].(float64)
		var infoContratista map[string]interface{}
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/"+strconv.Itoa(int(idContratista)), &infoContratista); err == nil {
			infoContrato["Contratista"] = infoContratista
			return infoContrato
		} else {
			return infoContrato
		}
	} else {
		return
	}

}

// ListaContratoContratoSuscrito ...
// @Title ListaContratoContratoSuscrito
// @Description get Disponibilidad by vigencia
// @Param	vigencia	query	string	false	"vigencia de la lista"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	rangoinicio	query	string	false	"rango inicial del periodo a consultar"
// @Param	rangofin	query	string	false	"rango final del periodo a consultar"
// @Param	query	query	string	false	"query de filtrado para la lista de los cdp"
// @Success 200 {object} models.Disponibilidad
// @Failure 403
// @router ListaContratoContratoSuscrito/:vigencia [get]
func (c *Contrato_generalController) ListaContratoContratoSuscrito() {
	var infoContrato []interface{}
	var respuesta []map[string]interface{}
	var limit int64 = 10
	var offset int64
	var startrange string
	var endrange string
	var query string
	var querybase string
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if r := c.GetString("rangoinicio"); r != "" {
		startrange = r

	}

	if r := c.GetString("rangofin"); r != "" {
		endrange = r

	}
	if r := c.GetString("query"); r != "" {
		querybase = r

	}
	if startrange != "" && endrange != "" {
		query = querybase + ",FechaRegistro__gte:" + startrange + ",FechaRegistro__lte:" + endrange

	} else if querybase != "" {
		query = "," + querybase
	}
	if querybase != "" {
		query = "," + querybase
	}
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	_, err1 := strconv.Atoi(vigenciaStr)
	//UnidadEjecutora, err2 := c.GetInt("UnidadEjecutora")
	if err1 == nil { //&& err2 == nil {
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general?limit="+strconv.FormatInt(limit, 10)+"&offset="+strconv.FormatInt(offset, 10)+"&query=VigenciaContrato:"+vigenciaStr+",ContratoSuscrito.Vigencia:"+vigenciaStr+query, &infoContrato); err == nil {
			if infoContrato != nil {
				done := make(chan interface{})
				defer close(done)
				resch := utilidades.GenChanInterface(infoContrato...)
				chdisponibilidades := utilidades.Digest(done, FormatoInfoContratoContratoSuscrito, resch, nil)
				for contrato := range chdisponibilidades {
					respuesta = append(respuesta, contrato.(map[string]interface{}))
				}
				c.Data["json"] = respuesta
			} else {
				c.Data["json"] = map[string]interface{}{"Code": "E_0458", "Body": nil, "Type": "error"}
			}
		} else {
			c.Data["json"] = map[string]interface{}{"Code": "E_0458", "Body": err.Error(), "Type": "error"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"Code": "E_0458", "Body": "Not enough parameter", "Type": "error"}
	}

	c.ServeJSON()
}
