package controllers

import (
	"fmt"
	"sappo/models"
	"sappo/utils"

	"github.com/astaxie/beego"
)

type LoginchangeController struct {
	beego.Controller
}

func (c *LoginchangeController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Ctx.SetCookie("prg", "", -1, "/")
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login_change.html"

}

func (c *LoginchangeController) Post() {
	// 获取表单信息
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")       //旧密码
	newpwd := c.Input().Get("newpwd") //新密码
	tel := c.Input().Get("tel")
	prg := c.Input().Get("prg")
	pwdmd5 := utils.Md5(pwd) //转MD5加密

	// 验证用户名及密码
	user, err := models.GetUser(uname, pwdmd5)
	if err != nil {
		beego.Error(err)
	}
	if uname == user.Uname && //从数据库读取数据后并赋值
		pwdmd5 == user.Pwd {
		prg = user.Prgco //用户密码相同时再赋值Prg审批码

		err := models.UpdatetUser(uname, newpwd, prg, tel)
		fmt.Println(uname, newpwd, prg, tel)
		if err != nil {
			beego.Error(err)

		}

	}

	c.Redirect("/", 302)
	return
}
