package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "CertificacionDocumentosAprobados",
			Router: `/certificacion_documentos_aprobados/:dependencia/:mes/:anio`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "CertificacionVistoBueno",
			Router: `/certificacion_visto_bueno/:dependencia/:mes/:anio`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "ObtenerDependenciaOrdenador",
			Router: `/dependencia_ordenador/:docordenador`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "GetContratosDocente",
			Router: `/get_contratos_docente/:numDocumento`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "ObtenerInfoCoordinador",
			Router: `/informacion_coordinador/:id_dependencia_oikos`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "ObtenerInfoOrdenador",
			Router: `/informacion_ordenador/:numero_contrato/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "PagoAprobado",
			Router: `/pago_aprobado/:numero_contrato/:vigencia/:mes/:anio`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "GetSolicitudesCoordinador",
			Router: `/solicitudes_coordinador/:doccoordinador`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "GetSolicitudesOrdenador",
			Router: `/solicitudes_ordenador/:docordenador`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method: "GetSolicitudesSupervisor",
			Router: `/solicitudes_supervisor/:docsupervisor`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CambioEstadoContratoValidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CambioEstadoContratoValidoController"],
		beego.ControllerComments{
			Method: "ValidarCambioEstado",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"],
		beego.ControllerComments{
			Method: "GetContratoByContratoSuscritoId",
			Router: `/GetContratoByContratoSuscritoId/:id/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"],
		beego.ControllerComments{
			Method: "ListaContratoContratoSuscrito",
			Router: `/ListaContratoContratoSuscrito/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method: "Cancelar",
			Router: `/cancelar`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method: "Expedir",
			Router: `/expedir`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method: "ValidarDatosExpedicion",
			Router: `/validar_datos_expedicion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "ActualizarVinculaciones",
			Router: `/actualizar_vinculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "AdicionarHoras",
			Router: `/adicionar_horas`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "AnularAdicionDocente",
			Router: `/anular_adicion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "AnularDesvinculacionDocente",
			Router: `/anular_desvinculacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesCancelados",
			Router: `/docentes_cancelados`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesDesvinculados",
			Router: `/docentes_desvinculados`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDisponibilidadController"],
		beego.ControllerComments{
			Method: "ListarApropiaciones",
			Router: `/listar_apropiaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDocumentoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDocumentoResolucionController"],
		beego.ControllerComments{
			Method: "GetContenidoResolucion",
			Router: `/get_contenido_resolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "Calcular_total_de_salarios",
			Router: `/Precontratacion/calcular_valor_contratos`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "Calcular_total_de_salarios_seleccionados",
			Router: `/Precontratacion/calcular_valor_contratos_seleccionados`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesCargaHoraria",
			Router: `/Precontratacion/docentes_x_carga_horaria`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "InsertarPrevinculaciones",
			Router: `/Precontratacion/insertar_previnculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesPrevinculados",
			Router: `/docentes_previnculados`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesPrevinculadosAll",
			Router: `/docentes_previnculados_all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"],
		beego.ControllerComments{
			Method: "GetResolucionesAprobadas",
			Router: `/get_resoluciones_aprobadas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"],
		beego.ControllerComments{
			Method: "GetResolucionesInscritas",
			Router: `/get_resoluciones_inscritas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"],
		beego.ControllerComments{
			Method: "InsertarResolucionCompleta",
			Router: `/insertar_resolucion_completa`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"],
		beego.ControllerComments{
			Method: "ValidarContrato",
			Router: `/:dedicacion/:numHoras`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
