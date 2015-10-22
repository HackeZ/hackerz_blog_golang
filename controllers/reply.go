package controllers

import (
	"hackerz_blog_golang/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	tid := c.Input().Get("tid")
	err := models.AddReply(tid, c.Input().Get("nickname"), c.Input().Get("content"))

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+tid, 302)
}

func (c *ReplyController) Delete() {
	if !checkLogin(c.Ctx) {
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteReply(c.Input().Get("rid"))

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+tid, 302)
}
