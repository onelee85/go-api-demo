package main

import (
	_ "user/db"
	_ "user/routers"
	_ "user/util"

	"github.com/astaxie/beego"
)

func main() {
	beego.Informational("server run  mode:", beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
