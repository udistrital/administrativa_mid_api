// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),

		beego.NSNamespace("/calculo_salario",
			beego.NSInclude(
				&controllers.CalculoSalarioController{},
			),
		),

		beego.NSNamespace("/consultar_disponibilidades",
			beego.NSInclude(
				&controllers.VerificarDisponibilidadesController{},
			),
		),

		beego.NSNamespace("/validar_contrato",
			beego.NSInclude(
				&controllers.ValidarContratoController{},
			),
		),

		beego.NSNamespace("/cancelacion_valida",
			beego.NSInclude(
				&controllers.CancelacionValidaController{},
			),
		),

		beego.NSNamespace("/validarCambioEstado",
			beego.NSInclude(
				&controllers.CambioEstadoContratoValidoController{},
			),
		),

		beego.NSNamespace("/informacionDocentes",
			beego.NSInclude(
				&controllers.ListarDocentesVinculacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
