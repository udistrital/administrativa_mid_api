package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
	. "github.com/udistrital/golog"
)

// PreliquidacionController operations for Preliquidacion
type GestionPrevinculacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *GestionPrevinculacionesController) URLMapping() {
	c.Mapping("InsertarPrevinculaciones", c.InsertarPrevinculaciones)
	c.Mapping("CalcularTotalDeSalarios", c.CalcularTotalSalarios)
	c.Mapping("ListarDocentesCargaHoraria", c.ListarDocentesCargaHoraria)
}

// Calcular_total_de_salarios_seleccionados ...
// @Title Calcular_total_de_salarios_seleccionados
// @Description createCalcular_total_de_salarios_seleccionados
// @Success 201 {int} int
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos_seleccionados [post]
func (c *GestionPrevinculacionesController) Calcular_total_de_salarios_seleccionados() {

	var v []models.VinculacionDocente
	var total int
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}

	v, err = CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	total = int(CalcularTotalSalario(v))
	c.Data["json"] = total

	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/calcular_valor_contratos [post]
func (c *GestionPrevinculacionesController) CalcularTotalSalarios() {

	var v []models.VinculacionDocente
	var totalesDisponibilidad int
	var total int
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	v, err = CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	totalesSalario := CalcularTotalSalario(v)
	vigencia := strconv.Itoa(int(v[0].Vigencia.Int64))
	periodo := strconv.Itoa(v[0].Periodo)
	disponibilidad := strconv.Itoa(v[0].Disponibilidad)

	err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_valores_totales_x_disponibilidad/"+vigencia+"/"+periodo+"/"+disponibilidad+"", &totalesDisponibilidad)
	if err != nil {
		beego.Error("ERROR al calcular total de contratos", err)
		c.Abort("403")
	}
	total = int(totalesSalario) + totalesDisponibilidad
	c.Data["json"] = total

	c.ServeJSON()
}

// InsertarPrevinculaciones ...
// @Title InsetarPrevinculaciones
// @Description create InsertarPrevinculaciones
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /Precontratacion/insertar_previnculaciones [post]
func (c *GestionPrevinculacionesController) InsertarPrevinculaciones() {

	var v []models.VinculacionDocente
	var idRespuesta int

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		beego.Error("Error al hacer unmarshal", err)
		c.Abort("403")
	}
	v, err = CalcularSalarioPrecontratacion(v)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	err = sendJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/InsertarVinculaciones/", "POST", &idRespuesta, &v)
	c.Data["json"] = idRespuesta
	if err != nil {
		beego.Error("Error al insertar docentes", err)
		c.Abort("403")
	}

	c.ServeJSON()
}

// GestionPrevinculacionesController ...
// @Title ListarDocentesCargaHoraria
// @Description create ListarDocentesCargaHoraria
// @Param vigencia query string false "año a consultar"
// @Param periodo query string false "periodo a listar"
// @Param tipo_vinculacion query string false "vinculacion del docente"
// @Param facultad query string false "facultad"
// @Param nivel_academico query string false "nivel_academico"
// @Success 201 {object} models.Docentes_x_Carga
// @Failure 403 body is empty
// @router /Precontratacion/docentes_x_carga_horaria [get]
func (c *GestionPrevinculacionesController) ListarDocentesCargaHoraria() {

	vigencia := c.GetString("vigencia")
	periodo := c.GetString("periodo")
	tipoVinculacion := c.GetString("tipo_vinculacion")
	facultad := c.GetString("facultad")
	nivelAcademico := c.GetString("nivel_academico")

	docentesXcargaHoraria, err := ListarDocentesHorasLectivas(vigencia, periodo, tipoVinculacion, facultad, nivelAcademico)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}
	newDocentesXcargaHoraria := models.ObjetoCargaLectiva{}

	//BUSCAR CATEGORÍA DE CADA DOCENTE
	for _, pos := range docentesXcargaHoraria.CargasLectivas.CargaLectiva {
		catDocente := models.ObjetoCategoriaDocente{}
		emptyCatDocente := models.ObjetoCategoriaDocente{}
		q := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudUrano") + "/categoria_docente/" + vigencia + "/" + periodo + "/" + pos.DocDocente
		err = getXml(q, &catDocente.CategoriaDocente)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}

		pos.CategoriaNombre, pos.IDCategoria, err = Buscar_Categoria_Docente(vigencia, periodo, pos.DocDocente)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		if catDocente.CategoriaDocente != emptyCatDocente.CategoriaDocente {
			newDocentesXcargaHoraria.CargasLectivas.CargaLectiva = append(newDocentesXcargaHoraria.CargasLectivas.CargaLectiva, pos)
		}
	}

	//RETORNAR CON ID DE TIPO DE VINCULACION DE NUEVO MODELO
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDTipoVinculacion, pos.NombreTipoVinculacion = HomologarDedicacion_ID("old", pos.IDTipoVinculacion)
		if pos.IDTipoVinculacion == "3" {
			pos.HorasLectivas = "20"
			pos.NombreTipoVinculacion = "MTO"
		}
		if pos.IDTipoVinculacion == "4" {
			pos.HorasLectivas = "40"
			pos.NombreTipoVinculacion = "TCO"
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}

	//RETORNAR FACULTTADES CON ID DE OIKOS, HOMOLOGACION
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDFacultad, err = HomologarFacultad("old", pos.IDFacultad)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos
	}
	//RETORNAR PROYECTOS CURRICUALRES HOMOLOGADOS!!
	for x, pos := range newDocentesXcargaHoraria.CargasLectivas.CargaLectiva {
		pos.IDProyecto, err = HomologarProyectoCurricular(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		pos.DependenciaAcademica, err = strconv.Atoi(pos.IDProyecto)
		if err != nil {
			beego.Error(err)
			c.Abort("403")
		}
		newDocentesXcargaHoraria.CargasLectivas.CargaLectiva[x] = pos

	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = newDocentesXcargaHoraria.CargasLectivas.CargaLectiva
	c.ServeJSON()

}

func CalcularSalarioPrecontratacion(docentes_a_vincular []models.VinculacionDocente) (docentes_a_insertar []models.VinculacionDocente, err error) {
	nivelAcademico := docentes_a_vincular[0].NivelAcademico
	vigencia := strconv.Itoa(int(docentes_a_vincular[0].Vigencia.Int64))
	var a string
	var categoria string

	for x, docente := range docentes_a_vincular {
		p, err := EsDocentePlanta(docente.IdPersona)
		if err != nil {
			return docentes_a_insertar, err
		}
		if p && strings.ToLower(nivelAcademico) == "posgrado" {
			categoria = docente.Categoria + "ud"
		} else {
			categoria = docente.Categoria
		}

		var predicados string
		if strings.ToLower(nivelAcademico) == "posgrado" {
			salarioMinimo, err := CargarSalarioMinimo()
			if err != nil {
				return docentes_a_insertar, err
			}
			predicados = "valor_salario_minimo(" + strconv.Itoa(salarioMinimo.Valor) + "," + vigencia + ")." + "\n"
			docente.NumeroSemanas = 1
		} else if strings.ToLower(nivelAcademico) == "pregrado" {
			a, err := CargarPuntoSalarial()
			if err != nil {
				return docentes_a_insertar, err
			}
			predicados = "valor_punto(" + strconv.Itoa(a.ValorPunto) + ", " + vigencia + ")." + "\n"
		}

		predicados = predicados + "categoria(" + docente.IdPersona + "," + strings.ToLower(categoria) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "vinculacion(" + docente.IdPersona + "," + strings.ToLower(docente.Dedicacion) + ", " + vigencia + ")." + "\n"
		predicados = predicados + "horas(" + docente.IdPersona + "," + strconv.Itoa(docente.NumeroHorasSemanales*docente.NumeroSemanas) + ", " + vigencia + ")." + "\n"
		reglasbase := CargarReglasBase("CDVE")
		reglasbase = reglasbase + predicados
		m := NewMachine().Consult(reglasbase)

		contratos := m.ProveAll("valor_contrato(" + strings.ToLower(nivelAcademico) + "," + docente.IdPersona + "," + vigencia + ",X).")
		for _, solution := range contratos {
			a = fmt.Sprintf("%s", solution.ByName_("X"))
		}
		f, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return docentes_a_vincular, err
		}
		salario := f
		docentes_a_vincular[x].ValorContrato = salario

	}

	return docentes_a_vincular, nil

}

func CargarPuntoSalarial() (p models.PuntoSalarial, err error) {
	var v []models.PuntoSalarial

	err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/punto_salarial/?sortby=Vigencia&order=desc&limit=1", &v)
	if err != nil {
		err = fmt.Errorf("He fallado en punto_salarial (get) función CargarPuntoSalarial, %s", err)
	}
	return v[0], err
}

func CargarSalarioMinimo() (p models.SalarioMinimo, err error) {
	var v []models.SalarioMinimo

	err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/salario_minimo/?sortby=Vigencia&order=desc&limit=1", &v)
	if err != nil {
		err = fmt.Errorf("He fallado en salario_minimo (get) función CargarSalarioMinimo, %s", err)
	}

	return v[0], err
}

func EsDocentePlanta(idPersona string) (docentePlanta bool, err error) {
	var temp map[string]interface{}
	var esDePlanta bool

	err = getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAcademica")+"/"+"consultar_datos_docente/"+idPersona, &temp)
	if err != nil {
		esDePlanta = false
		return false, err
	}
	jsonDocentes, err := json.Marshal(temp)
	if err != nil {
		return false, err
	}

	var tempDocentes models.ObjetoDocentePlanta
	json.Unmarshal(jsonDocentes, &tempDocentes)

	if tempDocentes.DocenteCollection.Docente[0].Planta == "true" {
		esDePlanta = true
	} else {
		esDePlanta = false
	}

	return esDePlanta, nil
}

func BuscarIdProveedor(DocumentoIdentidad int) (id_proveedor_docente int) {

	var Idproveedor int
	queryInformacionProveedor := "?query=NumDocumento:" + strconv.Itoa(DocumentoIdentidad)
	var informacionProveedor []models.InformacionProveedor
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/"+queryInformacionProveedor, &informacionProveedor); err2 == nil {
		if informacionProveedor != nil {
			Idproveedor = informacionProveedor[0].Id
		} else {
			Idproveedor = 0
		}

	}

	return Idproveedor
}

func CalcularTotalSalario(v []models.VinculacionDocente) (total float64) {

	var sumatoria float64
	for _, docente := range v {
		sumatoria = sumatoria + docente.ValorContrato
	}

	return sumatoria
}

// tipos de vinculación
const (
	tipoVinculacion = iota + 1
	tipoCancelacion
	tipoAdicion
	tipoReduccion
)

//ESTA FUNCIÓN LISTA LOS DOCENTES PREVINCULADOS EN TRUE O FALSE

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculadosAll
// @Description create ListarDocentesPrevinculadosAll
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados_all [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculadosAll() {
	idResolucion := c.GetString("id_resolucion")
	var v = []models.VinculacionDocente{}
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion
	var vinc models.VinculacionDocente

	//If resoluciones (GET)
	err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+idResolucion, &res)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	var llenarVinculacion = func(v *models.VinculacionDocente) {
		documentoIdentidad, _ := strconv.Atoi(v.IdPersona)

		v.NombreCompleto = BuscarNombreProveedor(documentoIdentidad)
		v.NumeroDisponibilidad = BuscarNumeroDisponibilidad(v.Disponibilidad)
		v.Dedicacion = BuscarNombreDedicacion(v.IdDedicacion.Id)
		v.LugarExpedicionCedula = BuscarLugarExpedicion(v.IdPersona)
		v.TipoDocumento = BuscarTipoDocumento(v.IdPersona)
		v.NumeroHorasSemanales, v.ValorContrato = Calcular_totales_vinculacion_pdf(v.IdPersona, idResolucion)
		v.NumeroMeses = strconv.FormatFloat(float64(v.NumeroSemanas)/4, 'f', 2, 64) + " meses"
		v.ValorContratoFormato = FormatMoney(int(v.ValorContrato), 2)
	}
	switch res.IdTipoResolucion.Id {
	case tipoVinculacion:
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/get_vinculaciones_agrupadas/"+idResolucion, &v)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

		for x, pos := range v {
			llenarVinculacion(&pos)
			v[x] = pos
		}

	case tipoCancelacion:
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		//If 2 modificacion_vinculacion (get)
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		//for vinculaciones
		for x, pos := range modvin {
			//If 3 vinculacion_docente para el join (get)
			err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.VinculacionDocenteCancelada.Id), &vinc)
			if err != nil {
				beego.Error(err)
				c.Abort("400")
			}
			v = append(v, vinc)
			llenarVinculacion(&vinc)
			v[x] = vinc
		}
	default:
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+idResolucion, &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		//for vinculaciones
		for x, pos := range modvin {
			//If 3 vinculacion_docente para el join (get)
			err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.VinculacionDocenteRegistrada.Id), &vinc)
			if err != nil {
				beego.Error(err)
				c.Abort("400")
			}
			v = append(v, vinc)
			llenarVinculacion(&vinc)
			v[x] = vinc
		}
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}

//ESTA FUNCIÓN LISTA LOS DOCENTES PREVINCULADOS EN TRUE

// GestionPrevinculacionesController ...
// @Title ListarDocentesPrevinculados
// @Description create ListarDocentesPrevinculados
// @Param id_resolucion query string false "resolucion a consultar"
// @Success 201 {int} models.VinculacionDocente
// @Failure 403 body is empty
// @router /docentes_previnculados [get]
func (c *GestionPrevinculacionesController) ListarDocentesPrevinculados() {
	idResolucion, err := c.GetInt("id_resolucion")
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	query := "?limit=-1&query=IdResolucion.Id:" + strconv.Itoa(idResolucion) + ",Estado:true"
	var v = []models.VinculacionDocente{}
	var res models.Resolucion
	var modres []models.ModificacionResolucion
	var modvin []models.ModificacionVinculacion

	err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/resolucion/"+strconv.Itoa(idResolucion), &res)
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	beego.Debug(res.IdTipoResolucion)
	switch res.IdTipoResolucion.Id {
	case tipoVinculacion:
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &v)

		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

	case tipoCancelacion:
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+strconv.Itoa(idResolucion), &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}

		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		arreglo := make([]string, len(modvin))
		for x, pos := range modvin {
			arreglo[x] = strconv.Itoa(pos.VinculacionDocenteCancelada.Id)
		}
		identificadoresvinc := strings.Join(arreglo, "|")
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:False,Id__in:"+identificadoresvinc+"&limit=-1", &v)

		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
	default: //los otros tipos
		err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+strconv.Itoa(idResolucion), &modres)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modres[0].Id), &modvin)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		arreglo := make([]string, len(modvin))
		for x, pos := range modvin {
			arreglo[x] = strconv.Itoa(pos.VinculacionDocenteRegistrada.Id)
		}
		identificadoresvinc := strings.Join(arreglo, "|")
		err = getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?query=Estado:True,Id__in:"+identificadoresvinc+"&limit=-1", &v)
		if err != nil {
			beego.Error(err)
			c.Abort("400")
		}
		for x, pos := range v {
			documentoIdentidad, _ := strconv.Atoi(pos.IdPersona)
			v[x].NombreCompleto = BuscarNombreProveedor(documentoIdentidad)
			v[x].NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
			v[x].Dedicacion = BuscarNombreDedicacion(pos.IdDedicacion.Id)
			v[x].LugarExpedicionCedula = BuscarLugarExpedicion(pos.IdPersona)
			v[x].TipoDocumento = BuscarTipoDocumento(pos.IdPersona)
			v[x].ValorContratoFormato = FormatMoney(int(v[x].ValorContrato), 2)
		}

	}
	for x, pos := range v {
		documentoIdentidad, _ := strconv.Atoi(pos.IdPersona)

		pos.NombreCompleto = BuscarNombreProveedor(documentoIdentidad)
		pos.NumeroDisponibilidad = BuscarNumeroDisponibilidad(pos.Disponibilidad)
		pos.Dedicacion = BuscarNombreDedicacion(pos.IdDedicacion.Id)
		pos.LugarExpedicionCedula = BuscarLugarExpedicion(pos.IdPersona)
		pos.TipoDocumento = BuscarTipoDocumento(pos.IdPersona)
		pos.ValorContratoFormato = FormatMoney(int(v[x].ValorContrato), 2)

		v[x] = pos
	}
	if v == nil {
		v = []models.VinculacionDocente{}
	}
	c.Ctx.Output.SetStatus(201)

	c.Data["json"] = v
	c.ServeJSON()

}

func ListarDocentesHorasLectivas(vigencia, periodo, tipo_vinculacion, facultad, nivel_academico string) (docentes_a_listar models.ObjetoCargaLectiva, err error) {

	tipoVinculacionOld := HomologarDedicacion_nombre(tipo_vinculacion)
	facultadOld, err := HomologarFacultad("new", facultad)
	if err != nil {
		return docentes_a_listar, err
	}

	var temp map[string]interface{}
	var docentesXCarga models.ObjetoCargaLectiva

	for _, pos := range tipoVinculacionOld {
		t := "http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAcademica") + "/" + "carga_lectiva/" + vigencia + "/" + periodo + "/" + pos + "/" + facultadOld + "/" + nivel_academico

		err = getJsonWSO2(t, &temp)
		if err != nil {
			return docentesXCarga, err
		}
		jsonDocentes, err := json.Marshal(temp)
		if err != nil {
			return docentesXCarga, err
		}

		var tempDocentes models.ObjetoCargaLectiva
		json.Unmarshal(jsonDocentes, &tempDocentes)
		docentesXCarga.CargasLectivas.CargaLectiva = append(docentesXCarga.CargasLectivas.CargaLectiva, tempDocentes.CargasLectivas.CargaLectiva...)

	}

	return docentesXCarga, nil

}

func Buscar_Categoria_Docente(vigencia, periodo, documento_ident string) (categoria_nombre, categoria_id_old string, err error) {
	var temp map[string]interface{}
	var nombreCategoria string
	var idCategoriaOld string

	err = getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudUrano")+"/"+"categoria_docente/"+vigencia+"/"+periodo+"/"+documento_ident, &temp)
	if err != nil {
		return categoria_nombre, categoria_id_old, err
	}
	if temp != nil {

		jsonDocentes, err := json.Marshal(temp)

		if err != nil {
			return categoria_nombre, categoria_id_old, err
		}
		var tempDocentes models.ObjetoCategoriaDocente
		json.Unmarshal(jsonDocentes, &tempDocentes)
		nombreCategoria = tempDocentes.CategoriaDocente.Categoria
		idCategoriaOld = tempDocentes.CategoriaDocente.IDCategoria

	}
	return nombreCategoria, idCategoriaOld, nil
}

func HomologacionTotal() {

}

func HomologarProyectoCurricular(proyecto_old string) (proyecto string, err error) {
	var id_proyecto string
	var temp map[string]interface{}

	err = getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+"proyecto_curricular_cod_proyecto/"+proyecto_old, &temp)
	if err != nil {
		return proyecto, err
	}

	json_proyecto_curricular, err := json.Marshal(temp)

	if err != nil {
		return proyecto, err
	}
	var temp_proy models.ObjetoProyectoCurricular
	err = json.Unmarshal(json_proyecto_curricular, &temp_proy)
	if err != nil {
		return proyecto, err
	}
	id_proyecto = temp_proy.Homologacion.IDOikos

	return id_proyecto, nil
}

func HomologarFacultad(tipo, facultad string) (facultad_old string, err error) {
	var id_facultad string
	var temp map[string]interface{}
	var string_consulta_servicio string

	if tipo == "new" {
		string_consulta_servicio = "facultad_gedep_oikos"
	} else {
		string_consulta_servicio = "facultad_oikos_gedep"
	}

	err = getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudHomologacion")+"/"+string_consulta_servicio+"/"+facultad, &temp)
	if err != nil {
		return facultad_old, err
	}
	if temp != nil {
		json_facultad, err := json.Marshal(temp)

		if err != nil {
			return facultad_old, err
		}

		var temp_proy models.ObjetoFacultad
		json.Unmarshal(json_facultad, &temp_proy)

		if tipo == "new" {
			id_facultad = temp_proy.Homologacion.IdGeDep
		} else {
			id_facultad = temp_proy.Homologacion.IdOikos
		}

	} else {
		return id_facultad, fmt.Errorf("No hay datos de respuesta de las APIs")
	}

	return id_facultad, nil

}

func HomologarDedicacion_nombre(dedicacion string) (vinculacion_old []string) {
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

	for _, pos := range arreglo_homologacion {
		if pos.Nombre == dedicacion {
			id_dedicacion_old = append(id_dedicacion_old, pos.Old)
		}
	}

	return id_dedicacion_old
}

func HomologarDedicacion_ID(tipo, dedicacion string) (vinculacion_old, nombre_vinculacion string) {
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
		panic(err) //nunca esperado
	}

	for _, pos := range arreglo_homologacion {
		if tipo == "new" {
			comparacion = pos.New
			resultado = pos.Old
		} else {
			comparacion = pos.Old
			resultado = pos.New
		}

		if comparacion == dedicacion {
			id_dedicacion_old = resultado
			nombre_dedicacion = pos.Nombre
		}
	}

	return id_dedicacion_old, nombre_dedicacion
}

func BuscarNombreProveedor(DocumentoIdentidad int) (nombre_prov string) {

	var nom_proveedor string
	queryInformacionProveedor := "?query=NumDocumento:" + strconv.Itoa(DocumentoIdentidad)
	var informacion_proveedor []models.InformacionProveedor
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/"+queryInformacionProveedor, &informacion_proveedor); err2 == nil {
		if informacion_proveedor != nil {
			nom_proveedor = informacion_proveedor[0].NomProveedor
		} else {
			nom_proveedor = ""
		}

	}

	return nom_proveedor

}

func BuscarTipoDocumento(Cedula string) (nombre_tipo_doc string) {
	var tipo_documento string
	var temp []models.InformacionPersonaNatural
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural/?limit=-1&query=Id:"+Cedula, &temp); err2 == nil {
		if temp != nil {
			tipo_documento = temp[0].TipoDocumento.ValorParametro
		} else {
			tipo_documento = "N/A"
		}
	} else {
		fmt.Println("error en json", err2)
		tipo_documento = "N/A"
	}

	return tipo_documento

}

func BuscarNombreDedicacion(id_dedicacion int) (nombre_dedicacion string) {
	var nom_dedicacion string
	query := "?limit=-1&query=Id:" + strconv.Itoa(id_dedicacion)
	var dedicaciones []models.Dedicacion
	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/dedicacion"+query, &dedicaciones); err2 == nil {
		if dedicaciones != nil {
			nom_dedicacion = dedicaciones[0].Descripcion
		} else {
			nom_dedicacion = ""
		}

	}

	return nom_dedicacion
}

func BuscarNumeroDisponibilidad(IdCDP int) (numero_disp int) {

	var temp []models.Disponibilidad
	var numero_disponibilidad int
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudKronos")+"/"+beego.AppConfig.String("NscrudKronos")+"/disponibilidad?limit=-1&query=DisponibilidadApropiacion.Id:"+strconv.Itoa(IdCDP), &temp); err == nil {
		if temp != nil {
			numero_disponibilidad = int(temp[0].NumeroDisponibilidad)

		} else {
			numero_disponibilidad = 0
		}

	} else {
		fmt.Println("Error en disponibilidad (get) función BuscarNumeroDisponibilidad:", err)
	}
	return numero_disponibilidad

}

func BuscarLugarExpedicion(Cedula string) (nombre_lugar_exp string) {

	var nombre_ciudad string
	var temp []models.InformacionPersonaNatural
	var temp2 []models.Ciudad

	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_persona_natural/?limit=-1&query=Id:"+Cedula, &temp); err2 == nil {
		if temp != nil {
			id_ciudad := temp[0].IdCiudadExpedicionDocumento
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ciudad/?limit=-1&query=Id:"+strconv.Itoa(int(id_ciudad)), &temp2); err2 == nil {
				if temp2 != nil {
					nombre_ciudad = temp2[0].Nombre

				} else {
					nombre_ciudad = "N/A"
				}

			} else {
				fmt.Println("error en json", err)
			}

		} else {
			nombre_ciudad = "N/A"
		}

	} else {
		fmt.Println("error en json", err2)
	}

	return nombre_ciudad

}

func Calcular_totales_vinculacion_pdf(cedula, id_resolucion string) (suma_total_horas int, suma_total_contrato float64) {

	query := "?limit=-1&query=IdPersona:" + cedula + ",IdResolucion.Id:" + id_resolucion
	var temp []models.VinculacionDocente
	var total_contrato int
	var total_horas int

	if err2 := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente"+query, &temp); err2 == nil {

		for _, pos := range temp {
			total_horas = total_horas + pos.NumeroHorasSemanales
			total_contrato = total_contrato + int(pos.ValorContrato)
		}

	} else {
		fmt.Println("error al guardar en json")
		total_horas = 0
		total_contrato = 0
	}

	return total_horas, float64(total_contrato)
}

// GestionPrevinculacionesController ...
// @Title GetVinculacionesAgrupadasCanceladas
// @Description Get de vinculaciones agrupadas canceladas para el PDF
// @Param id_resolucion 	path 	string	true  "resolucion a consultar"
// @Success 201 {object} []models.VinculacionDocente
// @Failure 403 :id_resolucion is empty
// @router /vinculaciones_agrupadas_canceladas/:id_resolucion [get]
func (c *GestionDesvinculacionesController) GetVinculacionesAgrupadasCanceladas() {
	id_resolucion := c.Ctx.Input.Param(":id_resolucion")

	var modRes []models.ModificacionResolucion
	var modVin []models.ModificacionVinculacion
	var v []models.VinculacionDocente
	var vinc models.VinculacionDocente

	//If 1 modificacion_resolucion (get)
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_resolucion/?query=ResolucionNueva:"+id_resolucion, &modRes); err == nil {
		//If 2 modificacion_vinculacion (get)
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/modificacion_vinculacion/?query=ModificacionResolucion:"+strconv.Itoa(modRes[0].Id), &modVin); err == nil {

			//for vinculaciones
			for _, pos := range modVin {
				//If 3 vinculacion_docente para el join (get)
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(pos.VinculacionDocenteCancelada.Id), &vinc); err == nil {
					v = append(v, vinc)
				} else { //If 3 vinculacion_docente para el join (get)
					fmt.Println("He fallado en If 3 vinculacion_docente para el join (get), solucioname!!!", err)
				}
			}

		} else { //If 2 modificacion_vinculacion (get)
			fmt.Println("He fallado en If 2 modificacion_vinculacion (get), solucioname!!!", err)
		}
	} else { //If 1 modificacion_resolucion (get)
		fmt.Println("He fallado en If 1 modificacion_resolucion (get), solucioname!!!", err)
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()

}
