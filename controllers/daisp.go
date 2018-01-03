package controllers

import (
	"sappo/models"

	"github.com/astaxie/beego"
)

type DaispController struct {
	beego.Controller
}

func (c *DaispController) Get() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsDaisp"] = true
	c.TplName = "daisp.html"

	var err error
	// 解析表单-从表单字段中获取内容
	prg := c.Ctx.GetCookie("prg")
	c.Data["Sappo"], err = models.GetAllSappos(prg)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Tid"] = c.Ctx.Input.Param("0")

}

//审批PO
func (c *DaispController) Post() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	// 解析表单-从表单字段中获取内容
	tid := c.Input().Get("tid")
	ebeln := c.Input().Get("ebeln") //订单号
	prg := c.Ctx.GetCookie("prg")
	flag := "X" //c.Input().Get("flag")

	var err error
	//写入sappo表中flag审批标记为X
	err = models.ModifyDaisp(tid, ebeln, flag)
	if err != nil {
		beego.Error(err)
	}
	//调取审批RFC同时写入sappo中上传SAP标记
	// 解析表单-从表单字段中获取内容
	tid = c.Input().Get("tid")
	ebeln = c.Input().Get("ebeln")
	prg = c.Ctx.GetCookie("prg")
	uppo := "X"
	//审批前检查sappo本地表是否还有未审批的PO行项目，因为sappo表是以行项目存取的。
	ok := models.Checkpo(ebeln, prg)
	if ok == true {
		//没有查到同po号的行项目，说明是最后一个审批PO行项目，故调取SAP RFC审批。
		err = models.PostSapPo(tid, ebeln, prg, uppo)
		if err != nil {
			beego.Error(err)
		}
	}

	c.Redirect("/daisp", 302)
}

//列出将要审批的PO
func (c *DaispController) View() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsDaisp"] = true
	c.TplName = "daisp_view.html"

	tid := c.Input().Get("tid")
	sappo, err := models.GetDaisp(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/daisp", 302)
		return
	}

	c.Data["sappo"] = sappo
	c.Data["Tid"] = tid
	f := sappo.Flag
	c.Data["Flag"] = checkFlag(f)
	//SAP Po,在查看html时，代码中加入以下隐藏字段，方便此页面执行POST或其它动作。
	////func (c *DaispController) Post()中使用
	//c.Data["Ebeln"] = sappo.Ebeln

}

func checkFlag(f string) bool {
	if f == "X" {
		return false
	}
	return true

}
