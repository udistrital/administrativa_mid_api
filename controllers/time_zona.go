package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func tiempo_bogota() time.Time {
	var tiempoBogota = time.Now()
	logs.Info(tiempoBogota)

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)
	tiempoBogota = tiempoBogota.In(loc)
	return tiempoBogota
}
