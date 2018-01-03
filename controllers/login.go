package controllers

import (
	"sappo/models"
	"sappo/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Ctx.SetCookie("prg", "", -1, "/")
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"

}

func (c *LoginController) Post() {
	// 获取表单信息
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	prg := c.Input().Get("prg")
	autoLogin := c.Input().Get("autoLogin") == "on"
	pwdmd5 := utils.Md5(pwd) //转MD5加密

	// 验证用户名及密码
	user, err := models.GetUser(uname, pwdmd5)
	if err != nil {
		beego.Error(err)

	}
	if uname == user.Uname && //从数据库读取数据后并赋值
		pwdmd5 == user.Pwd {
		prg = user.Prgco //用户密码相同时再赋值Prg审批码
		//	if uname == beego.AppConfig.String("adminName") &&
		//		pwd == beego.AppConfig.String("adminPass") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		//			c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwdmd5, maxAge, "/")
		c.Ctx.SetCookie("prg", prg, maxAge, "/")
	}

	c.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	// 验证用户名及密码,验证cookie中的用户密码与DB是否相同
	user, err := models.GetUser(uname, pwd)
	if err != nil {
		beego.Error(err)
		return false
	}
	return uname == user.Uname &&
		pwd == user.Pwd
}
