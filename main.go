package main

import (
	"net/url"
	"time"

	_ "github.com/udistrital/administrativa_mid_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/database"
	"github.com/udistrital/utils_oas/security"
	"github.com/udistrital/utils_oas/xray"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	flyway, err := database.BuildPostgresConnectionString()
	if err != nil {
		logs.Error("error consultando la cadena de conexión: %v", err)
		panic(err)
	}

	err = orm.RegisterDataBase("flywayAdmin", "postgres", flyway+"&timezone=UTC")
	if err != nil {
		logs.Error("error al conectarse a la base de datos: %v", err)
		panic(err)
	}

	amazon := "postgres://" + beego.AppConfig.String("UsercrudAgora") + ":" + url.QueryEscape(beego.AppConfig.String("PasscrudAgora")) + "@" + beego.AppConfig.String("HostcrudAgora") + "/" + beego.AppConfig.String("BdcrudAgora") + "?sslmode=disable&search_path=" + beego.AppConfig.String("SchcrudAgora") + "&timezone=UTC"
	if err := orm.RegisterDataBase("default", "postgres", amazon); err != nil {
		panic(err)
	}
}

func main() {
	allowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == beego.DEV {
		allowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Custom JSON error pages
	beego.ErrorHandler("400", badRequestJSONPage)
	beego.ErrorHandler("403", forgivenJSONPage)
	beego.ErrorHandler("404", notFoundJSONPage)
	beego.ErrorHandler("233", notValidJSONPage)

	apistatus.Init()
	auditoria.InitMiddleware()
	security.SetSecurityHeaders()
	xray.Init()

	beego.Run()
}
