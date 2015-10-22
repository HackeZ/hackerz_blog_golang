package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "rilegou" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplNames = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == pwd {
		maxAge := 0

		if autoLogin {
			maxAge = 1<<31 - 1
		}
		// 存储 Cookie
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	log.Println("POST??????")
	c.Redirect("/", 301)
	// return 是为了防止重新渲染页面
	return
}

func checkLogin(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == pwd
}
