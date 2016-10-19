package main

import (
	_ "hello/routers"

	"os"

	"github.com/astaxie/beego"
	"github.com/beego/beeweb/routers"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false

	setLog()
	//set static file access
	beego.SetStaticPath("/download", "download")

	//
	beego.Run()
}

func setLog() {
	routers.IsPro = beego.BConfig.RunMode == "prod"

	if routers.IsPro {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log.txt"}`)
	}
}

func enableXSRF() {
	//protect from xsrf
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEcGejJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 3600
}
