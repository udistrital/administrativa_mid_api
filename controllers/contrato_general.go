package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
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
