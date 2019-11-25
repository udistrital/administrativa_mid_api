package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
	. "github.com/udistrital/golog"
)

// CambioEstadoContratoValidoController operations for CambioEstadoContratoValido
type CambioEstadoContratoValidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CambioEstadoContratoValidoController) URLMapping() {
	c.Mapping("ValidarCambioEstado", c.ValidarCambioEstado)
}

// ValidarCambiosEstado ...
// @Title ValidarCambiosEstado
// @Description create ValidarCambiosEstado
// @Success 201 {int} models.EstadoContrato
// @Failure 403 body is empty
// @router / [post]
func (this *CambioEstadoContratoValidoController) ValidarCambioEstado() {

	var estados []models.EstadoContrato //0: actual y 1:siguiente
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	reglasbase, err := CargarReglasBase("AdministrativaContratacion")
	if err != nil {
		beego.Error(err)
		this.Abort("400")
	}

	m := NewMachine().Consult(reglasbase)

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &estados); err == nil {

		if m.CanProve(`estado(` + strings.ToLower(estados[0].NombreEstado) + `,` + strings.ToLower(estados[1].NombreEstado) + `).`) {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = "true"
		} else {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = "false"
		}

	} else {
		alertErr.Type = "Error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		fmt.Println("error1: ", err)
	}

	this.Data["json"] = alertErr
	this.ServeJSON()

}
