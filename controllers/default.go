package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Ctx.Input.Param("")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Fyn"] = map[string]string{
		"Email": "fyn@bohaisoft.cn",
	}
	c.TplName = "index.tpl"
}
