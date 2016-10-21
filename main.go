package main

import (
	_ "hello/routers"

	"os"

	"github.com/astaxie/beego"
)

func main() {

	//beego.BConfig.WebConfig.AutoRender = false

	setLog()

	setStatic()
	//
	beego.Run()
}

func setStatic() {
	beego.SetStaticPath("/download", "download")
	beego.SetStaticPath("/bower_components", "bower_components")
	beego.SetStaticPath("/public", "public")
}

func setLog() {
	isPro := beego.BConfig.RunMode == "prod"

	if isPro {
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
