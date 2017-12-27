package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// Contrato_genralController operations for Contrato_genral
type Contrato_genralController struct {
	beego.Controller
}

// URLMapping ...
func (c *Contrato_genralController) URLMapping() {
	c.Mapping("GetContratoById", c.GetContratoById)
}

// GetContratoById ...
// @Title GetContratoById
// @Description get Contrato_genral by id
// @Param	id		path 	string	true		"numero del contrato a consultar"
// @Param	vigencia		path 	string	true		"numero del contrato a consultar"
// @Success 200 {object} models.Contrato_genral
// @Failure 403
// @router GetContratoById/:id/:vigencia [get]
func (c *Contrato_genralController) GetContratoById() {
	idStr := c.Ctx.Input.Param(":id")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	var infoContrato []map[string]interface{}
	if _, err := strconv.Atoi(vigenciaStr); err == nil {
		if err = getJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/contrato_general?query=ContratoSuscrito.Vigencia:"+vigenciaStr+",Id:"+idStr, &infoContrato); err == nil {

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
