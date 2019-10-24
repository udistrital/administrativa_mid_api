package main

import (
	_ "github.com/udistrital/administrativa_mid_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

func init() {
	// orm.DefaultTimeLoc = time.UTC
	orm.Debug = true
	amazon := "postgres://" + beego.AppConfig.String("UsercrudAgora") + ":" + beego.AppConfig.String("PasscrudAgora") + "@" + beego.AppConfig.String("HostcrudAgora") + "/" + beego.AppConfig.String("BdcrudAgora") + "?sslmode=disable&search_path=" + beego.AppConfig.String("SchcrudAgora") + "&timezone=UTC"
	flyway := "postgres://" + beego.AppConfig.String("UsercrudAdmin") + ":" + beego.AppConfig.String("PasscrudAdmin") + "@" + beego.AppConfig.String("HostcrudAdmin") + "/" + beego.AppConfig.String("BdcrudAdmin") + "?sslmode=disable&search_path=" + beego.AppConfig.String("SchcrudAdmin") + "&timezone=UTC"

	if err := orm.RegisterDataBase("amazonAdmin", "postgres", amazon); err != nil {
		panic(err)
	}

	if err := orm.RegisterDataBase("flywayAdmin", "postgres", flyway); err != nil {
		panic(err)
	}

	if err := orm.RegisterDataBase("default", "postgres", amazon); err != nil {
		panic(err)
	}
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
	beego.ErrorHandler("400", badRequestJSONPage)
	beego.ErrorHandler("403", forgivenJSONPage)
	beego.ErrorHandler("404", notFoundJSONPage)
	beego.ErrorHandler("233", notValidJSONPage)

	if err := logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/administrativa_mid_api.log"}`); err != nil {
		beego.Info(err)
	}

	apistatus.Init()
	beego.ErrorController(&customerror.CustomErrorController{})
	beego.Run()
}
