package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
)

type ValidarContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidarContratoController) URLMapping() {
	c.Mapping("ValidarContrato", c.ValidarContrato)
}

func (c *ValidarContratoController) ValidarContrato() {
	dedicacion := c.Ctx.Input.Param(":dedicacion")
	numHorasStr := c.Ctx.Input.Param(":numHoras")

	reglasbase := CargarReglasBase("CVDE")

	m := NewMachine().Consult(reglasbase)

	var a string
	contratos := m.ProveAll(`cumple_tiempo(` + strings.ToLower(dedicacion) + `,` + numHorasStr + `,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	fmt.Println(a)
	validez, _ := strconv.Atoi(a)

	c.Data["json"] = validez

	c.ServeJSON()
}
