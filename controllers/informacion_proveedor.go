package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

// InformacionProveedorController operations for InformacionProveedor
type InformacionProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionProveedorController) URLMapping() {
	c.Mapping("contratoPersona", c.ContratoPersona)
}

// ContratoPersona ...
// @Title contratoProveedor
// @Description create InformacionProveedor
// @Param	body		body 	models.ContratoGeneral	true		"body for ContratoGeneral content"
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router /contratoPersona [post]

func (c *InformacionProveedorController) ContratoPersona() {
	var v []models.ContratoGeneral
	var datos string
	if err2 := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err2 == nil {
		query := "?query=" + datos
		fmt.Println(query)
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudArgo")+":"+beego.AppConfig.String("PortcrudArgo")+"/"+beego.AppConfig.String("NscrudArgo")+"/contrato_general/"+query, &v); err == nil {
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err2.Error()
		fmt.Println(err2)
	}
	c.ServeJSON()
}
