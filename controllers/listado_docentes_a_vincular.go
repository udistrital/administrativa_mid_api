package controllers

import (
	"fmt"
	"strconv"
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
// @Title ListarDocentesPrevinculados
// @Description create ListarDocentesPrevinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados [get]
func (c *ListarDocentesVinculacionController) ListarDocentesPrevinculados(){
	id_resolucion := c.GetString("id_resolucion")
	fmt.Println("resolucion a consultar")
	fmt.Println(id_resolucion)
	query := "?limit=-1&query=IdResolucion.Id:"+id_resolucion
	var v []models.VinculacionDocente

	if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudArgo")+"/"+beego.AppConfig.String("NscrudArgo")+"/vinculacion_docente"+query, &v); err2 == nil {
		for x, pos := range  v{
			documento_identidad,_ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto = BuscarNombreProveedor(documento_identidad);
			v[x].NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
		}

	}else{
		fmt.Println(err2)
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = 	v
	c.ServeJSON()
  //fmt.Println(v)

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

	docentes_x_carga_horaria := ListarDocentesHorasLectivas(vigencia, periodo, tipo_vinculacion, facultad)

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for x, pos := range  docentes_x_carga_horaria.CargasLectivas.CargaLectiva {
		docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].CategoriaNombre,docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDCategoria  = Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
	}

	//RETORNAR CON ID DE TIPO DE VINCULACION DE NUEVO MODELO
	for x, pos := range  docentes_x_carga_horaria.CargasLectivas.CargaLectiva {
		docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDTipoVinculacion, docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].NombreTipoVinculacion  = HomologarDedicacion_ID("old",pos.IDTipoVinculacion)
		if (docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDTipoVinculacion == "3"){
			docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].HorasLectivas	= "20"
			docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].NombreTipoVinculacion = "MTO"
		}
		if(docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDTipoVinculacion == "4"){
			docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].HorasLectivas	= "40"
			docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].NombreTipoVinculacion = "TCO"
		}
	}

	//RETORNAR FACULTTADES CON ID DE OIKOS, HOMOLOGACION
	for x, pos := range  docentes_x_carga_horaria.CargasLectivas.CargaLectiva {
		docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDFacultad = HomologarFacultad("old",pos.IDFacultad)
	}

	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range  docentes_x_carga_horaria.CargasLectivas.CargaLectiva {
		docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IDProyecto = HomologarProyectoCurricular("old",pos.IDProyecto)
	}


	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = 	docentes_x_carga_horaria .CargasLectivas.CargaLectiva
	c.ServeJSON()

}

func ListarDocentesHorasLectivas(vigencia, periodo, tipo_vinculacion, facultad string)(docentes_a_listar models.ObjetoCargaLectiva){

	tipo_vinculacion_old := HomologarDedicacion_nombre(tipo_vinculacion)
	facultad_old := HomologarFacultad("new",facultad)

	var temp map[string]interface{}
	var docentes_x_carga models.ObjetoCargaLectiva

	for _, pos := range  tipo_vinculacion_old {
		if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_academicos.HTTPEndpoint/carga_lectiva/"+vigencia+"/"+periodo+"/"+pos+"/"+facultad_old, &temp); err == nil && temp != nil{
		 jsonDocentes, error_json := json.Marshal(temp)

		 if error_json == nil {
			 var temp_docentes models.ObjetoCargaLectiva
			 json.Unmarshal(jsonDocentes, &temp_docentes)
			 docentes_x_carga.CargasLectivas.CargaLectiva = append(docentes_x_carga.CargasLectivas.CargaLectiva, temp_docentes.CargasLectivas.CargaLectiva...)
			 //c.Ctx.Output.SetStatus(201)
			 //c.Data["json"] = docentes_a_listar.CargasLectivas.CargaLectiva
		 }else{
			// c.Data["json"] = error_json.Error()
		 }
	 }else {
		 fmt.Println(err)

	 }
	 }

	 return docentes_x_carga;

}

func Buscar_Categoria_Docente(vigencia, periodo, documento_ident string)(categoria_nombre, categoria_id_old string){
	var temp map[string]interface{}
	var nombre_categoria string
	var id_categoria_old string
	//*****ojo, está quemada la cédula por falta de datos*****
	if err := getJsonWSO2("http://jbpm.udistritaloas.edu.co:8280/services/servicios_urano_pruebas/categoria_docente/"+vigencia+"/1/79708124", &temp); err == nil && temp != nil{
	 jsonDocentes, error_json := json.Marshal(temp)

	 if error_json == nil {
		 var temp_docentes models.ObjetoCategoriaDocente
		 json.Unmarshal(jsonDocentes, &temp_docentes)
		 nombre_categoria = temp_docentes.CategoriaDocente.Categoria
		 id_categoria_old = temp_docentes.CategoriaDocente.IDCategoria

	 }else{
		 fmt.Println(error_json.Error())
		// c.Data["json"] = error_json.Error()
	 }
 }else {
	 fmt.Println(err)

 }

 return nombre_categoria,id_categoria_old
}

func HomologacionTotal(){

}

func HomologarProyectoCurricular(tipo, proyecto string)(proyecto_old string){
	var id_proyecto_old string
	var comparacion string
	var resultado string
	homologacion_proyectos := `[
						{
							"old": "20",
							"new": "72"
						},
						{
							"old": "25",
							"new": "70"
						},
						{
							"old": "7",
							"new": "73"
						},
						{
							"old": "5",
							"new": "74"
						},
						{
							"old": "15",
							"new": "79"
						}
						]`

	 byt := []byte(homologacion_proyectos)
	 var arreglo_homologacion []models.Homologacion
	 if err := json.Unmarshal(byt, &arreglo_homologacion); err != nil {
			 panic(err)
	 }


	 for _, pos := range  arreglo_homologacion{
		 	if(tipo == "new"){
				comparacion = pos.New
				resultado = pos.Old
			}else{
				comparacion = pos.Old
				resultado = pos.New
			}

			if(comparacion == proyecto){
				id_proyecto_old = resultado
		}
 	}

	return id_proyecto_old
}

func HomologarFacultad(tipo, facultad string)(facultad_old string){

	var id_facultad_old string
	var comparacion string
	var resultado string
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
		 	if(tipo == "new"){
				comparacion = pos.New
				resultado = pos.Old
			}else{
				comparacion = pos.Old
				resultado = pos.New
			}

			if(comparacion == facultad){
				id_facultad_old = resultado
		}
 	}

	return id_facultad_old
}

func HomologarDedicacion_nombre(dedicacion string)(vinculacion_old []string){
	var id_dedicacion_old []string
	homologacion_dedicacion := `[
						{
							"nombre": "HCH",
							"old": "5",
							"new": "1"
						},
						{
							"nombre": "HCP",
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


	return id_dedicacion_old
}

func HomologarDedicacion_ID(tipo,dedicacion string)(vinculacion_old, nombre_vinculacion string){
	var id_dedicacion_old string
	var nombre_dedicacion string
	var comparacion string
	var resultado string
	homologacion_dedicacion := `[
						{
							"nombre": "HCH",
							"old": "5",
							"new": "1"
						},
						{
							"nombre": "HCP",
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
					 if(tipo == "new"){
					 comparacion = pos.New
					 resultado = pos.Old
				 }else{
					 comparacion = pos.Old
					 resultado = pos.New
				 }

				 if(comparacion == dedicacion){
					 id_dedicacion_old = resultado
					 nombre_dedicacion = pos.Nombre
			 }
 	}


	return id_dedicacion_old, nombre_dedicacion
}

func BuscarNombreProveedor(DocumentoIdentidad int)(nombre_prov string){

		var nom_proveedor string
		queryInformacionProveedor := "?query=NumDocumento:"+strconv.Itoa(DocumentoIdentidad)
		var informacion_proveedor []models.InformacionProveedor
		if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudArgo")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
			if(informacion_proveedor != nil){
				nom_proveedor = informacion_proveedor[0].NomProveedor
			}else{
				nom_proveedor = ""
			}

		}

		return nom_proveedor
		//docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IdProveedor = HomologarProyectoCurricular("old",pos.IDProyecto)

}


func BuscarNumeroDisponibilidad(IdCDP int)(numero_disp int){

		var temp []models.Disponibilidad
		var numero_disponibilidad int
		if err2 := getJson("http://10.20.0.254/financiera_api/v1/disponibilidad?limit=-1&query=Id:"+strconv.Itoa(IdCDP), &temp); err2 == nil {
			if(temp != nil){
				numero_disponibilidad = int(temp[0].NumeroDisponibilidad)
			
			}else{
				numero_disponibilidad = 0
			}

		}else{
			fmt.Println("error en json",err2)
		}
		return numero_disponibilidad
		//docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IdProveedor = HomologarProyectoCurricular("old",pos.IDProyecto)

}
