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
		c.Ctx.SetCookie("__uname", "", -1, "/")
		c.Ctx.SetCookie("__pwd", "", -1, "/")
		c.Ctx.SetCookie("__prg", "", -1, "/")
		c.Ctx.SetCookie("__prgpr", "", -1, "/")
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
	prgpr := c.Input().Get("prgpr")
	autoLogin := c.Input().Get("autoLogin") == "on"
	pwdmd5 := utils.Md5(pwd)     //转MD5加密
	unamemd5 := utils.Md5(uname) //转MD5加密

	// 验证用户名及密码
	user, err := models.GetUser(uname, pwdmd5)
	if err != nil {
		beego.Error(err)

	}
	if uname == user.Uname && //从数据库读取数据后并赋值
		pwdmd5 == user.Pwd {
		prg = user.Prgco   //用户密码相同时再赋值Prg审批码-PO
		prgpr = user.Prgcr ////用户密码相同时再赋值Prg审批码-PR
		//	if uname == beego.AppConfig.String("adminName") &&
		//		pwd == beego.AppConfig.String("adminPass") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		//			c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("__uname", unamemd5, maxAge, "/")
		c.Ctx.SetCookie("__pwd", pwdmd5, maxAge, "/")
		c.Ctx.SetCookie("__prg", prg, maxAge, "/")
		c.Ctx.SetCookie("__prgpr", prgpr, maxAge, "/")
	}

	c.Redirect("/", 302)
	return
}

//登录检查，cookie中的值是否等于DB表中值
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("__uname")
	if err != nil {
		return false
	}
	unamemd5 := ck.Value //已在建立时MD5加密

	ck, err = ctx.Request.Cookie("__pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	// 验证用户名及密码,验证cookie中的用户密码与DB是否相同
	user, err := models.GetUsermd5(unamemd5, pwd)
	if err != nil {
		beego.Error(err)
	}
	return unamemd5 == user.Unamemd5 &&
		pwd == user.Pwd
}
