package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"strings"
	"path/filepath"
	_ "github.com/udistrital/administrativa_mid_api/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestContratoPersona(t *testing.T) {
	query := "VigenciaContrato:2016"
	data := strings.NewReader(query)
	r, _ := http.NewRequest("POST", "http://localhost:8085/v1/informacion_proveedor/contratoPersona", data)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestContratoPersona", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Asunto: Envio y recibido de datos a contrato persona\n", t, func() {
		Convey("El codigo de estado debe ser 200", func() {
				So(w.Code, ShouldEqual, 200)
		})

		Convey("El resultado no deberia ser vacio", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
