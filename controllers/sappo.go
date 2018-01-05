package controllers

import (
	"sappo/models"
	//"time"

	"github.com/astaxie/beego"
	//"simonwaldherr.de/go/golibs/arg"
	//saprfc "simonwaldherr.de/go/saprfc"
)

type SappoController struct {
	beego.Controller
}

func (c *SappoController) Get() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsSappo"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "sappo.html"
	//连接SAP

	// 解析表单-从表单字段中获取内容
	prg := c.Ctx.GetCookie("prg")
	var err error

	err = models.GetSappo(prg)
	if err != nil {
		c.Abort("401")
		beego.Error(err)
	}

	//连接SAP
	c.Redirect("/", 302)
	return
}
