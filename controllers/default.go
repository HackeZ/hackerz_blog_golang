package controllers

import (
	"github.com/astaxie/beego"
	"hackerz_blog_golang/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkLogin(c.Ctx)
	c.TplNames = "home.html"

	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics

	category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = category
}
