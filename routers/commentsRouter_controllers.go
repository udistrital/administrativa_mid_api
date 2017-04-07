package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "ContratoPersona",
			Router: `/contratoPersona`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
