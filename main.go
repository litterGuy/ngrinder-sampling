package main

import (
	"github.com/astaxie/beego"
	"ngrinder-sampling/models"
	_ "ngrinder-sampling/routers"
)

func main() {
	models.Init()
	beego.AddFuncMap("handleJS",handleJS)
	beego.Run()
}

func handleJS(in string)(out string){
	out = in + "world"
	return
}

