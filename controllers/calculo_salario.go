package controllers

import (
	"fmt"
	"encoding/json"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	//. "github.com/mndrix/golog"
	"github.com/udistrital/administrativa_mid_api/models"
	. "github.com/udistrital/golog"
)

// PreliquidacionController operations for Preliquidacion
type CalculoSalarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *CalculoSalarioController) URLMapping() {
	//c.Mapping("CalcularSalarioContratacion", c.CalcularSalarioContratacion)
	c.Mapping("InsertarPrevinculaciones", c.InsertarPrevinculaciones)
}


// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router Contratacion/insertar_previnculaciones [post]
func (c *CalculoSalarioController) InsertarPrevinculaciones() {

	var v []models.VinculacionDocente
	var id_respuesta interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("docentes a contratar",v)
		v = CalcularSalarioPrecontratacion(v)

	} else {
		fmt.Println("ERROR")
		fmt.Println(err)

	}

	if err := sendJson("http://"+beego.AppConfig.String("UrlcrudArgo")+"/"+beego.AppConfig.String("NscrudArgo")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &id_respuesta, &v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = "Error al insertar docentes"
	}
	c.ServeJSON()
}


func CalcularSalarioPrecontratacion(docentes_a_vincular []models.VinculacionDocente)(docentes_a_insertar []models.VinculacionDocente) {
	//id_resolucion := 141
	nivel_academico := docentes_a_vincular[0].NivelAcademico
	var a string
	var categoria string


	for x, docente := range docentes_a_vincular {
		//docentes_a_vincular[x].NombreCompleto = docente.PrimerNombre + " " + docente.SegundoNombre + " " + docente.PrimerApellido + " " + docente.SegundoApellido
		//docentes_a_vincular[x].IdPersona = BuscarIdProveedor(docente.DocumentoIdentidad);

		if EsDocentePlanta(docente.IdPersona) && strings.ToLower(nivel_academico) == "posgrado" {
			categoria = categoria + "ud"
		} else {
			categoria = docente.Categoria
		}

		var predicados string
		if strings.ToLower(nivel_academico) == "posgrado" {
			predicados = `valor_salario_minimo(` + strconv.Itoa(CargarSalarioMinimo().Valor) + `,2016).` + "\n"
		} else if strings.ToLower(nivel_academico) == "pregrado" {
			predicados = `valor_punto(` + strconv.Itoa(CargarPuntoSalarial().ValorPunto) + `, 2016).` + "\n"
		}

		predicados = predicados + `categoria(` + docente.IdPersona + `,` + strings.ToLower(categoria) + `, 2016).` + "\n"
		predicados = predicados + `vinculacion(` + docente.IdPersona + `,` + strings.ToLower(docente.Dedicacion) + `, 2016).` + "\n"
		predicados = predicados + `horas(` + docente.IdPersona + `,` + strconv.Itoa(docente.NumeroHorasSemanales*docente.NumeroSemanas) + `, 2016).` + "\n"
		reglasbase := CargarReglasBase("CDVE")
		reglasbase = reglasbase + predicados
		m := NewMachine().Consult(reglasbase)

		contratos := m.ProveAll(`valor_contrato(` + strings.ToLower(nivel_academico) + `,` + docente.IdPersona + `,2016,X).`)
		for _, solution := range contratos {
			a = fmt.Sprintf("%s", solution.ByName_("X"))
		}
		f, _ := strconv.ParseFloat(a, 64)
		salario := f
		docentes_a_vincular[x].ValorContrato = salario

	}

	f, _ := strconv.ParseFloat(a, 64)
	salario := int(f)

	fmt.Println(salario)

	return docentes_a_vincular

}

func CargarPuntoSalarial() (p models.PuntoSalarial) {
	var v []models.PuntoSalarial

	if err := getJson("http://10.20.0.254/core_amazon_crud/v1/punto_salarial/?sortby=Vigencia&order=desc&limit=1", &v); err == nil {
	} else {
	}
	return v[0]
}

func CargarSalarioMinimo() (p models.SalarioMinimo) {
	var v []models.SalarioMinimo

	if err := getJson("http://10.20.0.254/core_amazon_crud/v1/salario_minimo/?sortby=Vigencia&order=desc&limit=1", &v); err == nil {
	} else {
	}
	return v[0]
}

func EsDocentePlanta(idPersona string) (docentePlanta bool) {
	var v []models.DocentePlanta
	if err := getJson("http://10.20.0.127/urano/index.php?data=B-7djBQWvIdLAEEycbH1n6e-3dACi5eLUOb63vMYhGq0kPBs7NGLYWFCL0RSTCu1yTlE5hH854MOgmjuVfPWyvdpaJDUOyByX-ksEPFIrrQQ7t1p4BkZcBuGD2cgJXeD&documento="+idPersona, &v); err == nil {
		fmt.Println(v[0].Nombres)
		return true
	} else {
		//fmt.Println("false")
		return false
	}
}

func BuscarIdProveedor(DocumentoIdentidad int)(id_proveedor_docente int){

		var id_proveedor int
		queryInformacionProveedor := "?query=NumDocumento:"+strconv.Itoa(DocumentoIdentidad)
		var informacion_proveedor []models.InformacionProveedor
		if err2 := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudArgo")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
			if(informacion_proveedor != nil){
				id_proveedor = informacion_proveedor[0].Id
			}else{
				id_proveedor = 0
			}

		}

		return id_proveedor
		//docentes_x_carga_horaria.CargasLectivas.CargaLectiva[x].IdProveedor = HomologarProyectoCurricular("old",pos.IDProyecto)

}
