package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
	//."github.com/udistrital/golog"
)

type ValidarContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidarContratoController) URLMapping() {
	c.Mapping("ValidarContrato", c.ValidarContrato)
}

// ValidarContrato ...
// @Title ValidarContrato
// @Description create ValidarContrato
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router /:dedicacion/:numHoras [get]
func (c *ValidarContratoController) ValidarContrato() {
	dedicacion := c.Ctx.Input.Param(":dedicacion")
	numHorasStr := c.Ctx.Input.Param(":numHoras")

	reglasbase := CargarReglasBase("CDVE")
	m := NewMachine().Consult(reglasbase)
	
	var a string

	contratos := m.ProveAll(`cumple_tiempo(` + strings.ToLower(dedicacion) + `,` + numHorasStr + `,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	validez, _ := strconv.Atoi(a)

	c.Data["json"] = validez

	c.ServeJSON()
}
