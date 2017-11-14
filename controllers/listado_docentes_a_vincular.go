package controllers

import (
	"fmt"
	//"strconv"
	//"strings"
	"encoding/json"
	"github.com/astaxie/beego"
	//. "github.com/mndrix/golog"
	"github.com/udistrital/administrativa_mid_api/models"
	. "github.com/udistrital/golog"
)

//ListarDocentesVinculacionController operations for Preliquidacion
type ListarDocentesVinculacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ListarDocentesVinculacionController) URLMapping() {
	c.Mapping("ListarDocentesCargaHoraria", c.ListarDocentesCargaHoraria)
	
}

// ListarDocentesVinculacionController ...
// @Title ListarDocentesCargaHoraria
// @Description create ListarDocentesCargaHoraria
// @Param vigencia query string false "año a consultar"
// @Param periodo query string false "periodo a listar"
// @Param tipo_vinculacion query string false "vinculacion del docente"
// @Param facultad query string false "facultad"
// @Success 201 {object} models.Docentes_x_Carga
// @Failure 403 body is empty
// @router /docentes_x_carga_horaria [get]
func (c *ListarDocentesVinculacionController) ListarDocentesCargaHoraria() {
	vigencia := c.GetString("vigencia")
	periodo := c.GetString("periodo")
	tipo_vinculacion := c.GetString("tipo_vinculacion")
	facultad := c.GetString("facultad")

	fmt.Println(vigencia)
	fmt.Println(periodo)
	fmt.Println(tipo_vinculacion)
	fmt.Println(facultad)
	var temp map[string]interface{}
	var docentes_a_listar models.ObjetoCargaLectiva

	if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_academicos.HTTPEndpoint/carga_lectiva/"+vigencia+"/"+periodo+"/"+tipo_vinculacion+"/"+facultad, &temp); err == nil && temp != nil{
		jsonDocentes, error_json := json.Marshal(temp)
		fmt.Println(jsonDocentes)
		if error_json == nil {
			err := json.Unmarshal(jsonDocentes, &docentes_a_listar)
			fmt.Println(err)
			fmt.Println(docentes_a_listar)
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = docentes_a_listar
		}else{
			c.Data["json"] = error_json.Error()
		}
	}else {
		fmt.Println(err)
		
	}

	c.ServeJSON()

}

//err := json.Unmarshal(jsonDocentes, &docentes_a_listar)
		//fmt.Println(err)
		//fmt.Println(docentes_a_listar)