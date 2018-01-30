package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

// AprobacionPagoController operations for AprobacionPago
type AprobacionPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *AprobacionPagoController) URLMapping() {
	c.Mapping("ObtenerInfoCoordinador", c.ObtenerInfoCoordinador)

}

// GestionPrevinculacionesController ...
// @Title ObtenerInfoCoordinador
// @Description create ObtenerInfoCoordinador
// @Param id_dependencia_oikos query int true "Proyecto a obtener información coordinador"
// @Success 201 {int} models.InformacionCoordinador
// @Failure 403 :id_dependencia_oikos is empty
// @router /informacion_coordinador/:id_dependencia_oikos [get]
func (c *AprobacionPagoController) ObtenerInfoCoordinador() {
	id_oikos := c.GetString(":id_dependencia_oikos")
	var temp map[string]interface{}
	var temp_snies map[string]interface{}

	if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_homologacion_dependencias/proyecto_curricular_oikos/"+id_oikos, &temp); err == nil && temp != nil {
		json_proyecto_curricular, error_json := json.Marshal(temp)

		if error_json == nil {
			var temp_homologacion models.ObjetoProyectoCurricular
			json.Unmarshal(json_proyecto_curricular, &temp_homologacion)
			//fmt.Println(temp_homologacion.Homologacion.IDSnies)
			id_proyecto_snies := temp_homologacion.Homologacion.IDSnies
			fmt.Println("Este " + id_proyecto_snies)

			if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/academicaProxy/carrera_snies/918", &temp_snies); err == nil && temp_snies != nil {
				json_info_coordinador, error_json := json.Marshal(temp_snies)

				fmt.Println(json_info_coordinador)

				if error_json == nil {
					var temp_info_coordinador models.InformacionCoordinador
					json.Unmarshal(json_info_coordinador, &temp_info_coordinador)
					fmt.Println(temp_info_coordinador)
					//id_proyecto_snies = temp_homologacion.Homologacion.IDSnies

				} else {
					fmt.Println(error_json.Error())
					// c.Data["json"] = error_json.Error()
				}
			}

		} else {
			fmt.Println(error_json.Error())
			// c.Data["json"] = error_json.Error()
		}
	} else {
		fmt.Println(err)

	}

}
