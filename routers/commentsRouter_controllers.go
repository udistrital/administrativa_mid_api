package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "Calcular_total_de_salarios",
			Router: `Contratacion/calcular_valor_contratos`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "InsertarPrevinculaciones",
			Router: `Contratacion/insertar_previnculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

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

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Contrato_proveedor",
			Router: `/contratoPersona`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ListarDocentesVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ListarDocentesVinculacionController"],
		beego.ControllerComments{
			Method: "ListarDocentesPrevinculados",
			Router: `/docentes_previnculados`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ListarDocentesVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ListarDocentesVinculacionController"],
		beego.ControllerComments{
			Method: "ListarDocentesCargaHoraria",
			Router: `/docentes_x_carga_horaria`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"],
		beego.ControllerComments{
			Method: "ValidarContrato",
			Router: `/:dedicacion/:numHoras`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:VerificarDisponibilidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:VerificarDisponibilidadesController"],
		beego.ControllerComments{
			Method: "ListarApropiaciones",
			Router: `/listar_apropiaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
