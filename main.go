package main

import (
	_ "github.com/udistrital/administrativa_mid_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	"github.com/udistrital/utils_oas/apiStatusLib"
)

func init() {
	orm.RegisterDataBase("amazonAdmin", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAgora")+":"+beego.AppConfig.String("PasscrudAgora")+"@"+beego.AppConfig.String("HostcrudAgora")+"/"+beego.AppConfig.String("BdcrudAgora")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAgora")+"")
	orm.RegisterDataBase("flywayAdmin", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAdmin")+":"+beego.AppConfig.String("PasscrudAdmin")+"@"+beego.AppConfig.String("HostcrudAdmin")+"/"+beego.AppConfig.String("BdcrudAdmin")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAdmin")+"")
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAgora")+":"+beego.AppConfig.String("PasscrudAgora")+"@"+beego.AppConfig.String("HostcrudAgora")+"/"+beego.AppConfig.String("BdcrudAgora")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAgora")+"")
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Custom JSON error pages
	beego.ErrorHandler("400", BadRequestJsonPage)
	beego.ErrorHandler("403", forgivenJsonPage)
	beego.ErrorHandler("404", notFoundJsonPage)

	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/administrativa_mid_api.log"}`)

	apistatus.Init()
	beego.Run()
}
