package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

type CancelacionValidaController struct {
	beego.Controller
}

// URLMapping ...
func (c *CancelacionValidaController) URLMapping() {
	c.Mapping("ValidarCancelacion", c.ValidarCancelacion)
}

// ValidarCancelacion ...
// @Title ValidarCancelacion
// @Description create ValidarCancelacion
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router /:idResolucion [post]
func (c *CancelacionValidaController) ValidarCancelacion() {
	idResolucionStr := c.Ctx.Input.Param(":idResolucion")
	vinculaciones := CargarVinculacionesDocente(idResolucionStr)
	c.Data["json"] = "OK"
	for _, vinculacion := range vinculaciones {
		if ExisteLiquidacion(vinculacion.NumeroContrato.String, strconv.FormatInt(vinculacion.Vigencia.Int64, 10)) {
			c.Data["json"] = "NO"
		}
	}
	c.ServeJSON()
}

func CargarVinculacionesDocente(idResolucion string) (c []models.VinculacionDocente) {
	var vinculacionesDocente []models.VinculacionDocente

	if err := getJson("http://10.20.0.254/administrativa_amazon_crud/v1/vinculacion_docente/?query=IdResolucion.Id%3A"+idResolucion+"&limit=0", &vinculacionesDocente); err == nil {
		fmt.Println(vinculacionesDocente)
	} else {
	}
	return vinculacionesDocente
}

func ExisteLiquidacion(numeroContrato string, vigencia string) (r bool) {
	var resultado bool
	var liquidaciones []models.DetalleLiquidacion
	fmt.Println("http://10.20.0.254/titan_api_crud/v1/detalle_liquidacion/?query=NumeroContrato.Id%3A" + numeroContrato + "%2CVigenciaContrato%3A" + vigencia)
	if err := getJson("http://10.20.0.254/titan_api_crud/v1/detalle_liquidacion/?query=NumeroContrato.Id%3A"+numeroContrato+"%2CVigenciaContrato%3A"+vigencia, &liquidaciones); err == nil {
		if len(liquidaciones) > 0 {
			resultado = true
		} else {
			resultado = false
		}
	} else {
		fmt.Println(err.Error())
	}
	return resultado
}
