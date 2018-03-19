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
	//"github.com/udistrital/auditoria"
)

func init() {

	//auditoria.InitMiddleware()
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/gestion_previnculacion",
			beego.NSInclude(
				&controllers.GestionPrevinculacionesController{},
			),
		),

		beego.NSNamespace("/gestion_resoluciones",
			beego.NSInclude(
				&controllers.GestionResolucionesController{},
			),
		),

		beego.NSNamespace("/gestion_documento_resolucion",
			beego.NSInclude(
				&controllers.GestionDocumentoResolucionController{},
			),
		),

		beego.NSNamespace("/gestion_desvinculaciones",
			beego.NSInclude(
				&controllers.GestionDesvinculacionesController{},
			),
		),

		beego.NSNamespace("/consultar_disponibilidades",
			beego.NSInclude(
				&controllers.GestionDisponibilidadController{},
			),
		),

		beego.NSNamespace("/expedir_resolucion",
			beego.NSInclude(
				&controllers.ExpedirResolucionController{},
			),
		),

		beego.NSNamespace("/validar_contrato",
			beego.NSInclude(
				&controllers.ValidarContratoController{},
			),
		),

		beego.NSNamespace("/validarCambioEstado",
			beego.NSInclude(
				&controllers.CambioEstadoContratoValidoController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.Contrato_generalController{},
			),
		),

		beego.NSNamespace("/aprobacion_pago",
			beego.NSInclude(
				&controllers.AprobacionPagoController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
