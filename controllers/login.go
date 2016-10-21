package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type user struct {
	ID       string `form:"-"`
	Name     string `form:"username"`
	Password string `form:"password"`
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}

func (c *LoginController) Post() {
	//

	//方法1
	//	user := c.GetString("username")
	//	password := c.GetString("password")

	//方法2
	//u := user{}
	//if err := c.ParseForm(&u); err != nil {
	//	beego.Error("cannot parse form ", err.Error())
	//}

	var name string
	c.Ctx.Input.Bind(&name, "username")
	var password string
	c.Ctx.Input.Bind(&password, "password")
	beego.Debug("something: ", name, password)
	//c.TplName = "LoginController/get.tpl"
	c.Data["json"] = "hello"
	//c.ServeJSON()
	c.Redirect("/", http.StatusTemporaryRedirect)
}
