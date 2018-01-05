package controllers

import (
	"fmt"
	"sappo/models"

	"github.com/astaxie/beego"
)

type PriceController struct {
	beego.Controller
}

func (c *PriceController) Get() {
	//登录检查
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsShouy"] = true
	c.TplName = "price.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//price := [10]int{86, 1114, 106, 106, 1107, 111, 1133, 221, 783, 2478}
	matnr := c.Ctx.GetCookie("matnr")
	datetime, price, err := models.GetPricelistOut(matnr)
	fmt.Println(datetime, price)
	if err != nil {
		beego.Error(err)
		c.Redirect("/daisp", 302)
		return
	}
	c.Data["datetime"] = datetime
	c.Data["Price"] = price
	//c.Data["json"] = &price //json传递
	//c.ServeJSON()//json传递

	/*	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "index.tpl"

		c.Data["TrueCond"] = true
		c.Data["FalseCond"] = false

		type u struct {
			Name string
			Age  int
			Sex  string
		}
		User := &u{
			Name: "wusj",
			Age:  36,
			Sex:  "Male",
		}
		c.Data["User"] = User

		numb := []int{1, 2, 3, 4, 5, 6}
		c.Data["Numb"] = numb
	*/

}
