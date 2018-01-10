package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CambioEstadoContratoValidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CambioEstadoContratoValidoController"],
		beego.ControllerComments{
			Method: "ValidarCambioEstado",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"],
		beego.ControllerComments{
			Method: "ValidarCancelacion",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"],
		beego.ControllerComments{
			Method: "GetContratoByContratoSuscritoId",
			Router: `GetContratoByContratoSuscritoId/:id/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:Contrato_generalController"],
		beego.ControllerComments{
			Method: "ListaContratoContratoSuscrito",
			Router: `ListaContratoContratoSuscrito/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ExpedirResolucionController"],
		beego.ControllerComments{
			Method: "Expedir",
			Router: `/expedir`,
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
			Method: "AnularDesvinculacionDocente",
			Router: `/anular_desvinculacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesDesvinculados",
			Router: `/docentes_desvinculados`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDesvinculacionesController"],
		beego.ControllerComments{
			Method: "AdicionarHoras",
			Router: `adicionar_horas`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionDisponibilidadController"],
		beego.ControllerComments{
			Method: "ListarApropiaciones",
			Router: `/listar_apropiaciones`,
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
			Method: "Calcular_total_de_salarios",
			Router: `Precontratacion/calcular_valor_contratos`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "ListarDocentesCargaHoraria",
			Router: `Precontratacion/docentes_x_carga_horaria`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionPrevinculacionesController"],
		beego.ControllerComments{
			Method: "InsertarPrevinculaciones",
			Router: `Precontratacion/insertar_previnculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:GestionResolucionesController"],
		beego.ControllerComments{
			Method: "InsertarResolucionCompleta",
			Router: `/insertar_resolucion_completa`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Contrato_proveedor",
			Router: `/contratoPersona`,
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
