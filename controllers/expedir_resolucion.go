package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/administrativa_mid_api/models"
)

// ExpedirResolucionController operations for ExpedirResolucion
type ExpedirResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExpedirResolucionController) URLMapping() {
	c.Mapping("Expedir", c.Expedir)
}

// Expedir ...
// @Title Expedir
// @Description create Expedir
// @Success 201 {int} models.ExpedicionResolucion
// @Failure 403 body is empty
// @router /expedir [post]
func (c *ExpedirResolucionController) Expedir() {
	amazon := orm.NewOrm()
	flyway := orm.NewOrm()
	amazon.Using("amazonAdmin")
	flyway.Using("flywayAdmin")
	var m models.ExpedicionResolucion
	var temp string
	var cdve int
	var proveedor []models.InformacionProveedor
	var response interface{}
	//alerta = append(alerta, "success")
	vigencia, _, _ := time.Now().Date()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		v := m.Vinculaciones
		//numeroContratos := NumeroContratoDVE(vigencia)
		if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_general/maximo_dve", &cdve); err == nil {
			fmt.Println("Consulta exitosa consecutivo DVE")
		} else {
			fmt.Println("Error consecutivo DVE: ", err)
		}
		fmt.Println("He fallado un poquito aaaaaaaa, solucioname!!!", err)
		numeroContratos := cdve
		for _, vinculacion := range *v {
			fmt.Println("He fallado un poquito aaaaaaaa2, solucioname!!!", err)
			numeroContratos = numeroContratos + 1
			v := vinculacion.VinculacionDocente
			fmt.Println("He fallado un poquito aaaaaaaa3, solucioname!!!", err)
			fmt.Println("He fallado un poquito aaaaaaaa4 con el v, solucioname!!!", v)
			//if err = flyway.Read(&v); err == nil {
			fmt.Println("He fallado un poquito aaaaaaaa5, solucioname!!!", err)
			if v.NumeroContrato.String == "" && v.Vigencia.Int64 == 0 {
				contrato := vinculacion.ContratoGeneral
				acta := vinculacion.ActaInicio
				fmt.Println(contrato.Contratista)
				aux1 := 181
				fmt.Println("He fallado un poquito aux11111111111111111111111111111111111111111111, solucioname!!!", err)
				contrato.VigenciaContrato = vigencia
				fmt.Println("He fallado un poquito vigenciaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa, solucioname!!!", err)
				contrato.Id = "DVE" + strconv.Itoa(numeroContratos)
				fmt.Println("He fallado un poquito contrato idddddddddddddddddddddddddddddddddddddddddddddddddddd, solucioname!!!", err)
				contrato.FormaPago.Id = 240
				fmt.Println("He fallado un poquito FORMAPAGOOOOOOOOOOOOOOOOOOOOOOOO, solucioname!!!", err)
				contrato.DescripcionFormaPago = "Abono a Cuenta Mensual de acuerdo a puntos y horas laboradas"
				contrato.Justificacion = "Docente de Vinculacion Especial"
				contrato.UnidadEjecucion.Id = 205
				contrato.LugarEjecucion.Id = 4
				//contrato.Supervisor.Id = 192
				contrato.TipoControl = aux1
				contrato.ClaseContratista = 33
				contrato.TipoMoneda = 137
				contrato.OrigenRecursos = 149
				contrato.OrigenPresupueso = 156
				contrato.TemaGastoInversion = 166
				contrato.TipoGasto = 146
				contrato.RegimenContratacion = 136
				contrato.Procedimiento = 132
				contrato.ModalidadSeleccion = 123
				contrato.TipoCompromiso = 35
				contrato.TipologiaContrato = 46
				fmt.Println("He fallado un poquito aquiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii, solucioname!!!", err)
				contrato.FechaRegistro = time.Now()
				fmt.Println("He fallado un poquito tiempoooooooooooooooooooooooooooooooo, solucioname!!!", err)
				contrato.UnidadEjecutora = 1
				contrato.Condiciones = "Sin condiciones"
				fmt.Println("He fallado un poquito porqueeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee, solucioname!!!", err)
				fmt.Println("He fallado un poquito supervisorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr, solucioname!!!", err)
				fmt.Println("He fallado un poquito aaaaaaaa6, solucioname!!!", err)

				if err := getJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(contrato.Contratista), &proveedor); err == nil {
					fmt.Println("Consulta exitosa informacion proveedor")
					fmt.Println("He fallado un poquito aaaaaaaa7, solucioname!!!", err)
				} else {
					fmt.Println("Error informacion proveedor")
					fmt.Println("He fallado un poquito aaaaaaaa8, solucioname!!!", err)
				}
				fmt.Println("He fallado un poquito aaaaaaaa9, solucioname!!!", err)
				temp = proveedor[0].NumDocumento

				_, err = amazon.Raw("INSERT INTO argo.contrato_general(numero_contrato, vigencia, objeto_contrato, plazo_ejecucion, forma_pago, ordenador_gasto, sede_solicitante, dependencia_solicitante, contratista, unidad_ejecucion, valor_contrato, justificacion, descripcion_forma_pago, condiciones, unidad_ejecutora, fecha_registro, tipologia_contrato, tipo_compromiso, modalidad_seleccion, procedimiento, regimen_contratacion, tipo_gasto, tema_gasto_inversion, origen_presupueso, origen_recursos, tipo_moneda, tipo_control, observaciones, supervisor,clase_contratista, tipo_contrato, lugar_ejecucion) VALUES (?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", contrato.Id, contrato.VigenciaContrato, contrato.ObjetoContrato, contrato.PlazoEjecucion, contrato.FormaPago.Id, contrato.OrdenadorGasto, contrato.SedeSolicitante, contrato.DependenciaSolicitante, temp[0], contrato.UnidadEjecucion.Id, contrato.ValorContrato, contrato.Justificacion, contrato.DescripcionFormaPago, contrato.Condiciones, contrato.UnidadEjecutora, contrato.FechaRegistro.Format(time.RFC1123), contrato.TipologiaContrato, contrato.TipoCompromiso, contrato.ModalidadSeleccion, contrato.Procedimiento, contrato.RegimenContratacion, contrato.TipoGasto, contrato.TemaGastoInversion, contrato.OrigenPresupueso, contrato.OrigenRecursos, contrato.TipoMoneda, contrato.TipoControl, contrato.Observaciones, 192, contrato.ClaseContratista, contrato.TipoContrato, contrato.LugarEjecucion.Id).Exec()
				fmt.Println("He fallado un poquito aaaaaaaa10, solucioname!!!", err)
				fmt.Println("Consulta realizada")
				if err == nil {
					fmt.Println("He fallado un poquito aaaaaaaa11, solucioname!!!", err)
					aux1 := contrato.Id
					aux2 := contrato.VigenciaContrato
					var ce models.ContratoEstado
					var ec models.EstadoContrato
					ce.NumeroContrato = aux1
					fmt.Println("He fallado un poquito aaaaaaaa12, solucioname!!!", err)
					ce.Vigencia = aux2
					fmt.Println("He fallado un poquito aaaaaaaa13, solucioname!!!", err)
					ce.FechaRegistro = time.Now()
					ec.Id = 4
					ce.Estado = &ec

					if err2 := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_estado", "POST", &response, &ce); err2 == nil {
						fmt.Println("Mirenme, soy un post de contrato_estado :D", response)
					} else {
						fmt.Println("He fallado un poquito en contrato_estado, solucioname!!!", err2)
						amazon.Rollback()
						flyway.Rollback()
						return
					}

					if err == nil {
						a := vinculacion.VinculacionDocente
						fmt.Println("Desde aquiiiiiiiiiiiiii la aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa ", vinculacion.VinculacionDocente)
						var ai models.ActaInicio
						ai.NumeroContrato = aux1
						ai.Vigencia = aux2
						ai.Descripcion = acta.Descripcion
						ai.FechaInicio = acta.FechaInicio
						ai.FechaFin = acta.FechaFin
						fmt.Println(a.NumeroSemanas, "semanasaaaaaaaa")
						ai.FechaFin = CalcularFechaFin(acta.FechaInicio, a.NumeroSemanas)

						if err3 := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio", "POST", &response, &ai); err3 == nil {
							fmt.Println("Mirenme, soy un post de acta_inicio :D", response)
						} else {
							fmt.Println("He fallado un poquito en acta_inicio, solucioname!!!", err3)
							amazon.Rollback()
							flyway.Rollback()
							return
						}

						var cd models.ContratoDisponibilidad
						cd.NumeroCdp = vinculacion.Cdp.NumeroCdp
						cd.NumeroContrato = aux1
						cd.Vigencia = aux2
						cd.Estado = true
						cd.FechaRegistro = time.Now()
						cd.VigenciaCdp = vinculacion.Cdp.VigenciaCdp

						if err4 := sendJson("http://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad", "POST", &response, &cd); err4 == nil {
							fmt.Println("Mirenme, soy un post de contrato_disponibilidad :D", response)
						} else {
							fmt.Println("He fallado un poquito en contrato_disponibilidad, solucioname!!!", err4)
							amazon.Rollback()
							flyway.Rollback()
							return
						}

						//ACA
						if err == nil {
							//if err = flyway.Read(&a); err == nil {
							a.IdPuntoSalarial = vinculacion.VinculacionDocente.IdPuntoSalarial
							a.IdSalarioMinimo = vinculacion.VinculacionDocente.IdSalarioMinimo
							v := a
							v.NumeroContrato.String = aux1
							v.Vigencia.Int64 = int64(aux2)
							fmt.Println("Soyyyyyyyyyyyy la aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa ", a)
							fmt.Println("Soyyyyyyyyyyyy la vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv ", v)
							fmt.Println("Soyyyyyyyyyyyy la idresolucion ", a.IdResolucion)

							if err := sendJson("http://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/"+strconv.Itoa(v.Id), "PUT", &response, &v); err == nil {
								fmt.Println("Mirenme, soy un put de vinculacion_docente :D ", response)
							} else {
								fmt.Println("He fallado un poquito en vinculacion_docente, solucioname!!! ", err)
								amazon.Rollback()
								flyway.Rollback()
								return
								//	}
							} /*else {
								fmt.Println("He fallado un poquito leyendo a (vinculacion_docente), solucioname!!!", err)
								amazon.Rollback()
								flyway.Rollback()
								return
							}*/
						} //xD

						//Aca
					} else {
						fmt.Println("He fallado antes de ir a la vinculacion docente, solucioname!!!", err)
						amazon.Rollback()
						flyway.Rollback()
						return
					}
				} else {
					fmt.Println("He fallado un poquito leyendo a (vinculacion_docente), solucioname!!!", err)
					amazon.Rollback()
					flyway.Rollback()
					return
				}
			} else {
				aux1 := v.NumeroContrato
				aux2 := v.Vigencia
				var ce models.ContratoEstado
				ce.NumeroContrato = aux1.String
				ce.Vigencia = int(aux2.Int64)
				ce.FechaRegistro = time.Now()
				ce.Estado.Id = 1
				//_, err = o.Insert(&e)
				if err != nil {
					fmt.Println("He fallado un poquito en el segundo estado_contrato post, solucioname!!!", err)
					amazon.Rollback()
					flyway.Rollback()
					return
				}
			}
			//} //Cierre if Read V (vinculacion.vinculacionDocente)
		} //Cierre for de v (m.vinculacion)

		c.Data["json"] = v
	} else {
		/*fmt.Println("ERROR")
		fmt.Println(err)
		c.Data["json"] = "Error al listar disponibilidades"*/
	}

	c.ServeJSON()
}

/*func NumeroContratoDVE(vigencia int) (consecutivo int) {
	o := orm.NewOrm()
	var temp []models.TotalContratos
	_, err := o.Raw("SELECT COALESCE(MAX(substring(numero_contrato,4)::INTEGER), 0) total FROM argo.contrato_general WHERE numero_contrato LIKE 'DVE%' AND vigencia = " + strconv.Itoa(vigencia) + ";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp[0].NumeroTotal
}*/

func CalcularFechaFin(fecha_inicio time.Time, numero_semanas int) (fecha_fin time.Time) {

	semanas := float32(numero_semanas)
	meses := semanas / 4
	fmt.Println("meses", meses)
	numero_dias := (meses * 30) + 1
	f_i := fecha_inicio
	after := f_i.AddDate(0, 0, int(numero_dias))
	fmt.Println(after)
	return after

}
