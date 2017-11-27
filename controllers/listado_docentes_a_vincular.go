package controllers

import (
	"fmt"
	//"strconv"
	//"strings"
	"encoding/json"
	"github.com/astaxie/beego"
	//. "github.com/mndrix/golog"
	"github.com/udistrital/administrativa_mid_api/models"
	//. "github.com/udistrital/golog"
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

	tipo_vinculacion_old := HomologarDedicacion(tipo_vinculacion)
	facultad_old := HomologarFacultad(facultad)

	var temp map[string]interface{}
	var docentes_a_listar models.ObjetoCargaLectiva

	for _, pos := range  tipo_vinculacion_old {
		if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_academicos.HTTPEndpoint/carga_lectiva/"+vigencia+"/"+periodo+"/"+pos+"/"+facultad_old, &temp); err == nil && temp != nil{
		 jsonDocentes, error_json := json.Marshal(temp)

		 if error_json == nil {
			 var temp_docentes models.ObjetoCargaLectiva
			 json.Unmarshal(jsonDocentes, &temp_docentes)
			 docentes_a_listar.CargasLectivas.CargaLectiva = append(docentes_a_listar.CargasLectivas.CargaLectiva, temp_docentes.CargasLectivas.CargaLectiva...)
			 c.Ctx.Output.SetStatus(201)
			 c.Data["json"] = docentes_a_listar.CargasLectivas.CargaLectiva
		 }else{
			 c.Data["json"] = error_json.Error()
		 }
	 }else {
		 fmt.Println(err)

	 }
	 }




	c.ServeJSON()

}

func HomologarFacultad(facultad string)(facultad_old string){

	var id_facultad_old string
	homologacion_facultad := `[
						{
							"old": "33",
							"new": "14"
						},
						{
							"old": "24",
							"new": "17"
						},
						{
							"old": "23",
							"new": "35"
						},
						{
							"old": "101",
							"new": "65"
						},
						{
							"old": "32",
							"new": "66"
						}
						]`

	 byt := []byte(homologacion_facultad)
	 var arreglo_homologacion []models.Homologacion
	 if err := json.Unmarshal(byt, &arreglo_homologacion); err != nil {
			 panic(err)
	 }

	 for _, pos := range  arreglo_homologacion{
			if(pos.New == facultad){
				id_facultad_old = pos.Old
		}
 	}
	fmt.Println("resultado homologacion facultad")
	fmt.Println(id_facultad_old)
	return id_facultad_old
}

func HomologarDedicacion(dedicacion string)(vinculacion_old []string){
	var id_dedicacion_old []string
	homologacion_dedicacion := `[
						{
							"nombre": "HCP",
							"old": "5",
							"new": "1"
						},
						{
							"nombre": "HCS",
							"old": "4",
							"new": "2"
						},
						{
							"nombre": "TCO|MTO",
							"old": "2",
							"new": "4"
						},{
							"nombre": "TCO|MTO",
							"old": "3",
							"new": "3"
						}
						]`

	 byt := []byte(homologacion_dedicacion)
	 var arreglo_homologacion []models.HomologacionDedicacion
	 if err := json.Unmarshal(byt, &arreglo_homologacion); err != nil {
			 panic(err)
	 }

	 for _, pos := range  arreglo_homologacion{
			if(pos.Nombre == dedicacion){
				id_dedicacion_old = append(id_dedicacion_old, pos.Old)
		}
 	}

	fmt.Println("resultado homologacion dedicacion")
	fmt.Println(id_dedicacion_old)
	return id_dedicacion_old
}

//err := json.Unmarshal(jsonDocentes, &docentes_a_listar)
		//fmt.Println(err)
		//fmt.Println(docentes_a_listar)
