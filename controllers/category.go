package controllers

import (
	"hackerz_blog_golang/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	// 获取操作动作
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

	c.Data["IsCategory"] = true
	c.TplNames = "category.html"
	c.Data["IsLogin"] = checkLogin(c.Ctx)

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
