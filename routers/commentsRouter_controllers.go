package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "CalcularSalarioContratacion",
			Router: `Contratacion/:idVinculacion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"],
		beego.ControllerComments{
			Method: "ValidarCancelacion",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"],
		beego.ControllerComments{
			Method: "ValidarContrato",
			Router: `/:dedicacion/:numHoras`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
