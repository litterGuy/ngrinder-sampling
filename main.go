package main

import (
	"github.com/astaxie/beego"
	"ngrinder-sampling/models"
	_ "ngrinder-sampling/routers"
)

func main() {
	models.Init()

	beego.Run()
}
