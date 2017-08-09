package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Contrato_proveedor",
			Router: `/contrato_proveedor`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})
<<<<<<< HEAD

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "CalcularSalarioPrecontratacion",
			Router: `Precontratacion/:nivelAcademico/:idProfesor/:numHoras/:numSemanas/:categoria/:dedicacion`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CalculoSalarioController"],
		beego.ControllerComments{
			Method: "CalcularSalarioContratacion",
			Router: `Contratacion/:idVinculacion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:ValidarContratoController"],
		beego.ControllerComments{
			Method: "ValidarContrato",
			Router: `/:dedicacion/:numHoras`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_mid_api/controllers:CancelacionValidaController"],
		beego.ControllerComments{
			Method: "ValidarCancelacion",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

=======
>>>>>>> issue/2
}
