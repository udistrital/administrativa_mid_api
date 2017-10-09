package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
	. "github.com/mndrix/golog"
	//."github.com/udistrital/administrativa_mid_api/mndrix/golog"
)

// CambioEstadoContratoValidoController operations for CambioEstadoContratoValido
type CambioEstadoContratoValidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CambioEstadoContratoValidoController) URLMapping() {
	c.Mapping("ValidarCambioEstado", c.ValidarCambioEstado)
}

func (this *CambioEstadoContratoValidoController) ValidarCambioEstado() {

	var estados []models.EstadoContrato //0: actual y 1:siguiente

	reglasbase := CargarReglasBase("AdministrativaContratacion")

	m := NewMachine().Consult(reglasbase)

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &estados); err == nil {

		if m.CanProve(`estado(` + strings.ToLower(estados[0].NombreEstado) + `,` + strings.ToLower(estados[1].NombreEstado) + `).`) {
			this.Data["json"] = "true"
		} else {
			this.Data["json"] = "false"
		}

	} else {
		this.Data["json"] = err.Error()
		fmt.Println("error1: ", err)
	}

	this.ServeJSON()

}
