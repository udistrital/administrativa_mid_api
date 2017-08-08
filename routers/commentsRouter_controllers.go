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
}
