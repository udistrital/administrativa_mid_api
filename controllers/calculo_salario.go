package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	//. "github.com/mndrix/golog"
	."github.com/udistrital/golog"
	"github.com/udistrital/administrativa_mid_api/models"
)

// PreliquidacionController operations for Preliquidacion
type CalculoSalarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *CalculoSalarioController) URLMapping() {
	c.Mapping("CalcularSalarioContratacion", c.CalcularSalarioContratacion)
	c.Mapping("CalcularSalarioPrecontratacion", c.CalcularSalarioPrecontratacion)
}


// CalcularSalarioContratacion ...
// @Title CalcularSalarioContratacion
// @Description create CalcularSalarioContratacion
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router Contratacion/:idVinculacion [get]
func (c *CalculoSalarioController) CalcularSalarioContratacion() {
	idVinculacionStr := c.Ctx.Input.Param(":idVinculacion")
	fmt.Println(idVinculacionStr)
	vinculacionDocente := CargarVinculacionDocente(idVinculacionStr)
	fmt.Println("aca va el putazo de la primera funcion")
	fmt.Println(vinculacionDocente)
	escalafon := CargarEscalafon(strconv.Itoa(vinculacionDocente.IdPersona))
	fmt.Println("21132312132")
	fmt.Println(vinculacionDocente.IdPersona)
	fmt.Println("21132312132")
	if EsDocentePlanta(strconv.Itoa(vinculacionDocente.IdPersona)) && strings.ToLower(vinculacionDocente.IdResolucion.NivelAcademico) == "posgrado" {
		fmt.Println(EsDocentePlanta(strconv.Itoa(vinculacionDocente.IdPersona)))
		escalafon = escalafon + "ud"
	}

	predicados := `valor_punto(` + strconv.Itoa(CargarPuntoSalarial().ValorPunto) + `, 2016).` + "\n"
	predicados = predicados + `categoria(` + strconv.Itoa(vinculacionDocente.IdPersona) + `,` + strings.ToLower(escalafon) + `, 2016).` + "\n"
	fmt.Println("MADRAZOOOOOOOOOO ENTRE LOS PREDICADOS DE MIERDA")
	fmt.Println(vinculacionDocente.IdDedicacion.NombreDedicacion)
	predicados = predicados + `vinculacion(` + strconv.Itoa(vinculacionDocente.IdPersona) + `,` + strings.ToLower(vinculacionDocente.IdDedicacion.NombreDedicacion) + `,2016).` + "\n"
	predicados = predicados + `horas(` + strconv.Itoa(vinculacionDocente.IdPersona) + `,` + strconv.Itoa(vinculacionDocente.NumeroHorasSemanales*vinculacionDocente.NumeroSemanas) + `,2016).` + "\n"
	reglasbase := CargarReglasBase("CDVE")
	reglasbase = reglasbase + predicados
	//fmt.Println(reglasbase)
	m := NewMachine().Consult(reglasbase)
	var a string
	contratos := m.ProveAll(`valor_contrato(` + strings.ToLower(vinculacionDocente.IdResolucion.NivelAcademico) + `,` + strconv.Itoa(vinculacionDocente.IdPersona) + `,2016,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	f, _ := strconv.ParseFloat(a, 64)
	salario := int(f)
	c.Data["json"] = salario
	c.ServeJSON()

}

// CalcularSalarioPrecontratacion ...
// @Title CalcularSalarioPrecontratacion
// @Description create CalcularSalarioPrecontratacion
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router Precontratacion/:id_resolucion/:nivel_academico [post]
func (c *CalculoSalarioController) CalcularSalarioPrecontratacion() {
	id_resolucion:= c.Ctx.Input.Param(":id_resolucion")
	nivel_academico:= c.Ctx.Input.Param(":nivel_academico")

	var categoria string
	var docentes_precontratados []models.DocentePrecontratado
	fmt.Println(id_resolucion)

	if err := getJson("http://10.20.0.254/administrativa_amazon_api/v1/precontratado/"+id_resolucion, &docentes_precontratados); err == nil {

	}else {
		fmt.Println(err)
	}

	for x, docente := range docentes_precontratados {
		docentes_precontratados[x].NombreCompleto= docente.PrimerNombre + " " + docente.SegundoNombre + " " + docente.PrimerApellido + " " + docente.SegundoApellido;

		if EsDocentePlanta(strconv.Itoa(docente.Id)) && strings.ToLower(nivel_academico) == "posgrado" {
			categoria = categoria + "ud"
		}else{
			categoria = docente.Categoria
		}

		var predicados string
		if strings.ToLower(nivel_academico) == "posgrado" {
			predicados = `valor_salario_minimo(` + strconv.Itoa(CargarSalarioMinimo().Valor) + `,2016).` + "\n"
		} else if strings.ToLower(nivel_academico) == "pregrado" {
			predicados = `valor_punto(` + strconv.Itoa(CargarPuntoSalarial().ValorPunto) + `, 2016).` + "\n"
		}

		predicados = predicados + `categoria(` + strconv.Itoa(docente.Id) + `,` + strings.ToLower(categoria) + `, 2016).` + "\n"
		predicados = predicados + `vinculacion(` + strconv.Itoa(docente.Id) + `,` + strings.ToLower(docente.Dedicacion) + `, 2016).` + "\n"
		predicados = predicados + `horas(` + strconv.Itoa(docente.Id) + `,` + strconv.Itoa(docente.HorasSemanales*docente.Semanas) + `, 2016).` + "\n"
		reglasbase := CargarReglasBase("CDVE")
		reglasbase = reglasbase + predicados
		m := NewMachine().Consult(reglasbase)
		var a string
		contratos := m.ProveAll(`valor_contrato(` + strings.ToLower(nivel_academico) + `,` + strconv.Itoa(docente.Id) + `,2016,X).`)
		for _, solution := range contratos {
			a = fmt.Sprintf("%s", solution.ByName_("X"))
		}
		f, _ := strconv.ParseFloat(a, 64)
		salario := int(f)
		docentes_precontratados[x].ValorContrato = salario


	}

	c.Data["json"] = docentes_precontratados
	c.ServeJSON()

}

func CargarEscalafon(idPersona string) (e string) {
	escalafon := ""
	var v []models.CategoriaPersona

	if err := getJson("http://10.20.0.254/hvapi/v1/categoria_persona/?query=PersonaId%3A"+idPersona, &v); err == nil {
		escalafon = v[0].IdTipoCategoria.NombreCategoria
		fmt.Println(escalafon)
	} else {
		fmt.Println(err)
	}
	return escalafon
}

func CargarVinculacionDocente(idVinculacion string) (a models.VinculacionDocente) {
	var v []models.VinculacionDocente
	fmt.Println("putazo numero 2")
	fmt.Println(idVinculacion)

	if err := getJson("http://10.20.2.121:8080/v1/vinculacion_docente/?query=Id:"+idVinculacion, &v); err == nil {
		fmt.Println(v[0])
		fmt.Println("putazo if de error models")
		fmt.Println(v[0])
		return v[0]
	} else {
		fmt.Println("aca estoy escalafon gonoorea")
		fmt.Println(err.Error())
	}
	fmt.Println("putazo return models")
	fmt.Println(v)
	return
}

func CargarPuntoSalarial() (p models.PuntoSalarial) {
	var v []models.PuntoSalarial

	if err := getJson("http://10.20.0.254/core_api/v1/punto_salarial/?sortby=Vigencia&order=desc&limit=1", &v); err == nil {
	} else {
	}
	return v[0]
}

func CargarSalarioMinimo() (p models.SalarioMinimo) {
	var v []models.SalarioMinimo

	if err := getJson("http://10.20.0.254/core_api/v1/salario_minimo/?sortby=Vigencia&order=desc&limit=1", &v); err == nil {
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
