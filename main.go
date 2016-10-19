package main

import (
	_ "hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/download", "download")
	beego.Run()
}
