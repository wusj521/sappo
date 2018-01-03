package controllers

import (
	"sappo/models"

	"github.com/astaxie/beego"
)

type LoginaddController struct {
	beego.Controller
}

func (c *LoginaddController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Ctx.SetCookie("prg", "", -1, "/")
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login_add.html"

}

func (c *LoginaddController) Post() {
	// 获取表单信息
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	tel := c.Input().Get("tel")
	prg := c.Input().Get("prg")
	//	autoLogin := c.Input().Get("autoLogin") == "on"

	err := models.InsertUser(uname, pwd, prg, tel)
	if err != nil {
		beego.Error(err)

	}

	c.Redirect("/", 302)
	return
}
