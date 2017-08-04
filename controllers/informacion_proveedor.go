package controllers

import (
	"github.com/udistrital/administrativa_mid_api/models"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
)

// InformacionProveedorController operations for InformacionProveedor
type InformacionProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionProveedorController) URLMapping() {
	c.Mapping("contratoProveedor", c.Contrato_proveedor)
}

// ContratoPersona ...
// @Title contratoProveedor
// @Description create InformacionProveedor
// @Param	body		body 	models.ContratoGeneral	true		"body for ContratoGeneral content"
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router /contratoPersona [post]

func (c *InformacionProveedorController) Contrato_proveedor() {
	var persona_natural []models.InformacionPersonaNatural
	var informacion_proveedor []models.InformacionProveedor
	var datos []models.ContratoGeneral
	var contrato_proveedor []models.ContratoProveedor

	if err2 := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err2 == nil {
/*		query := "?limit=-1&query=" + datos
		fmt.Println(query)*/
		queryPersonaNatural := "?query=Id:"+strconv.Itoa(datos[7].Contratista)

		queryInformacionProveedor := "?query=NumDocumento:"+strconv.Itoa(datos[7].Contratista)
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudArgo")+"/informacion_persona_natural/"+queryPersonaNatural, &persona_natural); err == nil {
		if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudArgo")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
			fmt.Println("persona NATURAL")
			fmt.Println(persona_natural[0])
			fmt.Println("informacion PROVEEDOR")
			fmt.Println(informacion_proveedor[0])
			contrato_proveedor[0].InformacionProveedor = informacion_proveedor[0]
			contrato_proveedor[0].InformacionPersonaNatural = persona_natural[0]
			c.Data["json"] = contrato_proveedor[0]
		}else {
			c.Data["json"] = err2.Error()
		}

		}else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err2.Error()
		fmt.Println(err2)
	}
	fmt.Println(c)
	c.ServeJSON()
}
