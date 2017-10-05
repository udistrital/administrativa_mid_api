package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	. "github.com/mndrix/golog"
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
	escalafon := CargarEscalafon(strconv.Itoa(vinculacionDocente.IdPersona))
	fmt.Println(escalafon)
	if EsDocentePlanta(strconv.Itoa(vinculacionDocente.IdPersona)) && strings.ToLower(vinculacionDocente.IdResolucion.NivelAcademico) == "posgrado" {
		escalafon = escalafon + "ud"
	}

	predicados := `valor_punto(` + strconv.Itoa(CargarPuntoSalarial().ValorPunto) + `, 2016).` + "\n"
	predicados = predicados + `categoria(` + strconv.Itoa(vinculacionDocente.IdPersona) + `,` + strings.ToLower(escalafon) + `, 2016).` + "\n"
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
// @router Precontratacion/:nivelAcademico/:idProfesor/:numHoras/:numSemanas/:categoria/:dedicacion [post]
func (c *CalculoSalarioController) CalcularSalarioPrecontratacion() {
	nivelAcademico := c.Ctx.Input.Param(":nivelAcademico")
	idPersonaStr := c.Ctx.Input.Param(":idProfesor")
	numHorasStr := c.Ctx.Input.Param(":numHoras")
	numHoras, _ := strconv.Atoi(numHorasStr)
	numSemanasStr := c.Ctx.Input.Param(":numSemanas")
	numSemanas, _ := strconv.Atoi(numSemanasStr)
	categoria := c.Ctx.Input.Param(":categoria")
	vinculacion := c.Ctx.Input.Param(":dedicacion")
	if EsDocentePlanta(idPersonaStr) && strings.ToLower(nivelAcademico) == "posgrado" {
		categoria = categoria + "ud"
	}	

	var predicados string
	if strings.ToLower(nivelAcademico) == "posgrado" {
		predicados = `valor_salario_minimo(` + strconv.Itoa(CargarSalarioMinimo().Valor) + `,2016).` + "\n"
	} else if strings.ToLower(nivelAcademico) == "pregrado" {
		predicados = `valor_punto(` + strconv.Itoa(CargarPuntoSalarial().ValorPunto) + `, 2016).` + "\n"
	}
	predicados = predicados + `categoria(` + idPersonaStr + `,` + strings.ToLower(categoria) + `, 2016).` + "\n"
	predicados = predicados + `vinculacion(` + idPersonaStr + `,` + strings.ToLower(vinculacion) + `, 2016).` + "\n"
	predicados = predicados + `horas(` + idPersonaStr + `,` + strconv.Itoa(numHoras*numSemanas) + `, 2016).` + "\n"
	reglasbase := CargarReglasBase("CDVE")
	reglasbase = reglasbase + predicados
	m := NewMachine().Consult(reglasbase)
	var a string
	contratos := m.ProveAll(`valor_contrato(` + strings.ToLower(nivelAcademico) + `,` + idPersonaStr + `,2016,X).`)
	for _, solution := range contratos {
		a = fmt.Sprintf("%s", solution.ByName_("X"))
	}
	f, _ := strconv.ParseFloat(a, 64)
	salario := int(f)
	c.Data["json"] = salario
	c.ServeJSON()

}

func CargarEscalafon(idPersona string) (e string) {
	escalafon := ""
	var v []models.CategoriaPersona

	if err := getJson("http://10.20.0.254/hvapi/v1/categoria_persona/?query=PersonaId%3A"+idPersona, &v); err == nil {
		escalafon = v[0].IdTipoCategoria.NombreCategoria
	} else {
	}
	return escalafon
}

func CargarVinculacionDocente(idVinculacion string) (a models.VinculacionDocente) {
	var v []models.VinculacionDocente

	fmt.Println(idVinculacion)

	if err := getJson("http://10.20.0.254/administrativa_api/v1/vinculacion_docente/?query=Id%3A"+idVinculacion, &v); err == nil {
		fmt.Println(v)
		return v[0]
	} else {
		fmt.Println(err.Error())
	}
	return models.VinculacionDocente{}
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
	var v bool

	if err := getJson("http://10.20.2.17:8085/v1/docente_planta/"+idPersona, &v); err == nil {
	} else {
	}
	return v
}
