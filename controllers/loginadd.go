package controllers

import (
	"sappo/models"
	"sappo/utils"

	"github.com/astaxie/beego"
)

type LoginaddController struct {
	beego.Controller
}

func (c *LoginaddController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("__uname", "", -1, "/")
		c.Ctx.SetCookie("__pwd", "", -1, "/")
		c.Ctx.SetCookie("__prg", "", -1, "/")
		c.Ctx.SetCookie("__prgpr", "", -1, "/")
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login_add.html"

}

func (c *LoginaddController) Post() {
	// 获取表单信息
	uname := c.Input().Get("uname")
	unamemd5 := utils.Md5(uname) //转MD5加密
	pwd := c.Input().Get("pwd")
	tel := c.Input().Get("tel")
	prg := c.Input().Get("prg")
	prgpr := c.Input().Get("prgpr")

	//	autoLogin := c.Input().Get("autoLogin") == "on"

	err := models.InsertUser(uname, unamemd5, pwd, prg, prgpr, tel)
	if err != nil {
		beego.Error(err)

	}

	c.Redirect("/", 302)
	return
}
