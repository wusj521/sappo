package controllers

import (
	"github.com/astaxie/beego"
	"sappo/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	//检查是否操作
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

}
